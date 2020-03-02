package flowcontrolloop

import (
	"fmt"
	"strings"
)

func main() {
	// 简单的映射使用
	mapOne()
	// 映射的文法
	mapTwo()
	// 修改映射
	mapThree()
	// 返回给定字符串的字符个数
	wordCount("Sunshine")
}

// 映射
// key => value
// 映射的零值为nil，nil映射既没有键，也不能添加键
// make函数会返回给定类型的映射，并将其初始化备用
var mapsOne map[string]Sunshine

func mapOne() {
	mapsOne = make(map[string]Sunshine)
	mapsOne["Sundial Dreams"] = Sunshine{
		1,
		"Sunshine",
	}
	fmt.Println(mapsOne["Sundial Dreams"])
}

// 映射的文法
// 映射的文法与结构体相似，不过必须有键名
// 若顶级类型只是一个类型名，你可以在文法的元素中省略它
func mapTwo() {
	var mapsTwo = map[string]Sunshine{
		"Sundial Dreams": {
			X:    1,
			Name: "Sundial Dreams",
		},
		"Sunshine": {
			X:    2,
			Name: "Sunshine",
		},
	}
	fmt.Println(mapsTwo)
}

// 修改映射
func mapThree() {
	mapsThree := map[string]int{
		"Sundial Dreams": 1,
		"Sunshine":       2,
	}
	fmt.Println(mapsThree)
	mapsThree["Sunshine"] = 666
	fmt.Println(mapsThree)
	delete(mapsThree, "Sundial Dreams")
	fmt.Println(mapsThree)
	v, ok := mapsThree["Sunshine"]
	fmt.Println("The value is: ", v, "Present?", ok)
}

// 返回给定字符串的字符个数
func wordCount(s string) map[string]int {
	resultMap := make(map[string]int)
	stringArray := strings.Fields(s)
	for _, str := range stringArray {
		resultMap[str] = resultMap[str] + 1
	}
	return resultMap
}
