package flowcontrolloop

import "fmt"

func main() {
	// 声明一个简单的结构体
	structOne()
	// 使用"."来访问结构体字段
	structTwo()
	// 通过结构体指针修改初始值
	structThree()
	// 结构体的声明文法
	structFour()
}

// 结构体
// struct就是一组字段
type Sunshine struct {
	X    int
	Name string
}

func structOne() {
	fmt.Println(Sunshine{1, "Sundial Dreams"})
}

// 结构体字段通过"."来访问
func structTwo() {
	sunshine := Sunshine{1, "Sundial Dreams"}
	sunshine.X = 2
	sunshine.Name = "Sunshine"
	fmt.Println(sunshine)
}

// 结构体指针
// 结构体字段可以通过结构体指针来访问
// 结构体指针p，可以通过(*p).X来访问Sunshine结构体的X字段
// 这么写比较啰嗦
// 所以Go语言还允许隐式间接引用，直接写p.X也可以进行访问
func structThree() {
	sunshine := Sunshine{1, "Sundial Dreams"}
	pointerFive := &sunshine
	(*pointerFive).X = 2
	(*pointerFive).Name = "Sunshine"
	fmt.Println(sunshine)
	pointerFive.X = 3
	pointerFive.Name = "mysunshinedreams"
	fmt.Println(sunshine)
}

// 结构体文法
// 通过直接列出字段的值来创建一个结构体
// 使用"Name:"语法可以仅列出部分字段(字段名的顺序无关)
// 特殊的前缀"&"返回一个指向结构体的指针
func structFour() {
	sunshineOne := Sunshine{1, "Sundial Dreams"}
	sunshineTwo := Sunshine{Name: "Sunshine"}
	sunshineThree := Sunshine{}
	fmt.Println(sunshineOne)
	fmt.Println(sunshineTwo)
	fmt.Println(sunshineThree)
}
