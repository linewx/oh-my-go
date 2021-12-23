package advanced

import (
	"fmt"
	"testing"
)

//迭代变量
func iterVar() {
	var funclist []func()
	for index, _ := range [10]int{} {
		//defer trace2(index)()
		funclist = append(funclist, func() {
			fmt.Println(index)
		})
	}

	for _, theFunc := range funclist {
		theFunc()
	}
}

//迭代变量
func iterVarFix1() {
	var funclist []func()
	for index, _ := range [10]int{} {
		//defer trace2(index)()
		funclist = append(funclist, func(x int) func() {
			return func() {
				fmt.Println(x)
			}
		}(index))
	}

	for _, theFunc := range funclist {
		theFunc()
	}
}

func iterVarFix2() {
	var funclist []func()
	for index, _ := range [10]int{} {
		theIndex := index
		funclist = append(funclist, func() {
			fmt.Println(theIndex)
		})
	}

	for _, theFunc := range funclist {
		theFunc()
	}
}

func iterVarWithDefer() {
	for index, _ := range [10]int{} {
		defer func() { fmt.Println(index) }()
	}

}

func iterVarWithDeferFix() {
	for index, _ := range [10]int{} {
		theIndex := index
		defer func() { fmt.Println(theIndex) }()
	}
}

func iterVarWithDeferFix2() {
	for index, _ := range [10]int{} {
		defer trace2(index)
	}
}


func trace2(x int) {
	fmt.Println(x)
}

func TestIterVar(t *testing.T) {
	fmt.Println("wrong iter var ...")
	iterVar()
	fmt.Println("iter var fix1...")
	iterVarFix1()
	fmt.Println("iter var fix2...")
	iterVarFix2()
	fmt.Println("wrong iter var with defer...")
	iterVarWithDefer()
	fmt.Println("iter var with defer fix ...")
	iterVarWithDeferFix()

	fmt.Println("iter var with defer fix2 ...")
	iterVarWithDeferFix2()


}
