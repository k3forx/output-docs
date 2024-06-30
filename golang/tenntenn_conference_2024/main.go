package main

import "fmt"

func main() {
	// pointerに変換する
	ptr := Ptr[bool](false)
	fmt.Println(ptr)

	// スライスに関数を適用する
	var sum int
	Apply([]int{10, 20}, func(i, v int) {
		sum += v
	})
	fmt.Println(sum)

	// スライスの要素をフィルターする
	ns := []int{1, 2, 3, 4}
	ms := Filter(ns, func(i, v int) bool {
		return v%2 == 0
	})
	fmt.Println(ms)

	// スライスを別のスライスに変換する
	var ss []string = Map([]int{10, 20}, func(n int) string {
		return fmt.Sprintf("0x%x", n)
	})
	fmt.Println(ss)

	// 任意の2つの型のフィールドXとYを持つ構造体のポインタを返す
	var t *Tuple[int, string] = New(10, "apple")
	fmt.Println(t.X, t.Y)
}

func Ptr[T any](v T) *T {
	return &v
}

func Apply[E any](s []E, f func(int, E)) {
	for i, v := range s {
		f(i, v)
	}
}

func Filter[E any](s []E, f func(int, E) bool) []E {
	res := make([]E, 0, len(s))
	for i, v := range s {
		if f(i, v) {
			res = append(res, v)
		}
	}
	return res
}

func Map[S ~[]E, E, V any](s S, f func(E) V) []V {
	res := make([]V, len(s))
	for i, v := range s {
		res[i] = f(v)
	}
	return res
}

type Tuple[X, Y any] struct {
	X X
	Y Y
}

func New[T1, T2 any](t1 T1, t2 T2) *Tuple[T1, T2] {
	return &Tuple[T1, T2]{
		X: t1,
		Y: t2,
	}
}
