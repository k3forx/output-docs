# Contexts and structs

## Introduction

The documentation for contexts state:

> Contexts should not be stored inside a struct type, but instead passed to each function that needs it.

## Prefer contexts passed as arguments

```go
// Worker fetches and adds works to a remote work orchestration server.
type Worker struct { /* … */ }

type Work struct { /* … */ }

func New() *Worker {
  return &Worker{}
}

func (w *Worker) Fetch(ctx context.Context) (*Work, error) {
  _ = ctx // A per-call ctx is used for cancellation, deadlines, and metadata.
}

func (w *Worker) Process(ctx context.Context, work *Work) error {
  _ = ctx // A per-call ctx is used for cancellation, deadlines, and metadata.
}
```

- この引数として`context.Context`を渡すデザインを用いることで、ユーザーはcallごとにデッドラインやキャンセレーション、metadataをセットできる

## Storing context in structs leads to confusion
