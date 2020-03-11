package method

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// 读取字节流
	readerOne()
}

// io包指定了io.Reader接口，它表示从数据流的末尾进行读取
// Go标准库包含了该接口的许多实现，包括文件、网络连接，压缩和加密等
// io.Reader接口中有一个Read方法：
// func (T) Read(b []byte) (n int, err error)
// Read用数据填充给定的字节切片并返回填充的字节数和错误值
// 在遇到字节流的末尾时，它会返回一个io.EOF错误
func readerOne() {
	r := strings.NewReader("Sundial Dreams")
	b := make([]byte, 8)
	for {
		// 从Reader中读取字节流，每次读取8个字节
		n, err := r.Read(b)
		// 打印返回的信息
		fmt.Printf("n = %v, error = %v, b = %v\n", n, err, b)
		// 打印读取的字节数组信息
		fmt.Printf("b[:n] = %q\n", b[:n])
		// 遇到文件终止符，终止读取
		if err == io.EOF {
			break
		}
	}
}

// 练习：Reader
// 实现一个Reader类型，它产生一个ASCII字符的'A'的无限流
type SunshineReader struct {
}

func (sunshineReader SunshineReader) Reader(b []byte) (n int, err error) {
	b[0] = 'A'
	return 1, nil
}

// 练习：rot13Reader
// 有一种常见的模式是将io.Reader包装成另一个io.Reader，然后通过某种方式修改其数据流
// 例如，gzip.NewReader函数接受一个io.Reader（已压缩的数据流）并返回一个同样实现io.Reader的*gzip.Reader（解压后的数据流）
// 编写一个实现io.Reader并从另一个io.Reader中读取数据的rot13Reader，通过应用rot13代换密码对数据流进行修改
type rot13Reader struct {
	r io.Reader
}

func rot13(p byte) byte { //字母转换
	switch {
	case p >= 'A' && p <= 'M':
		p += 13
	case p >= 'N' && p <= 'Z':
		p -= 13
	case p >= 'a' && p <= 'm':
		p += 13
	case p >= 'n' && p <= 'z':
		p -= 13
	}
	return p
}

func (rot13Reader rot13Reader) Read(b []byte) (int, error) {
	rot := make([]byte, 8)
	n, err := rot13Reader.r.Read(rot)
	for index, value := range rot[:n] {
		b[index] = rot13(value)
	}
	return n, err
}
