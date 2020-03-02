package flowcontrolloop

import "fmt"

func main() {
	// 声明一个指针
	pointerOne()
	// 指向内存地址的指针
	pointerTwo()
	// 指向底层值的指针
	pointerThree()
	// 通过指针修改原始值
	pointerFour()
}

// Go指针
// Go语言拥有指针，指针保存了值的内存地址
// 类型"*T"是指向T类型的指针，其零值为nil
func pointerOne() {
	var pointerOne *int
	fmt.Println(pointerOne)
}

// &操作符会生成一个指向其操作数的指针
func pointerTwo() {
	i := 64
	pointerTwo := &i
	fmt.Println(pointerTwo)
}

// *操作符表示指针指向的底层值
func pointerThree() {
	i := 64
	pointerThree := &i
	fmt.Println(*pointerThree)
}

// 通过指针设置值
// 这就是所说的间接引用或者重定向
// 与C不同的是，Go没有指针运算
func pointerFour() {
	i := 64
	pointerFour := &i
	*pointerFour = 128
	fmt.Println(i)
}
