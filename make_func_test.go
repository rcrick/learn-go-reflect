package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeFunc(t *testing.T) {
	// return value type shoud above agrument value type
	var intSum func(int, int) int64
	var floatSum func(float32, float32) float64
	var stringSum func(string, string) string

	MakeSum(&intSum)
	MakeSum(&floatSum)
	MakeSum(&stringSum)

	assert.Equal(t, int64(3), intSum(1, 2))
	assert.Equal(t, float64(float32(1.2)+float32(1.2)), floatSum(1.2, 1.2))
	assert.Equal(t, "12", stringSum("1", "2"))
}
