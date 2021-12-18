package basic

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
	"unicode"
)

func TestString(t *testing.T) {

	//1.创建： 普通创建
	var s1 string
	assert.Equal(t, "", s1)

	s2 := "hello"
	assert.Equal(t, "hello", s2)

	//2.使用

	//2.1 strings package, strings operation including contain, index, join ...
	s21 := "hello world"
	assert.True(t, strings.Contains(s21, "hello"))
	assert.Equal(t, 2, strings.Index(s21, "l"))
	assert.Equal(t, s21, strings.Join([]string{"hello", "world"}, " "))
	assert.True(t, strings.Compare("hello", "world") < 0)

	//2.2 unicode, operation on unicode code point
	assert.True(t, unicode.IsLower('c'))
	assert.True(t, unicode.IsDigit('1'))
	assert.Equal(t, 'C', unicode.ToUpper('c'))

	//2.3 strconv, do string conversion, like conversion between number/bool type and string
	result1, _ := strconv.ParseFloat("0.1", 64)
	assert.Equal(t, 0.1, result1)

	result2, _ := strconv.ParseFloat("0.1", 32)
	fmt.Println(result2) //surprising!!!, result is 0.10000000149011612

	assert.Equal(t, "true", strconv.FormatBool(true))

	//2.2

}
