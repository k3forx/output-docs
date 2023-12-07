# `Handler` インターフェイスの `WithAttrs` メソッド

## シグネチャ

```go
type Handler interface {
	...
	WithAttrs(attrs []Attr) Handler
	...
}
```

## 詳細

- `handelState` が詳細を持つ

```go
// handleState holds state for a single call to commonHandler.handle.
// The initial value of sep determines whether to emit a separator
// before the next key, after which it stays true.
type handleState struct {
	h       *commonHandler
	buf     *buffer.Buffer
	freeBuf bool           // should buf be freed?
	sep     string         // separator to write before next key
	prefix  *buffer.Buffer // for text: key prefix
	groups  *[]string      // pool-allocated slice of active groups, for ReplaceAttr
}
```

## メモ

- bufferにAttrsを入れて新しいhandlerを生成している
  - bufの生成にはsync.Poolが使われている
  - appendKeyやappendValue中で、bufに値を入れている
  - そのbufを持った新しいhandlerを返す
  - deferの部分も書く
- preformattedAttrsは使われていない？
  - 聞く
- なぜ初期化は1024バイト？
  - ある程度大きな容量を最初に確保しておいて、容量を追加するときのappendによる容量の再確保等を防ぐ
- なぜ大きい場合はリセットしないのか？
  - メガバイトのログエントリがbufに入った時を考えてみる
  - 仮にそのbufをpoolに戻したとしても、再利用されない可能性もある&GCからも回収されない可能性もある
  - そういった状況を避けるために大きすぎるbufはPutを呼ばないような設計になっている
