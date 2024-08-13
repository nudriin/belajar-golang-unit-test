package ch_6_table_test

import (
	"testing"

	"github.com/nudriin/belajar-golang-unit-test/ch_1_helper"
	"github.com/stretchr/testify/assert"
)

func TestTable(t *testing.T) {

	// Membuat table untuk test
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "SayHello(Nurdin)",
			request:  "Nurdin",
			expected: "Hello Nurdin",
		},
		{
			name:     "SayHello(Sasuke)",
			request:  "Sasuke",
			expected: "Hello Sasuke",
		},
		{
			name:     "SayHello(Naruto)",
			request:  "Naruto",
			expected: "Hello Naruto",
		},
	}

	// ! Membuat table test dengan data yang berbeda-beda
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ch_1_helper.SayHello(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}
