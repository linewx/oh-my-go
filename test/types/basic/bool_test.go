package basic

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestBool(t *testing.T) {

	// 创建： 普通创建
	var b1 bool
	assert.Equal(t, false, b1)

	b2 := true //字面量
	assert.Equal(t, true, b2)
	assert.Equal(t, "bool", reflect.TypeOf(b2).String())

	//字面量类型
	assert.Equal(t, "bool", reflect.TypeOf(false).String())





}