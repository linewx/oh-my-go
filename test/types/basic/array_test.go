package basic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArray(t *testing.T) {
	var array1[3]int // 只声明，隐试初始化
	assert.Equal(t, 3, len(array1), "the length of array1 should be array1")
	assert.Equal(t, 0, array1[0], "the first element of array1 should be 0")

	array2 := [3]int{} //用字面量声明
	assert.Equal(t, 3, len(array2), "the length of array should be array1")
	assert.Equal(t, 0, array2[0], "the first element of array should be 0")
	assert.Equal(t, 0, array2[2], "the first element of array should be 0")

	array3 := [3]int{1,2}
	assert.Equal(t, 3, len(array2), "the length of array should be array1")
	assert.Equal(t, 0, array3[0], "the first element of array should be 0")
	assert.Equal(t, 0, array3[2], "the first element of array should be 0")

}