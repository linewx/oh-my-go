package advanced

import (
	"fmt"
	"testing"
	"time"
)

//1.7 defer function
func doDefer() {
	defer trace()()
	time.Sleep(5 * time.Second)
}

//1.7.1 trace func
func trace() func() {
	startTime := time.Now()
	return func() { fmt.Print("it cost ", time.Since(startTime)) }
}


func TestDeferFunc(t *testing.T) {
	doDefer()
}

