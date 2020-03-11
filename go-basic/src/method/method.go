package method

import (
	"fmt"
	"strings"
)

func main() {
	// 简单的方法声明
	sunshine := Sunshine{1, "Sundial Dreams"}
	fmt.Println(sunshine.methodOne())
	// 函数和方法的对比
	fmt.Println(functionOne(sunshine))
	// string类型的方法
	name := Name("sunshine")
	fmt.Println(name.methodTwo())
	// 指针接收者
	sunshine.methodThree()
	fmt.Println(sunshine.methodOne())
	// 指针接收者和函数等同
	functionTwo(&sunshine)
	fmt.Println(functionOne(sunshine))
}

// Go语言中没有类这个定义
// 你可以通过结构体类型定义方法
// 方法是一类带特殊"接收者"参数的函数
// 方法接收者在它自己的参数列表内，位于func关键字和方法名之间
type Sunshine struct {
	No   int
	Name string
}

// 方法methodOne拥有一个名称为sunshine，类型为Sunshine结构体的接收者
func (sunshine Sunshine) methodOne() string {
	return sunshine.Name
}

// 方法即函数
// 方法仅仅是一个带有接收者的函数
// 下面的写法就是一个函数
func functionOne(sunshine Sunshine) string {
	return sunshine.Name
}

// 也可以为非结构体类型声明方法
// 比如为一个string类型的声明方法
type Name string

func (name Name) methodTwo() string {
	return strings.ToUpper(string(name))
}

// 指针接收者
// 也可以为指针接收者声明方法
// 针对于T类型，接收者的类型可以用*T的文法（T不能是类*int指针类型）
// 指针接收者的方法可以修改接收者指向的值，由于方法经常需要修改它的接收者，指针接收者比值接收者更为常用
func (sunshine *Sunshine) methodThree() {
	sunshine.No = 2
	sunshine.Name = "SunshineInterface"
}

// 指针与函数
// 继续把指针接受的方法改为函数
func functionTwo(sunshine *Sunshine) {
	sunshine.No = 2
	sunshine.Name = "SunshineInterface"
}

// 方法与指针重定向
// 比较两种类型的程序
// 带指针参数的函数必须接受一个指针
// 指针接收者则既可以是值，也可以是个指针
// 这是因为methodThree()方法有一个指针接收者，Go会将语句sunshine.methodThree()解释为(&sunshine).methodThree()

// 同样的事情也可以发生在相反的方向
// 接受一个值作为参数的函数必须接受一个指定类型的值
// 也就是函数的参数类型时固定的，是string类型，就不应该传入一个指针类型
// 而以值为接收者的方法被调用时，接收者既能为值，也可以为指针
// 也是是说既可以使用sunshine.methodThree()调用，也可以使用(&sunshine).methodThree()调用
// Go语言会自动将指针与真指针类型之间进行相互转换
func methodFour() {
	sunshine := Sunshine{
		No:   1,
		Name: "Sundial Dreams",
	}

	p := &sunshine
	fmt.Println(p.methodOne())
	fmt.Println(sunshine.methodOne())
}

// 选择值或指针作为接收者
// 使用指针接收者的原因：；
// 1. 方法能够修改接收者指向的值
// 2. 可以避免在每次调用方法时复制该值，若值的类型为大型结构体，使用方法会更加的高效
// 通常来说，所有给定类型的方法都应该具有值或指针接收者，但并不应该二者混用
