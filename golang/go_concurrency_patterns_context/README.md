# Go Concurrency Patterns: Context

## Introduction

Goのサーバーにおいて、やってくるそれぞれのリクエストは、それぞれのゴルーチンの中で処理される。リクエストのハンドラーはしばしば、データベースやRPCサービスのようなバックエンドに接続するために追加でゴルーチンをスタートさせることもある。リクエストに作用するゴルーチンの集合は通常、エンドユーザーのアイデンティティや認証トークン、リクエストのデッドラインといったリクエストに特有な値にアクセスする必要がある。リクエストがキャンセルされたり、タイムアウトしたりした時、そのリクエストに作用する全てのゴルーチンは、システムがそれらが使っていたいかなるリソースも回収できるように、素早く終了しなければならない。

Googleでは、リクエストスコープな値やキャンセルシグナル、API境界をまたぐデッドラインを、リクエストを処理するのに関わる全てのゴルーチンに簡単に渡せるようにするために渡せるようにするために `context` パッケージを開発した。

## Context

`context` パッケージのコアの部分は `Context` 型である。

```golang
// A Context carries a deadline, cancellation signal, and request-scoped values
// across API boundaries. Its methods are safe for simultaneous use by multiple
// goroutines.
type Context interface {
    // Done returns a channel that is closed when this Context is canceled
    // or times out.
    Done() <-chan struct{}

    // Err indicates why this context was canceled, after the Done channel
    // is closed.
    Err() error

    // Deadline returns the time when this Context will be canceled, if any.
    Deadline() (deadline time.Time, ok bool)

    // Value returns the value associated with key or nil if none.
    Value(key interface{}) interface{}
}
```

`Done` メソッドは `Context` それ自身の上で動作している関数に対するキャンセルのシグナルとして作用するチャネルを返す。

### Derived contexts

`context` パッケージは、既存のコンテキストから新しい `Context` の値を派生するための関数を提供する。これらのコンテキストはツリーを形成する。`Context` がキャンセルされた時、それから派生した全ての `Context` もキャンセルされるということである。

`Background` はいかなる `Context` ツリーの根になり、決してキャンセルされることはない。

```golang
// Background returns an empty Context. It is never canceled, has no deadline,
// and has no values. Background is typically used in main, init, and tests,
// and as the top-level Context for incoming requests.
func Background() Context
```

`WithCancel` と `WithTimeout` は、親の `Context` より早くキャンセルされうる派生 `Context` の値を返します。リクエストのハンドラーが処理を終わった時、処理されるリクエストに紐づく `Context` は一般的にはキャンセルされる。`WithCancel` は、複数のレプリカを使っている時、無駄なリクエストをキャンセルするのに役に立つ。`WithTimeout` はバックエンドサーバーに対するリクエストに対してデッドラインを設定するのに役立つ。

```golang
// WithCancel returns a copy of parent whose Done channel is closed as soon as
// parent.Done is closed or cancel is called.
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)

// A CancelFunc cancels a Context.
type CancelFunc func()

// WithTimeout returns a copy of parent whose Done channel is closed as soon as
// parent.Done is closed, cancel is called, or timeout elapses. The new
// Context's Deadline is the sooner of now+timeout and the parent's deadline, if
// any. If the timer is still running, the cancel function releases its
// resources.
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
```

`WithValue` は `Context` にリクエストに特有の値を紐づけるための方法を提供する。

```golang
// WithValue returns a copy of parent whose Value method returns val for key.
func WithValue(parent Context, key interface{}, val interface{}) Context
```

## Example: Google Web Search

### The server program

### Package userip

### Package google

## Adapting code for Contexts

## Conclusion