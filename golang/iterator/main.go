package main

import "fmt"

func main() {
	for v := range g(f) {
		fmt.Println(v)
	}
}

func f(yield func(int) bool) {
	for _, num := range []int{1, 0, 2, 3, 4, 5} {
		if !yield(num) {
			return
		}
	}
}

func g(f func(yield func(int) bool)) func(yield func(string) bool) {
	maps := map[int]string{
		1: "one", 2: "two", 3: "three",
		4: "four", 5: "five",
	}
	return func(yield func(string) bool) {
		for v := range f {
			str, ok := maps[v]
			if len(str) > 3 {
				break
			}
			if ok && !yield(str) {
				return
			}
		}
	}
}
