# Structured Logging with slog

- https://go.dev/blog/slog

## Introduction

- For servers, logging is an important way for developers to observe the detailed behavior of the system, and often the first place they go to debug it. Logs therefore tend to be voluminous, and the ability to search and filter them quickly is essential.
  - サーバーにとって、ロギングは開発者がシステムの詳細な動作を観察するための重要な方法であり、しばしばデバッグのために最初に行く場所でもある。そのため、ログは膨大な量になりがちであり、それらを素早く検索し、フィルタリングする能力は不可欠である。
- Over time, we've learned that structured logging is important to Go programmers. It has consistently ranked high in our annual survey, and many packages in the Go ecosystem provide it.
  - 時間をかけて、構造化ロギングがGoプログラマーにとって重要であることがわかりました。私たちの年次調査でも常に上位にランクインしており、Goエコシステムの多くのパッケージがロギングを提供しています。

## A tour of `slog`

```go
package main

import "log/slog"

func main() {
    slog.Info("hello, world")
}
```

---

以下ように `Attrs` 型と `LogAttrs` メソッドを呼び出すことはメモリアロケーションを最小にする。

```go
slog.LogAttrs(context.Background(), slog.LevelInfo, "hello, world",
    slog.String("user", os.Getenv("USER")))
```

There is a lot more to `slog`:
* As the call to `LogAttrs` shows, you can pass a `context.Context` to some log functions so a handler can extract context information like trace IDs.
  * `context.Context` を何かしらのログ関数に渡すことができ、handlerはトレースIDのようなcontextの情報を取り出すことできる
* You can call `Logger.With` to add attributes to a logger that will appear in all of its output, effectively factoring out the common parts of several log statements. This is not only convenient, but it can also help performance, as discussed below.
  * `Logger.With` 呼び出すことで、全ての出力に現れるようなattributesをロガーに追加できる

## Performance

- We wanted `slog` to be fast. For large-scale performance gains, we designed the `Handler` interface to provide optimization opportunities. The `Enabled` method is called at the beginning of every log event, giving the handler a chance to drop unwanted log events quickly. The WithAttrs and WithGroup methods let the handler format attributes added by Logger.With once, rather than at each logging call. This pre-formatting can provide a significant speedup when large attributes, like an http.Request, are added to a Logger and then used in many logging calls.