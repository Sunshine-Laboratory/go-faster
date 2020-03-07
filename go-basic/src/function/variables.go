package function

import "fmt"

// var用于声明一个变量
// 声明格式：参数名称 参数类型
// var可以用于包级或函数级别
var language string

func main() {
	// 函数级变量
	var faction int
	fmt.Println(faction, language)
}

// 赋给变量初始值
func assign() {
	var faction = "C++"
	var a, b, c = false, 1, 'a'
	fmt.Println(faction, a, b, c)
}

// 短变量声明
// ":="可以在类型明确的地方代替var声明
// 函数外的每个语句必须以关键词开始(var, func等)，因此":="不能再函数外使用
func shortAssign() {
	a := 3
	b := 'b'
	c := true
	d, e, f := 'd', false, "Sunshine"
	fmt.Println(a, b, c, d, e, f)
}
