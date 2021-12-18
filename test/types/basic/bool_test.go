package basic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBool(t *testing.T) {

	// 创建： 普通创建
	var b1 bool
	assert.Equal(t, false, b1)

	b2 := true
	assert.Equal(t, true, b2)
}