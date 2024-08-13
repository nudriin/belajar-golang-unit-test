package ch_2_assertion

import (
	"fmt"
	"testing"

	"github.com/nudriin/belajar-golang-unit-test/ch_1_helper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSayHelloAssert(t *testing.T) {
	result := ch_1_helper.SayHello("Nurdin")
	assert.Equal(t, "Hello Nurdins", result) // (t, expected, actual)
	assert.NotEqual(t, "", result)           // (t, expected, actual)
	fmt.Println("di print ke sini")
}

func TestSayHelloRequire(t *testing.T) {
	result := ch_1_helper.SayHello("Nurdin")
	require.Equal(t, "Hello Nurdins", result) // (t, expected, actual)
	require.NotEqual(t, "", result)           // (t, expected, actual)
	fmt.Println("tidak di print ke sini")
}
