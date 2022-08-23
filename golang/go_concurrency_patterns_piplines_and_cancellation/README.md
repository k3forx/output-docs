# Go Concurrency Patterns: Pipelines and cancellation

## Introduction

Goの同時並行の原則は効率的なI/Oや複数のCPUを効率的に使用するストリーミングデータ処理の構築を容易にする。

## What is a pipeline?

形式的なパイプラインの定義はGoに存在しない、様々な種類の同時実行のプログラムの1つに過ぎない。非公式には、パイプラインはチャネルによって紐づいた一連のステージである、

## Squaring numbers

3つのステージのパイプラインを考える。

最初のステージ `gen` は整数のリストを、その整数を発するチャネルに変換する関数である。`gen` 関数はチャネルに整数を送るゴルーチンを開始させ、全ての値を送り終えたらチャネルを閉じる。

```golang
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}
```

2番目のステージ `sq` はチャネルから整数を受け取り、受け取ったそれぞれの整数の2乗を発するチャネルを返す。

```golang
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
```

## Fan-out, fan-in

チャネルが閉じるまで、複数の関数がそのチャネルから読み込むことができる、このことはfan-outと呼ばれる。このことは、CPU使用やI/Oを並列するようなワーカーのグループ間にタスクを分散する方法を提供する。

1つの関数は複数のインプットから読み込んで、

私たちは、同じインプットのチャネルから読み込みを行う2つの `sq` インスタンスを走らせるようにパイプラインを変更することができる。私たちは新しい関数 `merge` を導入する。

```golang
func main() {
    in := gen(2, 3)

    // Distribute the sq work across two goroutines that both read from in.
    c1 := sq(in)
    c2 := sq(in)

    // Consume the merged output from c1 and c2.
    for n := range merge(c1, c2) {
        fmt.Println(n) // 4 then 9, or 9 then 4
    }
}
```

`merge` 関数は、それぞれのインバウンドのチャネルに対してアウトバンドのチャネルに値をコピーするような1つのゴルーチンをスタートさせることで、チャネルのリストを1つのチャネルに変換する。一度全てのアウトバンドのゴルーチンがスタートすると、`merge` はもう一つのゴルーチンを開始させ、チャネルが全て終了した後にアウトバンドのチャネルを閉じる。

閉じたチャネルへの送信はパニックになるので、doneを呼ぶ前に全ての送信が終了していることを保証することが重要です。`sync.WaitGroup` はこのsynchronizationをarrangeする簡単な方法を提供します。

```golang
func merge(cs ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int)

    // Start an output goroutine for each input channel in cs.  output
    // copies values from c to out until c is closed, then calls wg.Done.
    output := func(c <-chan int) {
        for n := range c {
            out <- n
        }
        wg.Done()
    }
    wg.Add(len(cs))
    for _, c := range cs {
        go output(c)
    }

    // Start a goroutine to close out once all the output goroutines are
    // done.  This must start after the wg.Add call.
    go func() {
        wg.Wait()
        close(out)
    }()
    return out
}
```

## Stopping short

私たちのパイプライン関数にはあるパターンがあります。

- ステージは全ての送信操作が終了した時に、アウトバンドのチャネルを閉じます
- ステージは、インバウンドのチャネルが閉じられるまで、それらから値を受け取り続けます

このパターンはそれそれの受け取りステージが...し、全ての値がダウンストリームの送ることが成功したら全てのゴルーチンがexitすることを保証している。

しかし、現実のパイプラインでは、ステージは全てのインバウンドの値を常に受け取るわけではない。時々、これは仕様による。受信用のゴルーチンは処理を行うために値の一部しか必要なかもしれない。より多くの場合、受信した値が以前のステージのエラーを表すため、ステージは早期に終了する。いずれのケースでも、受信チャネルは残りの値が届くのを待つ必要はなく、先のステージは後のステージが必要としない値を生成することを止めさせたい。

私たちのパイプラインの例では、もしステージが全てのインバウンドの値を消費するのを失敗したら、それらの値を送ろうとしているゴルーチンは無限にブロックされるだろう。

```golang
    // Consume the first value from the output.
    out := merge(c1, c2)
    fmt.Println(<-out) // 4 or 9
    return
    // Since we didn't receive the second value from out,
    // one of the output goroutines is hung attempting to send it.
}
```

これはリソースリークである、ゴルーチンはメモリとランタイムリソースを消費し、ゴルーチンスタックのヒープリファレンスはデータをガベージコレクタされることから守る。ゴルーチンはガベージコレクタされない。

下流のステージがすべてのインバウンド値を受信できなかった場合でも、パイプラインの上流ステージが終了するように手配する必要があります。これを実現する一つの方法は、送信チャンネルをバッファを持つように変更することです。バッファは一定数の値を保持することができます。バッファに余裕があれば、送信操作は直ちに完了します。

```golang
c := make(chan int, 2) // buffer size 2
c <- 1  // succeeds immediately
c <- 2  // succeeds immediately
c <- 3  // blocks until another goroutine does <-c and receives 1
```

チャネルの作成時に送られる値の数が知られている時、バッファはコードをシンプルにすることができる。例えば、バッファされたチャネルに整数のリストをコピーし、新しいゴルーチンを生成することを避けるように `gen` 関数を変更することができる。

```golang
func gen(nums ...int) <-chan int {
    out := make(chan int, len(nums))
    for _, n := range nums {
        out <- n
    }
    close(out)
    return out
}
```

パイプラインのブロックされたゴルーチンに戻る時、`merge` 関数によって返されるアウトバンドのチャネルにバッファを加えることを考えないといけないかもしれない。

```golang
func merge(cs ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int, 1) // enough space for the unread inputs
    // ... the rest is unchanged ...
```

## Explicit cancellation

`main` が `out` からの全ての値を受け取ることなく終了することになった時、アップストリームのステージにあるごルーチンに、それらが送ろうとしている値を破棄することを伝えなければならない。