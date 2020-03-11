package method

import "fmt"

func main() {
	// 接口定义及意义
	interfaceOne()
	// 空指针接口值
	interfaceNilOne()
	// 接口值为nil
	interfaceNilTwo()
	// 空接口
	interfaceNone()
	// 类型断言
	interfaceAssertion()
	// 类型选择
	interfaceSelect()
	// Stringer，toString()方法的实现
	stringerOne()
}

// 接口
// "接口类型"是一组方法签名定义的集合
// 接口类型的变量可以保存任何实现了这些方法的值
// 下面举个例子可以明白接口的意义
type SunshineInterface interface {
	SunshineMethod() string
}

type SunshineName string

func (name SunshineName) SunshineMethod() string {
	return string(name)
}

type SunshineStruct struct {
	No   int
	Name string
}

func (sunshineStruct *SunshineStruct) SunshineMethod() string {
	return sunshineStruct.Name
}

func interfaceOne() {
	var sunshineInterface SunshineInterface
	sunshineStruct := SunshineStruct{
		No:   1,
		Name: "Sundial Dreams",
	}
	sunshineInterface = &sunshineStruct
	fmt.Println(sunshineInterface.SunshineMethod())

	sunshineName := SunshineName("SunshineInterface")
	sunshineInterface = sunshineName

	fmt.Println(sunshineInterface.SunshineMethod())

	// 这里会报编译错误，因为string类型没有实现SunshineMethod()方法
	// sunshineString := "SunshineInterface"
	// sunshineInterface = sunshineString
}

// 接口与隐式实现
// 类型通过实现一个接口的所有方法来实现该接口
// 既然无需专门显式声明，也就没有"implements"关键字
// 隐式实现从接口的实现中解耦的定义，这样的接口实现可以出现在任何的包中，无需提前准备
// 因此，也就无需在每一个实现上增加新接口的名称，这样同时也鼓励了明确接口的定义
type sunshineInterface interface {
	M()
}

type sunshineStructOne struct {
	Name string
}

type sunshineName string

// 此方法表示使用sunshineStructOne类型的接收者实现了sunshineInterface接口，但无需显示声明已经实现
func (sunshineStruct *sunshineStructOne) M() {
	if sunshineStruct == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(sunshineStruct.Name)
}

func (name sunshineName) M() {
	fmt.Println("Sunshine")
}

// 接口值
// 接口也是值，也可以像其他值一样传递
// 接口值可以作为函数的参数或返回值
// 在内部，接口值可以看做包含值和具体类型的元组: (value, type)
// 接口值存储了一个具体底层类型的具体值
// 接口值调用方法时会执行其底层类型的同名方法
// 当然作为值，也存在为nil的可能性
// 底层值为nil的接口值，即便接口内的具体值为nil，方法仍然会被nil接收者调用
// 在一些语言中，这样的调用会触发一个空指针异常，但是在Go中通常会写一些优雅地方法进行处理
// 注意：保存了nil具体值的接口，其自身不为nil
func interfaceNilOne() {
	var sunshineInterface sunshineInterface

	var sunshineStruct *sunshineStructOne
	// 此时sunshineInterface底层值指向了nil
	sunshineInterface = sunshineStruct

	// 空指针调用
	describe(sunshineInterface)
	sunshineInterface.M()
	// 此时sunshineInterface底层值指向了堆中的对象
	sunshineInterface = &sunshineStructOne{
		Name: "Sunshine",
	}
	describe(sunshineInterface)
	sunshineInterface.M()
}

func describe(sunshineInterface sunshineInterface) {
	fmt.Printf("(%v, %T)\n", sunshineInterface, sunshineInterface)
}

// nil接口值
// 和底层值的概念不同，这里指的是，接口值为nil值
// 相当于在Java中，声明了一个IUserService接口，但是没有任何实现，所以在调用此接口的方法时，会抛出异常
// nil接口值既不保存值，也不保存类型
// 为nil接口调用方法会产生运行时错误，因为接口的元组并未包含能够指明该调用哪个具体的方法的类型
func interfaceNilTwo() {
	var sunshineInterface sunshineInterface
	describe(sunshineInterface)
	sunshineInterface.M()
}

