package main

import "fmt"

func main() {
	for v := range f {
		fmt.Println(v)
		if v == 3 {
			return
		}
	}
}

func f(yield func(int) bool) {
	for _, num := range []int{1, 2, 3, 4, 5} {
		if !yield(num) {
			return
		}
	}
}
