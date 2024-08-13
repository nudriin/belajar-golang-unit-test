package ch_4_before_after

import (
	"fmt"
	"testing"

	"github.com/nudriin/belajar-golang-unit-test/ch_1_helper"
	"github.com/stretchr/testify/assert"
)

// ! Membuat test main untuk mengatur eksekusi test
// ! Function ini akan di eksekusi per package 1 kali
// * Namanya wajib TestName
func TestMain(m *testing.M) {
	fmt.Println("Run sebuah code SEBELUM di run testnya")

	//! Memanggil semua unit test untuk di run
	m.Run()

	fmt.Println("Run sebuah code SESUDAH di run testnya")
}

func TestSayHello(t *testing.T) {
	result := ch_1_helper.SayHello("Nurdin")
	assert.Equal(t, "Hello Nurdin", result)
}
