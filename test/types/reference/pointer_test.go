package aggregate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPointer(t *testing.T) {

	//2.2

	//1. 创建
	a := 1
	assert.Equal(t, 1, a)
	var p1 *int
	p2 := &a //通过&取地址操作
	p3 := new(int)

	assert.True(t, nil == p1)
	assert.Equal(t, (*int)(nil), p1) // p1 is nil

	assert.Equal(t, p2, &a)
	assert.False(t, nil == p3) //p3 is not nil
	assert.Equal(t, 0, *p3)
}
