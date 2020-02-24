package function

import (
	"fmt"
)

func main() {
	fmt.Println(AddFunctionOne(123, 456))
	a, b := swap("Sundial", "Dreams")
	fmt.Println(a, b)
	fmt.Println(split(17))
}

// 声明一个函数
func AddFunctionOne(x int, y int) int {
	return x + y
}

// 参数类型相同，除最后一个参数，其他参数可以省略参数类型
func addFunctionTwo(x, y, z int) int {
	return x + y + z
}

// 多值返回
// 函数可以返回任意数量的返回值
func swap(x, y string) (string, string) {
	return y, x
}

// 命名返回值
// 命名的返回值会被定义为在函数顶部的变量
// 返回值的名称应具有一定的意义，它可以作为文档使用
// 无参return语句返回已命名的返回值，也就是"直接返回"
func split(sum int) (x, y int) {
	x = sum * 2 / 4
	y = (sum + 2) / 4
	return
}
