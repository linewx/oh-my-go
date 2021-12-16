package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArray(t *testing.T) {
	var array1[3]int // 只声明，隐试初始化
	assert.Equal(t, 3, len(array1), "the length of array1 should be array1")
	assert.Equal(t, 0, array1[0], "the first element of array1 should be 0")

	fmt.Printf(cap(array1))


}