package function

import "fmt"

// Go的基本类型
// bool
// string

// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64 uintptr
// uintptr：无符号整型指针，在32位系统下是4字节，在64位系统下8字节

// byte(uint8别名)
// rune(int32别名，表示一个Unicode码点)

// float32 float64
// complex64 complex128

// int uint uintptr在32位系统上通常为32bit，在64位系统上为64bit
// 当需要整整数值时应使用int类型
// 除非有特殊理由时，则使用固定大小或无符号的整数类型

func main() {
	// bool类型声明
	boolAssign()
	// int类型声明
	intAssign()
	// 零值
	zeroAssign()
	// 类型转换
	cast()
	// 类型推导
	typeInference()
}

func boolAssign() {
	var boolOne bool = false
	var boolTwo bool = true
	fmt.Println(boolOne, boolTwo)
}

func intAssign() {
	var intOne int = 64
	var intTwo int8 = 128
	var intThree int16 = 32768
	var intFour int32 = 2147483648
	var intFive int64 = 9223372036854775808
	fmt.Println(intOne, intTwo, intThree, intFour, intFive)
	var uintOne uint = 64
	var uintTwo uint8 = 256
	var uintThree uint16 = 65536
	var uintFour uint32 = 4294967296
	var uintFive uint64 = 18446744073709551616
	fmt.Println(uintOne, uintTwo, uintThree, uintFour, uintFive)
}

// 变量声明也可以"分组"成为一个语法块
var (
	UIntOne   uint32  = 64
	boolOne           = true
	stringOne string  = "Sunshine"
	floatOne  float32 = 5.4
	floatTwo  float64 = 123.6555
)

// 零值
// 没有明确初始值的变量会被赋予"零值"
// 数值类型的零值为0
// 布尔类型的零值为false
// 字符串的零值为""
func zeroAssign() {
	var i int
	var b bool
	var c string
	fmt.Println(i, b, c)
}

// 类型转换
// 表达式T(v)将值v转换为T类型
// Go在不同类型的项之间赋值时需要显式转换
func cast() {
	var i int = 64
	var f float32 = float32(i)
	var u uint = uint(f)
	fmt.Println(u)
}

// 类型推导
// 类型推导属于高级语言共有的特性
// 在声明一个变量时，不需要指定其类型(可用于:=和var声明的变量)
// 变量的类型由右值推导而出
// 右值声明为新的类型时，新变量的类型与其相同
func typeInference() {
	var x = 32
	y := x
	fmt.Println(y)
}
