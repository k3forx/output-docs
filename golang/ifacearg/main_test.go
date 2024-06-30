package main

import "testing"

func Benchmark_StringFromBasicInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bi := BasicInterfaceImpl{}
		StringFromBasicInterface(bi)
	}
}

func Benchmark_StringFromGeneralInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gi := GeneralInterfaceImpl{}
		StringFromGeneralInterface(gi)
	}
}