// 没有指定方法的接口称之为"空接口"
// 空接口可以保存任何类型的值，因为每个类型都至少实现了零个方法
// 空接口被用来处理未知类型的方法
// 例如：fmt.Print()可以接受类型为interface{}的任意数量的参数
func interfaceNone() {
	var interfaceNone interface{}
	describeInterfaceNone(interfaceNone)

	interfaceNone = 128
	describeInterfaceNone(interfaceNone)

	interfaceNone = "Sunshine"
	describeInterfaceNone(interfaceNone)
}

func describeInterfaceNone(interfaceNone interface{}) {
	fmt.Printf("(%v, %T)", interfaceNone, interfaceNone)
}

// 类型断言
// "类型断言"提供了访问接口值底层具体值的方式
// 比如：i := sunshineInterface.(T)
// 语句断言接口值sunshineInterface存储了类型T，并将其底层类型为T的值赋予变量i
// 若sunshineInterface并未保存类型T的值，该语句就会触发一个恐慌
// 为了判断一个接口值是否保存了一个特定的类型，类型断言可以返回两个值：
// * 底层值
// * 报告断言是否成功的布尔值
// i, ok := sunshineInterface.(T)
// 若sunshineInterface保存了T类型的底层值，则i是底层值，ok为true
// 若sunshineInterface没有保存T类型的底层值，则i是T类型的零值，ok为false，程序并不会产生恐慌
// 注意：这种语法和读取一个映射的相似之处
func interfaceAssertion() {
	var sunshineInterface sunshineInterface
	sunshineInterface = &sunshineStructOne{Name: "Sunshine"}

	i, ok := sunshineInterface.(*sunshineStructOne)
	fmt.Println(i, ok)

	var interfaceAssertion interface{} = "Sunshine"
	j, ok := interfaceAssertion.(string)
	fmt.Println(j, ok)

	k, ok := interfaceAssertion.(int)
	fmt.Println(k, ok)
}

// 类型选择
// "类型选择"是按照一种顺序从几个类型断言中选择分支的结构
// 类型选择与一般的switch语句相似，不过类型选择中的case为类型(而非值)
// 它们针对给定接口值所存储的值的类型进行比较
// switch v := sunshineInterface.(type) {
//     case T:
//         // v的类型为T
//     case S:
//         // v的类型为S
//     default:
//         // 没有匹配，v与sunshineInterface的类型相同
// }
// 类型选择中的声明与类型断言sunshineInterface.(T)的语法相同，只是类型T换成了关键字type
// 如果接口底层值是T或S类型，则变量v会按照T或S类型进行存储
// 否则默认情况下，变量v和i的接口类型和值相同
func interfaceSelect() {
	var sunshineInterface sunshineInterface

	sunshineInterface = &sunshineStructOne{Name: "Sunshine"}

	switch i := sunshineInterface.(type) {
	case *sunshineStructOne:
		fmt.Println(i.Name)
	case sunshineName:
		fmt.Println(i)
	default:
		fmt.Println("Unknown Type")
	}
}

// Stringer
// fmt包中定义的Stringer是最普遍的接口之一
// Stringer可以是用字符串描述当前存储的类型信息
// fmt包及其他包使用此接口来打印值
// 相当于重写toString()用法
type Student struct {
	Name string
	Age  int
}

func (stu Student) String() string {
	return fmt.Sprintf("%v (%v years)", stu.Name, stu.Age)
}

func stringerOne() {
	a := Student{
		Name: "Sundial Dreams",
		Age:  27,
	}
	b := Student{
		Name: "Sunshine",
		Age:  5,
	}
	fmt.Println(a, b)
}

// 练习
// 通过让 IPAddr 类型实现 fmt.Stringer 来打印点号分隔的地址。
type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func stringerTest() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
