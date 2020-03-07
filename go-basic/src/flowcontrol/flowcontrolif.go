package flowcontrolloop

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	// if语句
	ifOne(666)
	// 变量声明的if语句
	ifTwo()
	// switch语句
	switchOne()
	// switch的求值顺序
	switchTwo()
	// 优化"if else if"语句
	switchThree()
}

// if语句也不需要"()"包住条件表达式，但是内容需要使用"{}"包住
func ifOne(x int) string {
	if x > 10 {
		return "Sundial Dreams"
	} else if x > 20 {
		return "mysunshinedreams"
	}
	return "Sunshine"
}

// if语句可以在条件表达式前执行一个简单的语句
// 语句声明变量的作用域仅在if语句内
func ifTwo() string {
	if x := rand.Int(); x > 100 {
		return "Sundial Dreams"
	}
	return "Sunshine"
}

// switch是对长"if, else if"语句的简化
// Go自动提供每个case后面所需的break语句
// switch的case无需为常量，且取值不必为整数
func switchOne() {
	fmt.Print("Go program runs on ")
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Linux")
	case "darwin":
		fmt.Println("macOS")
	default:
		fmt.Println("windows")
	}
}

// switch的求值顺序
// switch的case语句从上到下执行，直到匹配成功
func switchTwo() {
	today := time.Now().Weekday()
	switch time.Sunday {
	case today:
		fmt.Println("Sunday is today")
	case today + 1:
		fmt.Println("Sunday is tomorrow")
	case today + 2:
		fmt.Println("Sunday is int two days")
	default:
		fmt.Println("Emm....")
	}
}

// 没有条件的switch语句
// 没有条件的switch语句等同于switch true
// 可以将"if else if"语句条理更清晰
func switchThree() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 17:
		fmt.Println("Afternoon")
	case t.Hour() < 20:
		fmt.Println("Evening")
	default:
		fmt.Println("Midnight")
	}
}
