package ch_1_helper

import (
	"fmt"
	"testing"
)

// ! Nama function untuk unit test harus diawali dengan Test

// * Aturan pembuatan testing, wajib ada (t *testing.T) pada paramter

func TestSayHelloWithFail(t *testing.T) {
	result := SayHello("Nurdins")
	if result != "Hello Nurdin" {
		t.Fail() //! menggagalkan tapi tetap lanjut ke kode selanjutnya selanjutnya
	}
	fmt.Println("Masih di print ke sini")
}

func TestSayHelloWithFailNow(t *testing.T) {
	result := SayHello("Nurdins")
	if result != "Hello Nurdin" {
		t.FailNow() //! menggagalkan tapi tidak lanjut ke kode selanjutnya selanjutnya
	}
	fmt.Println("tidak di print ke sini")
}

func TestSayHelloWithError(t *testing.T) {
	result := SayHello("Nurdins")
	if result != "Hello Nurdin" {
		t.Error("Result must be 'Hello Nurdin'") //! mengagalkan dengan error message dan memanggil Fail
	}
	fmt.Println("Masih di print ke sini")
}

func TestSayHelloWithFatal(t *testing.T) {
	result := SayHello("Nurdins")
	if result != "Hello Nurdin" {
		t.Fatal("Result must be 'Hello Nurdin'") //! mengagalkan dengan error message dan memanggil FailNow
	}
	fmt.Println("tidak di print ke sini")
}
