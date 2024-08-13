package ch_3_skip_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/nudriin/belajar-golang-unit-test/ch_1_helper"
	"github.com/stretchr/testify/require"
)

// !  Function ini akan di skip
func TestSkip(t *testing.T) {
	// ! Jika os nya windows maka kita skip
	if runtime.GOOS == "windows" {
		t.Skip("Can't run on Windows")
	}

	result := ch_1_helper.SayHello("Nurdin")
	require.Equal(t, "Hello Nurdins", result) // (t, expected, actual)
	require.NotEqual(t, "", result)           // (t, expected, actual)
	fmt.Println("tidak di print ke sini")
}
