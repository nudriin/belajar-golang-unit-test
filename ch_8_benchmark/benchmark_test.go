package ch_8_benchmark

import (
	"testing"

	"github.com/nudriin/belajar-golang-unit-test/ch_1_helper"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch_1_helper.SayHello("Nurdin")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("test_1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ch_1_helper.SayHello("Nurdin")
		}
	})
	b.Run("test_2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ch_1_helper.SayHello("Nurdin")
		}
	})
}

func BenchmarkTable(b *testing.B) {
	// Membuat table untuk test
	tests := []struct {
		name    string
		request string
	}{
		{
			name:    "SayHello(Nurdin)",
			request: "Nurdin",
		},
		{
			name:    "SayHello(Sasuke)",
			request: "Sasuke",
		},
		{
			name:    "SayHello(Naruto)",
			request: "Naruto",
		},
	}

	// ! Membuat table test dengan data yang berbeda-beda
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ch_1_helper.SayHello(test.request)
			}
		})
	}
}
