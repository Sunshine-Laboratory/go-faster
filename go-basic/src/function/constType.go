package function

import "fmt"

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

// 常量
// 常量使用const声明
// 常量类型可以是字符、字符串、布尔值或数值
// 常量不能使用":="语法声明
const IntOne = 32
const IntTwo int = 64
const ByteOne byte = 'a'
const BoolTwo bool = true
const AuthorName string = "Sunshine"

// 数值常量
// 数值常量是高精度的"值"
// 一个未指定类型的常量由上下文决定其类型
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}
