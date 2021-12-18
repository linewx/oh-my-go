package basic

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestFloat(t *testing.T) {

	// 创建： 普通创建
	var a int
	assert.Equal(t, 0, a)

	var f32 float32 = 0.1
	var f64 float64 = 0.1
	assert.True(t, f32 == 0.1)
	assert.True(t, f64 == 0.1)

	assert.NotEqual(t, 0.1, float64(f32)) //todo !!! ??? float64(f32) is 0.10000000149011612
	// assert.False(t, f64 == f32) mismatched types int64 and int32

	// 字面量类型
	assert.Equal(t, "float64", reflect.TypeOf(0.1).String())
}