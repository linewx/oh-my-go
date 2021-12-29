package aggregate

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArray(t *testing.T) {
	//1. 创建
	var array1 [3]int // 只声明，隐试初始化
	assert.Equal(t, 3, len(array1))
	assert.Equal(t, 0, array1[0])

	array2 := [3]int{} //用字面量声明
	assert.Equal(t, 3, len(array2))
	assert.Equal(t, 0, array2[0])
	assert.Equal(t, 0, array2[2])

	array3 := [3]int{1, 2} //初始化个别元素
	assert.Equal(t, 3, len(array3))
	assert.Equal(t, 1, array3[0])
	assert.Equal(t, 2, array3[1])
	assert.Equal(t, 0, array3[2])

	array4 := [3]int{1:2} //按照index 创建
	assert.Equal(t, 2, array4[1])
	assert.Equal(t, 0, array4[2])

	array5 := [...]int{99: -1} //长度由元素决定
	assert.Equal(t, -1, array5[99])
	assert.Equal(t, 0, array5[98])

	//2. 使用
	//2.1 for loop
	for index,value:= range array5 {
		fmt.Println(index, value)
	}

	//2.2 比较
	a1 := []int{1,2,3}
	a2 := []int{1,2,3}
	assert.Equal(t, a1, a2) //值对比，非指针对比

}
