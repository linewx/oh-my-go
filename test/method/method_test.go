package method

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Job struct {
	company string
	salary float64
}

type Address struct {
	city string
	district string
	detail string
}

type Person struct{
	name string
	mobile string
	age int
	Address
	job Job
}

//1. 创建method

//1.1 struct作为接受者
func (p Person) isAdult() bool{
	if p.age >= 18 {
		return true
	}else {
		return false
	}
}

//1.2 指针作为接受者
func (p *Person) isAdult2() bool{
	if p.age >= 18 {
		return true
	}else {
		return false
	}
}


//嵌套struct的方法
func (p *Job) getSalary() float64 {
	return p.salary
}

func (p *Address) getDistrict() string {
	return p.district
}

func TestMethod(t *testing.T) {
	jack := Person{name: "Jack", mobile: "181019320982", age: 19, Address: Address{city: "shanghai", district: "pudong", detail: "shijidadao"}, job: Job{company: "hp", salary: float64(10000)}}
	allen := Person{name: "Allen", mobile: "1810191282982", age: 17, Address: Address{city: "shanghai", district: "minhang", detail: "qibao"}}

	//2 如何使用
	assert.True(t, jack.isAdult()) // 2.1普通调用，调用者和被调用者都是struct
	assert.True(t, (&jack).isAdult()) // 2.2 调用者是指针，被调用者是struct

	assert.False(t, allen.isAdult2()) //2.3 调用者是struct,被调用者是pointer
	assert.False(t, (&allen).isAdult2()) // 调用者和被调用者都是pointer

	//嵌套struct的方法调用
	assert.Equal(t, float64(10000), jack.job.getSalary())

	//嵌套struct的匿名防范调用
	assert.Equal(t, "pudong", jack.getDistrict())

}





