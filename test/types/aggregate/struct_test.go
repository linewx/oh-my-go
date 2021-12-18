package aggregate

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestStruct(t *testing.T) {
	//1. 创建
	//1.1 创建Struct Type
	type Person struct {
		name string
		addr string
		age int
		mobile string
	}

	//1.2 利用Struct来创建变量
	var p1 Person
	assert.Equal(t, "aggregate.Person", reflect.TypeOf(p1).String())

	p2 := Person{"MiXiaoquan", "shanghai", 8, "18100000000"}
	assert.Equal(t, "MiXiaoquan", p2.name)

	p3 := Person{name: "JiangXiaoYa", age: 7}
	assert.Equal(t, 7, p3.age)

	//2. 使用
	//2.1 通过变量直接调用
	assert.Equal(t, "MiXiaoquan", p2.name)

	p2p := &p2
	assert.Equal(t, "MiXiaoquan", p2p.name)
	assert.Equal(t, "MiXiaoquan", (*p2p).name)









}
