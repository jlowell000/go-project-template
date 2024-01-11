package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_foobar(t *testing.T) {
	expected := "Alas, there's no code"
	actual := foobar()
	assert.Equal(t, expected, actual, "compare result")
}
