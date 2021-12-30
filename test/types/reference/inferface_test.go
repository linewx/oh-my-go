package aggregate

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"
)

//1. 创建
type Processor interface {
	process(data interface{}) interface{}
}


type AbstractStringProcessor interface {
	process(data interface{}) interface{}
}

type AbstractStringProcessor2 interface {
	process(string) interface{}
}

type StringProcessor struct {
}

type StringProcessor2 struct {

}

func (p StringProcessor2) process2(data interface{}) interface{} {
	return nil
}
//2. 使用

func (p StringProcessor) process(data interface{}) interface{} {

	switch x := data.(type) { //2.1 type switch 的用法
	case int:
		return fmt.Sprintf("%d", data)
	case bool:
		if x {
			return "TRUE"
		} else {
			return "FALSE"
		}
	case string:
		return x
	}

	return ""
}


//2.2 接受interface作为一个变量
type ProcessorManager struct {
	processors []Processor
}

func (p * ProcessorManager) addProcessor(processor Processor) {
	p.processors = append(p.processors, processor)
}


func (p *ProcessorManager) processAll (data interface{}) {
	for _,one := range p.processors {
		fmt.Println(one.process(data))
	}
}

func TestInterface(t *testing.T) {
	pm := ProcessorManager{}
	pm.addProcessor(StringProcessor{})
	pm.processAll(1)

	//2.3 type assertion
	var w io.Writer


	w = os.Stdout
	//a. w have only one Write method
	//w.Write


	//b. concrete type
	f := w.(*os.File)
	//b.1 f become os.File type, have many os.File method
	//f.Name()
	//f.Write()
	//f.Read()

	fmt.Println(reflect.TypeOf(f).Name())

	if _,ok := w.(*bytes.Buffer); ok {
		fmt.Println("success")
	}else {
		fmt.Println("fail")
	}

	//c. interface type
	rw := w.(io.ReadWriter)
	//c.1 only have method declared in ReadWriter
	// rw.Read()
	// rw.Write()
	reflect.ValueOf(rw).MethodByName("Write")



}
