# [Structured Logging with slog](https://go.dev/blog/slog)

## 導入

- 開発者にとって、ログはシステムの詳細な振る舞いを観察するための重要な手法であり、サーバーをデバックするために最初にみる場所だったりする
- なので、ログは膨大な量になりがちであり、それらを素早く検索し、フィルタリングする能力が不可欠となる
- 多くの構造化ロギングパッケージから選ぶことができるため、大規模なプログラムでは、依存関係を通じて複数のパッケージを含むことになることが多い
- mainプログラムは、ログ出力が一貫しているように、これらのロギング・パッケージをそれぞれ設定しなければならないかもしれない

## A Tour of `slog`

`slog`を使った最も単純なプログラムの例

```go
package main

import "log/slog"

func main() {
	slog.Info("hello, world")
}
```

出力は以下のようになる。

```bash
2023/09/28 21:13:37 INFO hello, world
```

これは `log.Printf` を使った時と`INFO`以外の出力は全く同じ (内部的には`log`パッケージのデフォルトロガーと同じ)

----

`log`パッケージとは違い、メッセージの後に書くことでkey-valueペアを簡単に追加できる

```go
slog.Info("hello, world", "user", os.Getenv("USER"))
```

出力は以下のようになる

```bash
2023/09/28 21:22:15 INFO hello, world user=kanata-miyahana
```

`slog`のトップレベルの関数はデフォルトロガーを明示的に呼ぶことができる。

```go
logger := slog.Default()
logger.Info("hello, world", "user", os.Getenv("USER"))
```

---

ロガーで使われる*handler*を変えることで、出力を変えることができる。`slog`は3つのビルトインhandlerを備えている。`TextHandler`は全ての情報を`key=value`の形式で出力する

```go
logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
logger.Info("hello, world", "user", os.Getenv("USER"))
```

出力は以下のようになる

```bash
time=2023-09-28T21:31:49.423+09:00 level=INFO msg="hello, world" user=kanata-miyahana
```

JSON形式の出力には`JSONHandler`を使う

```go
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
logger.Info("hello, world", "user", os.Getenv("USER"))
```

```bash
{"time":"2023-09-28T21:33:27.685697+09:00","level":"INFO","msg":"hello, world","user":"kanata-miyahana"}
```

`slog.Handler`インターフェイスを実装することで、誰でもhandlerを書くことができる。

---

今まで見てきたkey-valueのsyntaxは便利だが、頻繁に実行されるログ文では `Attr` 型を使い、`LogAttrs` メソッドを使った方が効率的。これらはメモリ割り当てを最小にする。

```go
slog.LogAttrs(context.Background(), slog.LevelInfo, "hello, world",
    slog.String("user", os.Getenv("USER")))
```

他の`slog`の機能

- `LogAttrs`の呼び出しが示すように、`context.Context`をいくつかのログ関数に渡すことで、handlerはトレースIDのようなコンテキスト情報を抽出することができます。(コンテキストをキャンセルしても、ログエントリが書き込まれるのを防ぐことはできません)。
- `Logger.With`を呼び出すと、ロガーのすべての出力に表示される属性を追加し、複数のログ文の共通部分を効果的にファクタリングすることができます。これは便利なだけでなく、後述するように、パフォーマンスにも役立ちます。
- 属性はグループにまとめることができる。これにより、ログ出力がより構造化され、同一でなければならないキーの曖昧さをなくすことができる。
- `LogValue`メソッドで型を指定することで、値がログにどのように表示されるかを制御できます。これを使用して、構造体のフィールドをグループとしてログに記録したり、機密データを編集したりできます。


## パフォーマンス

大規模なパフォーマンス向上のために、最適化の機会を提供する`Handler`インターフェイスを設計した。

```go
type Handler interface {
	Enabled(context.Context, Level) bool
	Handle(context.Context, Record) error
	WithAttrs(attrs []Attr) Handler
	WithGroup(name string) Handler
}
```

- パフォーマンス最適化作業に情報を提供するために、既存のオープンソースプロジェクトにおけるロギングの典型的なパターンを調査した
- ロギング・メソッドへの呼び出しの95％以上が、5つ以下の属性を渡していることを発見した。また、属性の種類を分類し、一握りの一般的な種類が大部分を占めていることを発見した。
- 一般的なケースを捉えたベンチマークを作成し、時間の経過を見るためのガイドとして使用した。
- メモリ割り当てに細心の注意を払うことで、最大の効果が得られた。

## 設計の過程

-----

# GopherConUK メモ

## The Problem

- An application may use many packages, each with its own logging
- Hard to keep log output consistent
- A standard library package can help logging packages work together

## A Tour of `slog`

## Goal

- Easy to use
- Fast
- Interoperates with existing packages, including `log`

## Architecture

- Logger (front api)
- Handler (backend api)
- Record (sending data format)

## Attrs and Values

- `slog.LogAttrs` を使った方がパフォーマンス的に良い

```go
// An Attr is a key-value pair.
type Attr struct {
	Key   string
	Value Value
}

// A Value can represent any Go value, but unlike type any,
// it can represent most small values without an allocation.
// The zero Value corresponds to nil.
type Value struct {
	_ [0]func() // disallow ==
	// num holds the value for Kinds Int64, Uint64, Float64, Bool and Duration,
	// the string length for KindString, and nanoseconds since the epoch for KindTime.
	num uint64
	// If any is of type Kind, then the value is in num as described above.
	// If any is of type *time.Location, then the Kind is Time and time.Time value
	// can be constructed from the Unix nanos in num and the location (monotonic time
	// is not preserved).
	// If any is of type stringptr, then the Kind is String and the string value
	// consists of the length in num and the pointer in any.
	// Otherwise, the Kind is Any and any is the value.
	// (This implies that Attrs cannot store values of type Kind, *time.Location
	// or stringptr.)
	any any
}
```

## Levels

- Levels are integers
- Some have names (Debug, Info, Warn, Error)
- There's room in between for your own
- Info is zero
- Fixed offset from OpenTelemetry level numbers

- `slog.LevelVar` を使えば任意にレベルを設定できる

## Groups

- `slog.Group` で入れ子になるようなJSONを生成できる
- `With` を使えば、値を渡すことを必要とせずに、必要な値を出力できる

```go
// rは*http.Request
logger := slog.Default().With("path", r.URL.path)
```

- 任意の構造体に `slog.Value` を返す `LogValue()` メソッドを定義すれば、それが使われる

### Problems

二つのアプリケーションが `value` という同じキーでそれぞれint/stringをvalueとするログを生み出していたらどうなるのか？

- この場合は `slog.WithGroup` を使うとよい

## Why does speed matter?

Logging can be in the inner loop

- Servers process thousands of requests per minute
- Each request can generate thousands of log messages
- Each log message may require non-trivial processing

**These are the sort of reasons!**

To really understand why you have to look at the history of structured logging in Go.

- 2014: logrus (xx imports) -> slow
- 2017: zap (xxx imports) -> fast
- 2017: zerolog (xx imports) -> faster

-> slog aims for Zap performance

zapがよくインポートされているし (多くの人に支持されているし)、Goの標準パッケージの目的は世界で最も早いログパッケージを提供することではない

## Four Steps to High Performance

1. Know your use cases
1. Write good benchmarks
1. Design the right API
1. Avoid allocation

### Know your use cases (使い道を知る)

- zapのログでどれだけのkey-valueペアが呼ばれているか調べた
