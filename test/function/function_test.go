package function

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//1. 创建
//1.1 normal function
func add(x int, y int) int {
	return x + y
}

//1.2 没有返回值
func hello(p *string) {
	*p = "hello " + *p
}

//1.3 命名返回值
func add2(x int, y int) (result int) {
	result = x + y
	return //省略了返回值
}

//1.4 函数参数
func accept(content string, visit func(string) string) string {
	return visit(content)
}

//1.5 可变参数
func addn(vars ...int) int {
	result := 0
	for _, value := range vars {
		result += value
	}
	return result
}

//1.6 多返回值
func stats(x int, y int) (sum int, mean float64) {
	sum = x + y
	mean = (float64)(x+y) / 2

	return
}

func TestFunction(t *testing.T) {
	assert.Equal(t, 3, add(1, 2))

	s1 := "Jack"
	hello(&s1)
	assert.Equal(t, "hello Jack", s1)

	assert.Equal(t, 3, add2(1, 2))

	//匿名函数
	assert.Equal(t, "hello Jack", accept("Jack", func(target string) string {
		return "hello " + target
	}))

	assert.Equal(t, 10, addn(1, 2, 3, 4))

	sum, mean := stats(1, 2)
	assert.Equal(t, 3, sum)
	assert.Equal(t, 1.5, mean)

}

