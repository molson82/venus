package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Simple : default true is true test
func TestSimple(t *testing.T) {
	assert.True(t, true, "True is true!")
}
