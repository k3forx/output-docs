package main

import (
	"iter"
)

func main() {

}

func NoValueIterator(yield func() bool) bool {
	return yield()
}

func FilterOdd(nums []int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for _, n := range nums {
			if n%2 == 1 {
				if !yield(n) {
					return
				}
			}
		}
	}
}
