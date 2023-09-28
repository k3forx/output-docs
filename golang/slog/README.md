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

## パフォーマンス

## 設計の過程