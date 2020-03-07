package flowcontrolloop

import "fmt"

func main() {
	// 简单循环
	loopOne()
	// 没有初始化语句、后置语句循环
	loopTwo()
	// for循环的变种 => while循环
	loopThree()
	// 无限循环
	loopFour()
}

// Go只有一种循环结构：for循环
// for循环由三部分组成，使用分号隔开
// * 初始化语句：在每一次迭代前执行
// * 条件表达式：在每次迭代前求值
// * 后置语句：在每次迭代的结尾执行
// 初始化语句通常为一个短变量声明，该变量声明仅在for语句的作用域内可见
// 一旦条件表达式结果=false，迭代将会终止
// Go中的for循环的三个语句部分不会用"()"进行分隔
func loopOne() int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}

// for循环的初始化语句和后置语句是可选的
func loopTwo() int {
	index := 0
	sum := 0
	for index < 10 {
		sum += index
		index++
	}
	return sum
}

// 你也可以去掉分号，那么其实就是一个while循环了
func loopThree() (index, sum int) {
	for sum < 5050 {
		sum += index
		index++
	}
	return index, sum
}

// 如果没有任何条件语句，那么会是一个无限循环
func loopFour() {
	for {
		fmt.Println("无限循环")
	}
}
