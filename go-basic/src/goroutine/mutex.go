package goroutine

import "sync"

func main() {

}

// sync.Mutex
// 我们已经看到信道非常适合在各个Go程之间通信
// 但是如果我们并不需要通信呢？比如说，我们只是想保证每一次只有一个Go程能够访问一个共享的变量，从而避免冲突？
// 这里大家就会想到互斥锁
// 互斥（mutual exclusion）
// Go标准库中提供了sync.Mutex互斥锁类型及其两个方法
// Lock Unlock
// 我们可以通过在代码前调用Lock方法，在使用完成之后调用Unlock方法，从而保证一段代码的互斥行为
// 我们也可以用defer语句来确保互斥锁一定会被解锁
type SunshineMutex struct {
	v   int
	mux sync.Mutex
}

func (sunshineMutex SunshineMutex) mutexOne() {
	sunshineMutex.mux.Lock()
	sunshineMutex.v++
	sunshineMutex.mux.Unlock()
}

func (sunshineMutex SunshineMutex) mutexTwo() {
	sunshineMutex.mux.Lock()
	defer sunshineMutex.mux.Unlock()
	sunshineMutex.v++
}
