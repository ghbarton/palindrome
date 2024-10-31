package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		name string
		data int
		exp  bool
	}{
		{
			name: "123",
			data: 123,
			exp:  false,
		},
		{
			name: "121",
			data: 121,
			exp:  true,
		},
		{
			name: "10",
			data: 10,
			exp:  false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			audit := isPalindrome(tc.data)
			assert.Equal(t, tc.exp, audit)
		})
	}
}
