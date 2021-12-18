package basic

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestComplex(t *testing.T) {

	//1.创建： 普通创建
	var c1 complex128
	var c2 complex64

	c3 := 1 + 1i //字面量

	assert.Equal(t, "complex128", reflect.TypeOf(c1).String())
	assert.Equal(t, "complex64", reflect.TypeOf(c2).String())

	assert.Equal(t, "complex128", reflect.TypeOf(c3).String()) // 字面量默认类型

	// assert.False(t, a64 == a32) mismatched types int64 and int32

	//字面量类型
	assert.Equal(t, "complex128", reflect.TypeOf(1 + 1i).String())


}