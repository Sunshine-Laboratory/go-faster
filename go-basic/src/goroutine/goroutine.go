package goroutine

import (
	"fmt"
	"golang.org/x/tour/tree"
	"time"
)

func main() {
	// 使用goroutine
	goroutineOne()
	// 阻塞式信道
	goroutineTwo()
	// 非阻塞式信道
	goroutineThree()
	// 在信道中使用range和close的用法
	goroutineFour()
}

// goroutine是Go运行时管理的轻量级线程
// go f(x, y, z)
// 会启动一个新的Go程并执行
// f(x, y, z)
// f, x, y, z的求值发生在当前的Go程中，而f的执行发生在新的Go程中
// Go程在相同的地址空间中运行，因此在访问共享的内存时必须进行同步
// sync包提供了同步的能力
func loop(s string) {
	time.Sleep(100 * time.Millisecond)
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func goroutineOne() {
	go loop("Sunshine")
	loop("The Go Faster's author: ")
}

// 信道
// 信道是带有类型的管道，你可以通过它并使用信道操作符"<-"来发送或接收值
// ch <- v 代表将v发送至信道ch
// v := <-ch 代表从ch接收值并赋给变量v
// 总结下就是，"<-"代表数据流向的方向
// 同映射或切片一样，信道在使用前必须创建
// ch := make(chan int)
// 默认情况下，发送和接收操作在另一端准备之前都会阻塞，这使得Go程可以在没有显式的锁或竞态变量的情况下进行同步
func sum(x []int, c chan int) {
	sum := 0
	for _, v := range x {
		sum += v
	}
	c <- sum
}

func goroutineTwo() {
	x := []int{1, 2, 3, 4, 5, 6}
	y := []int{6, 5, 4, 3, 2, 1}
	c := make(chan int)
	go sum(x, c)
	go sum(y, c)
	a, b := <-c, <-c
	fmt.Println(a + b)
}

// 带缓冲的信道
// 信道是可以带缓冲的，将缓冲长度作为第二个参数提供给make来初始化一个带缓冲的信道:
// ch := make(chan int, 100)
// 仅当信道的缓冲区被填满之后，向其发送的数据才会阻塞，当缓冲区为空时，接受方会阻塞
func goroutineThree() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// 	ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// range和close
// 发送者可以通过close关闭一个信道来表示没有需要发送的值
// 接收者可以通过为接收表达式分配第二个参数来测试信道是否被关闭
// 若没有值可以被接收，且信道被关闭，那么执行完
// v, ok := <-ch
// 之后，ok=false
// for循环
// for i := range c
// 会不断的从信道接收值，直到信道被关闭
// 注意：
// * 只有发送者能关闭信道，接收者不可以关闭信道，向一个已经关闭的信道发送数据会引起恐慌（panic）
// * 信道与文件不同，通常情况下无需关闭信道，只有在必须告诉接收者不再有需要发送的值时才有必要关闭，例如需要终止一个无限for循环
func fibonacciOne(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// 关闭信道传输
	close(c)
}

func goroutineFour() {
	c := make(chan int, 10)
	go fibonacciOne(10, c)
	for i := range c {
		// 信道关闭，自动退出循环
		fmt.Println(i)
	}
}

// select语句
// select语句可以使一个Go程可以等待多个通信操作
// select语句会阻塞到某个分支可以继续执行为止，这时就会执行该分支
// 当多个分支都准备好时，会随机选择一个执行
func fibonacciTwo(c chan int, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quit")
			return
		}
	}
}

func goroutineFive() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacciTwo(c, quit)
}

// 默认选择
// 当select中的其他分支都没有准备好时，default分支就会执行
// 为了在尝试发送或者接收时不发生阻塞，可以使用default分支
func goroutineSix() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("boom")
		default:
			fmt.Println("default")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// 练习：等价二叉查找树
// 不同的二叉树的叶节点上，可以保存相同的值序列
// 例如以下两个二叉树都保存了序列[1, 1, 2, 3, 5, 8, 13]
// 使用Go来实现检查两个二叉树是否保存了相同的序列
// 实现并测试Walk函数
// 函数tree.New(k)用于构造一个随机结构的已排序二叉查找树，它保存了值k, 2k, 3k, ...,10k
// 创建一个新的信道ch并且对其进行步进
// go Walk(tree.New(1), ch)
// 然否从信道中读取并打印10个值，应当是数字1, 2, 3, ..., 10
// 用Walk实现Same函数来检测ti和t2是否存储了相同的值
// 测试Same函数
// Same(tree.New(1), tree.New(1))应当返回true，而Same(tree.New(1), tree.New(2))应当返回false
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	if t.Right != nil {
		Walk(t.Right, ch)
	}
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}
	return true
}

func goroutineSeven() {
	fmt.Println(Same(tree.New(1), tree.New(2)))
	fmt.Println(Same(tree.New(1), tree.New(1)))
}
