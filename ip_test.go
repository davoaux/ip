package ip_test

import (
	"davoaux/ip"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	assert := assert.New(t)

	assert.True(ip.IsValid("1.2.3.4"))
	assert.True(ip.IsValid("123.45.67.89"))

	assert.False(ip.IsValid("1.2.3"))
	assert.False(ip.IsValid("1.2.3.4.5"))
	assert.False(ip.IsValid("123.456.78.90"))
	assert.False(ip.IsValid("123.045.067.089"))
}
