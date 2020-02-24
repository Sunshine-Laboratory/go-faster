package flowcontrolloop

import "fmt"

func main() {
	// defer调用
	deferOne()
	// defer的调用顺序
	deferStoredInStack()
}

// defer语句
// defer语句会将函数推迟到外层函数返回之后执行
// 推迟调用的函数其参数会立即求值，但直到外层函数返回之前该函数都不会被调用
func deferOne() {
	defer fmt.Println("Sunshine")
	fmt.Print("Hello, ")
}

// defer栈
// 推迟调用的函数调用会被压入到一个栈中，在外层函数返回时，被推迟的函数会按照后进先出的顺序调用
func deferStoredInStack() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("defer的调用顺序")
}
