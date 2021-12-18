package aggregate

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestSlice(t *testing.T) {
	/* 基本
	a. slice is reference type, which means function can modify the slice out of the function
	 */


	//1. 创建
	var sa11[]string //声明类型
	assert.Equal(t, 0, len(sa11))
	assert.Equal(t, 0, cap(sa11))

	sa12 := []string{"Apple", "Orange", "Banana"} //字面量
	assert.Equal(t, 3, len(sa12))
	assert.Equal(t, 3, cap(sa12))


	sa13 := make([]string, 10) //用make 创建
	assert.Equal(t, 10, len(sa13))
	assert.Equal(t, 10, cap(sa13))

	sa14 := make([]string, 10, 20)
	assert.Equal(t, 10, len(sa14))
	assert.Equal(t, 20, cap(sa14))

	s15 := [...]string{"apple", "banana", "orange", "grape", "mongo"}
	sa15 := s15[1:3]
	assert.Equal(t, 2, len(sa15))
	assert.Equal(t, 4, cap(sa15))

	//2. 使用




	//3. one more thing
	//3.1 sa1 is array, but sa1[:] is slice
	sa31 := [...]string{"Apple", "Orange", "Banana"}
	assert.Equal(t, "[3]string", reflect.TypeOf(sa31).String()) //sa3 is array
	assert.Equal(t, "[]string", reflect.TypeOf(sa31[:]).String()) //sa3[:] is slice
	assert.Equal(t, "[]string", reflect.TypeOf(sa31[1:2]).String()) //sa3[1:2] is slice

	//3.2 slice manipulation on string is string
	sa32 := "hello world"
	assert.Equal(t, "string", reflect.TypeOf(sa32).String()) //sa32 is string
	assert.Equal(t, "string", reflect.TypeOf(sa32[:]).String()) //sa32[:] is string too



}
