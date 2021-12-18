package basic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt(t *testing.T) {

	// 创建： 普通创建
	var a int
	assert.Equal(t, 0, a)

	var a32 int32 = 1
	var a64 int64 = 1
	var r rune = 1
	assert.True(t, a32 == 1)
	assert.True(t, a64 == 1)
	assert.True(t, a64 > 0)
	assert.True(t, r == 1)

	// assert.False(t, a64 == a32) mismatched types int64 and int32

}