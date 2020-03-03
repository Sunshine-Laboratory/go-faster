package flowcontrolloop

import (
	"fmt"
	"math"
)

func main() {

}

// 函数值
// 函数也是值，可以像其他值一样传递
// 函数值可以用作函数的参数或返回值
// 意思很简单，函数可以作为参数，放入到另一个函数中
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// 声明一个函数为变量
func functionValueOne() {
	pythagorean := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(pythagorean(5, 12))
}

// 作为3，4计算的函数放入进来
func funtionalValueTwo() {
	pythagorean := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(compute(pythagorean))
}

// 传入一个函数
func functionalValueThree() {
	// 两个入参，x, y，求x^y（幂等运算）
	fmt.Println(compute(math.Pow))
}

// 闭包
// Go函数可以是一个闭包，闭包是一个函数值，它引用了函数体之外的变量
// 该函数可以访问并赋予引用变量的值
// 换句话说，该函数被这些变量绑定到一起
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// 函数adder返回了一个闭包，每个闭包绑带在其各自的sum变量上
func functionaValueFour() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}

// 使用闭包完成斐波那契数列
func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		temp := x
		x, y = y, x+y
		return temp
	}
}
