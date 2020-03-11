package method

import (
	"fmt"
	"image"
	"image/color"
)

func main() {
	// 使用图像
	imageOne()
}

// 图像
// image包定义了Image接口
// type Image interface {
//     ColorModel() color.Model
//     Bounds() Rectangle
//     At(x, y int) color.Color
// }
// 注意：Bounds方法的返回值Rectangle实际上是一个image.Rectangle，它在image包中声明
// color.Model和color.Color类型也是接口，但是通常因为直接使用预定义实现的image.RGBA和image.RGBAModel而被忽视
// 这些接口和类型也由image/color定义
func imageOne() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

// 练习：图像
// 编写图像生成器
// 定义自己的image类型，实现必要的方法并调用pic.ShowImage
// Bounds应当返回一个image.Rectangle，例如image.Rect(0, 0, w, h)
// ColorModel应当返回color.Model
// At应当返回一个颜色，上一个图片生成器的值v对应于此的color.RGBA(v, v, 255, 255)
type Image struct {
}

func (img Image) ColorModel() color.Model {
	return color.RGBA64Model
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), uint8(255), uint8(255)}
}
