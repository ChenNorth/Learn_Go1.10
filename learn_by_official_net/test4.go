// https://tour.go-zh.org/methods/1

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Go 没有类。不过你可以为结构体类型定义方法。
// 方法就是一类带特殊的 接收者 参数的函数
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale1(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
// 带指针参数的函数必须接受一个指针

func (v Vertex) Scale2(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 把 Scale 方法重写为函数
func Scale3(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
// 而以指针为接收者的方法被调用时，接收者既能为值又能为指针

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64
// 可以为非结构体类型声明方法
func (f MyFloat) Abs2() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
// 你只能为在同一包内定义的类型的接收者声明方法，
// 而不能为其它包内定义的类型（包括 int 之类的内建类型）的接收者声明方法
// 接收者的类型定义和方法声明必须在同一包内；不能为内建类型声明方法



func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs()) // 方法只是个带接收者参数的函数
	// 或者
	fmt.Println(Abs(v)) // 正常的函数

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs2()) // 一个带 Abs 方法的数值类型 MyFloat

	// 可以为指针接收者声明方法，
	// 这意味着对于某类型 T ，接收者的类型可以用 *T 的文法
	v1 := Vertex{3, 4}
	v1.Scale1(10)
	fmt.Println(v1.Abs())
	// 移除v.Scale1函数声明中的*后的效果
	v2 := Vertex{3, 4}
	v2.Scale2(10)
	fmt.Println(v2.Abs())

	v3 := Vertex{3, 4}
	Scale3(&v3, 10)
	fmt.Println(Abs(v3))

	v4 := Vertex{3, 4}
	v4.Scale1(2)
	Scale3(&v4, 10)
	p := &Vertex{4, 3}
	p.Scale1(3)
	Scale3(p, 8)
	fmt.Println(v4, p)

	// 7未学习
	// 9未学习

	// 后接test5.go
}