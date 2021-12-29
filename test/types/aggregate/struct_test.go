package aggregate

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)


func TestStruct(t *testing.T) {
	//1. 创建
	//1.1 创建Struct Type
	type Person struct {
		name   string
		addr   string
		age    int
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

	//2.2 通过指针调用
	p2p := &p2
	assert.Equal(t, "MiXiaoquan", p2p.name)
	assert.Equal(t, "MiXiaoquan", (*p2p).name)

	//3. one more thing
	//3.1 嵌套struct
	type Order struct {
		orderTime time.Time
		name      string
		category  string
		buyer     Person
	}

	order := Order{buyer: Person{name: "xiaomi", age: 8, addr: "minhang"}, name: "computer", category: "3c"}
	assert.Equal(t, "3c", order.category)
	assert.Equal(t, "computer", order.name)
	assert.Equal(t, 8, order.buyer.age)

	//3.2 匿名成员
	type Department struct {
		name, description string
		Person
	}

	department := Department{
		name:        "R&D",
		description: "development",
		Person: Person{name: "xiaomi", age: 8, addr: "minhang"},
	}

	assert.Equal(t, "R&D", department.name)
	assert.Equal(t, "development", department.description)
	assert.Equal(t, "minhang", department.addr) //直接调用了匿名成员的变量

}
