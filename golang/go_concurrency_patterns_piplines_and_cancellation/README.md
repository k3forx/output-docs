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