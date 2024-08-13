package ch5subtest

import (
	"testing"

	"github.com/nudriin/belajar-golang-unit-test/ch_1_helper"
	"github.com/stretchr/testify/assert"
)

// Membuat sub test, mirip describe di jest
func TestSubTest(t *testing.T) {
	//! Menjalankan sub test
	t.Run("test 1", func(t *testing.T) {
		result := ch_1_helper.SayHello("Nurdin")
		assert.Equal(t, "Hello Nurdin", result)
	})
	t.Run("test 2", func(t *testing.T) {
		result := ch_1_helper.SayHello("Nurdin")
		assert.NotEqual(t, "Hello Nurdin", result, "Harus tidak sama")
	})
}
