package flowcontrolloop

import (
	"fmt"
	"strings"
)

func main() {
	// 数组声明
	arrayOne()
	// 简单的切片使用
	sliceOne()
	// 更改切片数据带来的连锁反应
	sliceTwo()
	// 切片的文法
	sliceThree()
	// 切片的默认行为
	sliceFour()
	// 重新切片
	sliceFive()
	// nil切片
	sliceNil()
	// 使用make制作数组
	sliceMake()
	// 切片中的切片
	sliceWithSlice()
	// 切片中追加元素
	sliceAppend()
	// 遍历切片
	sliceRangeOne()
	// 带有占位符的切片遍历
	sliceRangeTwo()
	// 仅有索引的切片遍历
	sliceRangeThree()
}

// 数组
// 类型[n]T表示拥有n个T类型的值的数组
var a1 [10]int
var arrays = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

// 数组长度是其类型的一部分，因此数组不能改变其大小
// Go语言提供了更加便利的方式使用数组
func arrayOne() {
	var stringArray [10]string
	stringArray[0] = "Sundial"
	stringArray[1] = "Dreams"
	stringArray[2] = "Sunshine"
	fmt.Println(stringArray)
	intArray := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(intArray)
}

// 切片
// 每个数组的大小都是固定的
// 切面为数组元素提供动态大小的、灵活的视角
// 切片比数组更常用
// "[]T"表示为元素类型为T的切片
// 切片通常由两个上下标来界定：上界和下界，二者以":"分隔，比如a[low : high]
// 它包括的元素是一个左闭右开的区间，即包括第一个元素，但不包括最后一个元素
// 比如a[0, 4]视为包含了数组索引下标为0, 1, 2, 3的四个元素
func sliceOne() {
	sliceOne := arrays[1:4]
	fmt.Println(sliceOne)
}

// 切片就像是数组的引用
// 切片不存储任何数据，它仅是描述了底层数组的一段
// 更改切片的元素的同时也会修改原始数组中的值
// 与它共享底层数组的切片也会进行修改
func sliceTwo() {
	sliceTwo := arrays[1:4]
	sliceThree := arrays[2:3]
	sliceFour := arrays[3:5]
	sliceTwo[2] = 666
	fmt.Println(sliceTwo)
	fmt.Println(sliceThree)
	fmt.Println(sliceFour)
	fmt.Println(arrays)
}

// 切片文法
// 类似于没有长度的数组文法
// 比如[3]bool{true, true, true}是一个数组文法
// 而下面的语句会创建一个和它一样的数组，以及一个引用了此数组的切片
// []bool{true, true, true}
func sliceThree() {
	sliceFive := []int{1, 2, 3, 4, 5}
	sliceSix := []struct {
		No   int
		Name string
	}{
		{1, "Sundial Dreams"},
		{2, "Sunshine"},
	}
	fmt.Println(sliceFive)
	fmt.Println(sliceSix)
}

// 切片的默认行为
// 可以使用切片的默认行为忽略上下边界
// 切片默认的下界为0，上界为数组的长度
// 注意，不能超过上界与下界，否则会报错
func sliceFour() {
	sliceSeven := arrays[:10]
	fmt.Println(sliceSeven)
	sliceSeven = arrays[2:]
	fmt.Println(sliceSeven)
	sliceSeven = arrays[:]
	fmt.Println(sliceSeven)
}

// 切片的长度和容量
// 切面拥有长度和容量
// 切片的长度就是它所包含的元素个数
// 切片的容量就是从它的一个元素开始，到其底层数组元素末尾的个数
// 切片的长度可以通过表达式len(s)和cap(s)来获取
// 可以通过重新切片扩展一个切片，给它提供足够的容量
func sliceFive() {
	sliceEight := arrays[:8]
	fmt.Println(sliceEight)
	sliceEight = sliceEight[2:6]
	fmt.Println(sliceEight)
	sliceEight = sliceEight[:2]
	fmt.Println(sliceEight)
}

// nil切片
// 切片的零值是nil
// nil切片的长度和容量为0，并且没有底层数组
func sliceNil() {
	var sliceNine []int
	fmt.Println(sliceNine, len(sliceNine), cap(sliceNine))
	if sliceNine == nil {
		fmt.Println("nil slice")
	}
}

// 使用make制作切片
// 切片可以使用内置函数make制作，同时也是创建动态数组的方式
// make函数会分配一个元素为零值的数组并返回一个引用了它的切片
func sliceMake() {
	sliceTen := make([]int, 5)
	fmt.Println(sliceTen)
	sliceEleven := make([]int, 5, 5)
	fmt.Println(sliceEleven)
	sliceEleven = sliceEleven[:cap(sliceEleven)]
	fmt.Println(sliceEleven)
	sliceEleven = sliceEleven[1:]
	fmt.Println(sliceEleven)
}

// 切片可以包括任何类型，甚至可以包括其它的切片
func sliceWithSlice() {
	sliceTwelve := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	sliceTwelve[0][0] = "X"
	sliceTwelve[2][2] = "O"
	sliceTwelve[1][2] = "X"
	sliceTwelve[1][0] = "O"
	sliceTwelve[0][2] = "X"
	for i := 0; i < len(sliceTwelve); i++ {
		fmt.Printf("%s\n", strings.Join(sliceTwelve[i], ""))
	}
}

// 向切片中追加元素
// 使用内建函数append()可以向切片中追加元素
// func append(s []T, vs ...T) []T
// s: 元素类型为T的切片
// vs: 需要追加到切片中的元素类型为T的值
// append()函数的结果是一个包含原切片所有元素加上新添加元素的切片
// 当给定的切片的底层数组容量不足时，会分配一个容量更大的数组，返回的切片会指向这个新分配的数组
// 切片的扩容不是和你的追加的元素个数有关，而是成2的倍数增长，以免发生多次扩容的情形
func sliceAppend() {
	sliceThirteen := []int{1, 2, 3, 0, 0, 0}
	fmt.Println(sliceThirteen)
	sliceThirteen = append(sliceThirteen, 4, 5, 6)
	fmt.Println(sliceThirteen)
}

// 切片的遍历
// for循环的range模式可以遍历切片
// for循环遍历切片会返回两个值
// * 当前元素的索引
// * 当前元素的副本
func sliceRangeOne() {
	sliceFourteen := arrays[:]
	for i, i2 := range sliceFourteen {
		fmt.Printf("索引: %d, 元素: %d", i, i2)
	}
}

// 可以使对索引和元素副本赋以"_"来进行忽略
func sliceRangeTwo() {
	sliceFifteen := arrays[:]
	for _, i2 := range sliceFifteen {
		fmt.Printf("元素: %d", i2)
	}
}

// 如果只需要索引，可以直接忽略元素副本
func sliceRangeThree() {
	sliceSixteen := arrays[:]
	for i := range sliceSixteen {
		fmt.Printf("索引: %d", i)
	}
}
