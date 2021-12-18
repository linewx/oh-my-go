package basic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConst(t *testing.T) {
	const a int = 1
	assert.Equal(t, 1, a, "a should 1")

	// a = 2, compile error

	const (
		MON int = iota + 1
		TUS
		WEZ
		THR
		FRI
		SAT
		SUN
	)
	assert.Equal(t, MON, 1, "mon should be 1")
	assert.Equal(t, SUN, 7, "sun should be 1")

	const (
		apple int = 1 << iota
		banana
		orange
	)

	assert.Equal(t, apple, 1, "")
	assert.Equal(t, banana, 2, "")
	assert.Equal(t, orange, 4, "")

	const (
		a1,b1 int = iota, iota + 1
		a2,b2
		a3,b3
	)

	assert.Equal(t, a1, 0)
	assert.Equal(t, a2, 1)
	assert.Equal(t, a3, 2)
	assert.Equal(t, b1, 1)
	assert.Equal(t, b2, 2)
	assert.Equal(t, b3, 3)

	//无类型常量
	c1 := 1 //无类型常量
	assert.Equal(t, c1, 1)

	var c2 float32 = 'a' //无类型常量
	assert.NotEqual(t, c2, 'a')
}