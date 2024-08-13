package ch_2_assertion

import (
	"testing"

	"github.com/nudriin/belajar-golang-unit-test/ch_1_helper"
	"github.com/stretchr/testify/assert"
)

func TestSayHello(t *testing.T) {
	result := ch_1_helper.SayHello("Nurdin")
	assert.Equal(t, "Hello Nurdin", result) // (t, expected, actual)
}
