package method

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// 运行时获取错误信息
	errorOne()
}

// 错误
// Go语言中使用error值来表示错误状态
// 与fmt.Stringer()类似，error类型是一个内建接口
// 与fmt.Stringer()类型，fmt在打印值时也会满足error
// 通常函数会返回一个error值，函数调用者需要判断这个错误是否等于nil，以此来进行错误处理
// error为nil时表示函数调用成功，非nil时表示函数调用失败
type SunshineError struct {
	When time.Time
	Why  string
}

// 需要实现内建接口，才能对错误信息进行输出
func (sunshineError *SunshineError) Error() string {
	return fmt.Sprintf("at %v, %s", sunshineError.When, sunshineError.Why)
}

// 调用产生一个错误
func run() error {
	return &SunshineError{When: time.Now(), Why: "类型转换错误"}
}

// 运行获得错误
func errorOne() {
	if e := run(); e != nil {
		fmt.Println(e)
	}
}

// 练习：错误
// 让Sqrt函数返回error信息
// Sqrt函数在获取到一个非负数的值时，需要返回非nil的错误值，复数同样也不被支持
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number:%v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}
