package basic

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestByte(t *testing.T) {

	//1.创建： 普通创建
	var b1 byte
	assert.Equal(t, byte(0), b1)

	s2 := "hello"
	b2 := []byte(s2)
	fmt.Println(b2)
	assert.Equal(t, 5, len(b2))

	assert.Equal(t, s2, string(b2))


	//2.使用

	//2.1 bytes package, strings operation including contain, index, join ...
	s21 := "hello world"
	b21 := []byte(s21)

	assert.True(t, bytes.Contains(b21, []byte("hello")))



}
