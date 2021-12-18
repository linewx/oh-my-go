package aggregate

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {

	//1. 创建
	var m1 map[string]int
	assert.True(t, m1 == nil)

	m2 := map[string]int{"yuwen": 99, "english": 95, "math": 100, "tiyu": 0} //字面量
	assert.Equal(t, 99, m2["yuwen"])

	m3 := map[string]int{}
	assert.Equal(t, 0, len(m3))

	m4 := make(map[string]int)
	assert.True(t, m4 != nil)

	//2. 使用
	//2.1 access map
	m21 := map[string]int{"yuwen": 99, "english": 95, "math": 100, "tiyu": 0}
	value,ok := m21["math"]
	assert.True(t, ok)
	assert.Equal(t, 100, value) //access eleement

	value2,ok2 := m21["meishu"]
	assert.False(t, ok2)
	assert.Equal(t, 0, value2) //todo: 0 even element not is map

	//2.2 loop map
	m5 :=  map[string]int{"yuwen": 99, "english": 95, "math": 100, "tiyu": 0}
	for key,value := range m5 { //loop map
		assert.Equal(t, "string", reflect.TypeOf(key).String())
		assert.Equal(t, "int", reflect.TypeOf(value).String())
	}
}
