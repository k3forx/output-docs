# `Value` vs `any`

## 疑問点

`Attr` 構造体のValueフィールドの型はなぜ `any` だとダメなのか？

```go
type Attr struct {
	Key   string
	Value Value // anyではダメなのか？
}
```


## 結論

- `any` が持つ実際の値へのポインタが `*byte` になっており、本来string型が持つポインタをanyが直接持つことによってメモリが少しだけ効率化されている

## 調査

## プログラム

```go
package main

import (
        "fmt"
        "unsafe"
)

type stringptr *byte

type Value struct {
        _   [0]func()
        num uint64
        any any
}

func StringValue(value string) Value {
        return Value{num: uint64(len(value)), any: stringptr(unsafe.StringData(value))}
}

func main() {
        str := "abcde"

        var s any = str
        fmt.Println(s)

        var ss Value = StringValue(str)
        fmt.Println(ss)

		fmt.Println(unsa)
}
```

### `s` のメモリを調べる

```bash
❯ dlv debug ./main.go
Type 'help' for list of commands.
(dlv) b main.go:24
Breakpoint 1 set at 0x10ae80f for main.main() ./main.go:24
(dlv) c
> main.main() ./main.go:24 (hits goroutine(1):1 total:1) (PC: 0x10ae80f)
    19:
    20: func main() {
    21:         str := "abc"
    22:
    23:         var s any = str
=>  24:         fmt.Println(s)
    25:
    26:         var ss Value = StringValue(str)
    27:         fmt.Println(ss)
    28: }
(dlv) p s
interface {}(string) "abc"
(dlv) p *((*runtime.iface)(uintptr(&s)))
runtime.iface {
        tab: *runtime.itab {
                inter: *(*"internal/abi.InterfaceType")(0x10),
                _type: *(*"internal/abi.Type")(0x8),
                hash: 125357496,
                _: [4]uint8 [7,8,8,24],
                fun: [1]uintptr [17639976],},
        data: unsafe.Pointer(0xc000014290),}
(dlv) x -fmt hex -count 16 -size 1 0xc000014290
0xc000014290:   0x5d   0xa8   0x0c   0x01   0x00   0x00   0x00   0x00
0xc000014298:   0x03   0x00   0x00   0x00   0x00   0x00   0x00   0x00
(dlv) x -fmt hex -count 1 -size 8 0xc000014290
0xc000014290:   0x00000000010ca85d
(dlv) x -fmt hex -count 5 -size 1 0x00000000010ca85d
0x10ca85d:   0x61   0x62   0x63   0x6e   0x69
(dlv) exit
```

- `any` の `data` は `unsafe.Pointer(0xc000014290)` となっている
- `0xc000014290` を調べると `0x00000000010ca85d` が表示された
- `0x00000000010ca85d` を調べると `0x61`, `0x62`, `0x63` が表示されて `abc` と対応している

### `ss` のメモリを調べる

```bash
❯ dlv debug ./main.go
Type 'help' for list of commands.
(dlv) b main.go:27
Breakpoint 1 set at 0x10ae884 for main.main() ./main.go:27
(dlv) c
abc
> main.main() ./main.go:27 (hits goroutine(1):1 total:1) (PC: 0x10ae884)
    22:
    23:         var s any = str
    24:         fmt.Println(s)
    25:
    26:         var ss Value = StringValue(str)
=>  27:         fmt.Println(ss)
    28: }
(dlv) p ss
main.Value {
        _: [0]func() [],
        num: 3,
        any: interface {}(main.stringptr) *97,}
(dlv) p ss.any
interface {}(main.stringptr) *97
(dlv) p *((*runtime.iface)(uintptr(&ss.any)))
runtime.iface {
        tab: *runtime.itab {
                inter: *(*"internal/abi.InterfaceType")(0x8),
                _type: *(*"internal/abi.Type")(0x8),
                hash: 174749434,
                _: [4]uint8 [15,8,8,54],
                fun: [1]uintptr [17639928],},
        data: unsafe.Pointer(0x10ca85d),}
(dlv) x -fmt hex -count 16 -size 1 0x10ca85d
0x10ca85d:   0x61   0x62   0x63   0x6e   0x69   0x6c   0x31   0x32
0x10ca865:   0x35   0x36   0x32   0x35   0x4e   0x61   0x4e   0x45
(dlv) exit
```

- `any` の `data` は `unsafe.Pointer(0x10ca85d)` となっている
- `0x10ca85d` を調べると `0x61`, `0x62`, `0x63` となっていて、`abc` と対応している

### バイトのサイズを調べてみる

先ほどのプログラムに以下の行を追加

```go
	fmt.Printf("byte size of str: %d\n", unsafe.Sizeof(str))
	fmt.Printf("byte size of s: %d\n", unsafe.Sizeof(s))
	fmt.Printf("byte size of ss: %d\n", unsafe.Sizeof(ss))
```

実行してみる

```bash
byte size of str: 16
byte size of s: 16
byte size of ss: 24
```

- `str` を `any` で表現すると、32byte (16byte + 16byte)
- `str` を `Value` で表現すると、24byte
