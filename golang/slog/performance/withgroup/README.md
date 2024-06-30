# `Handler` インターフェースの `WithGroup` メソッド

## メソッドの詳細

```go
type Handler interface {
	// ...
	WithGroup(name string) Handler
}
```

## 詳細

- `WithGroup` メソッドは `commonHandler` 構造体の `groups` にグループ名を追加する

```go
type commonHandler struct {
	json              bool // true => output JSON; false => output text
	opts              HandlerOptions
	preformattedAttrs []byte
	// groupPrefix is for the text handler only.
	// It holds the prefix for groups that were already pre-formatted.
	// A group will appear here when a call to WithGroup is followed by
	// a call to WithAttrs.
	groupPrefix string
	groups      []string // all groups started from WithGroup
	nOpenGroups int      // the number of groups opened in preformattedAttrs
	mu          *sync.Mutex
	w           io.Writer
}

// HandlerインターフェースのWithGroupメソッドでは以下のメソッドが呼ばれる
func (h *commonHandler) withGroup(name string) *commonHandler {
	// cloneメソッドでは特にgroupに関することは何もやっていなさそう
	h2 := h.clone()
	h2.groups = append(h2.groups, name)
	return h2
}
```

- `handle` メソッドでどう処理されているか？

```go
func (h *commonHandler) handle(r Record) error {
	// NOTE: newHandleStateでも特にgroupに関する処理はなさそう
	state := h.newHandleState(buffer.New(), true, "")
	defer state.free()
	if h.json {
		state.buf.WriteByte('{')
	}
	// Built-in attributes. They are not in a group.
	// nil
	stateGroups := state.groups
	state.groups = nil // So ReplaceAttrs sees no groups instead of the pre groups.
	rep := h.opts.ReplaceAttr
	// time
	if !r.Time.IsZero() {
		key := TimeKey
		val := r.Time.Round(0) // strip monotonic to match Attr behavior
		if rep == nil {
			state.appendKey(key)
			state.appendTime(val)
		} else {
			state.appendAttr(Time(key, val))
		}
	}
	// level
	key := LevelKey
	val := r.Level
	if rep == nil {
		state.appendKey(key)
		state.appendString(val.String())
	} else {
		state.appendAttr(Any(key, val))
	}
	// source
	if h.opts.AddSource {
		state.appendAttr(Any(SourceKey, r.source()))
	}
	key = MessageKey
	msg := r.Message
	if rep == nil {
		state.appendKey(key)
		state.appendString(msg)
	} else {
		state.appendAttr(String(key, msg))
	}
	state.groups = stateGroups // Restore groups passed to ReplaceAttrs.
	state.appendNonBuiltIns(r) // NOTE: ここでgroupが処理されている
	state.buf.WriteByte('\n')

	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.w.Write(*state.buf)
	return err
}
```

- `appendNonBuiltIns` メソッドをみてみる

```go
func (s *handleState) appendNonBuiltIns(r Record) {
	// preformatted Attrs
	if len(s.h.preformattedAttrs) > 0 {
		s.buf.WriteString(s.sep)
		s.buf.Write(s.h.preformattedAttrs)
		s.sep = s.h.attrSep()
	}
	// Attrs in Record -- unlike the built-in ones, they are in groups started
	// from WithGroup.
	// If the record has no Attrs, don't output any groups.
	// s.h.nOpenGroups = 0のはず
	nOpenGroups := s.h.nOpenGroups
	if r.NumAttrs() > 0 {
		s.prefix.WriteString(s.h.groupPrefix)
		s.openGroups()
		// 1になる
		nOpenGroups = len(s.h.groups)
		r.Attrs(func(a Attr) bool {
			s.appendAttr(a)
			return true
		})
	}
	if s.h.json {
		// Close all open groups.
		for range s.h.groups[:nOpenGroups] {
			s.buf.WriteByte('}')
		}
		// Close the top-level object.
		s.buf.WriteByte('}')
	}
}
```

## その他雑多めも

- unopened group
  - outputに出力していないグループのこと
- groupのValueでの表現
  - numがAttrsの数、anyは*Attr (Attrへのポインタ) になっている 
- unsafe.SliceDataを使って[]Attrを*Attrsにしている
  - https://pkg.go.dev/unsafe#SliceData
