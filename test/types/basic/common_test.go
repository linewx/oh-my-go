package basic

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestCommon(t *testing.T) {

	//1.创建： 普通创建
	var (
		width,height = 600,320
		cells = 100
		angle = math.Pi/6
	)

	assert.Equal(t, 600, width)
	assert.Equal(t, 320, height)
	assert.Equal(t, 100,  cells)
	assert.Equal(t, math.Pi/6, angle )


}
