// https://tour.go-zh.org/methods/9

package main

import (
	"fmt"
	"math"
)

// 接口类型 是由一组方法签名定义的集合
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	// math.Sqrt2参考https://studygolang.com/articles/6402
	// Sqrt2   = 1.414213562373095048801
	v := Vertex{3, 4}
	a = f  // a MyFloat 实现了 Abser
	a = &v // a *Vertex 实现了 Abser
	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	// a = v
	fmt.Println(a.Abs())

	// 接口与隐式实现
	// 类型通过实现一个接口的所有方法来实现该接口。
	// 隐式接口从接口的实现中解耦了定义，这样接口的实现可以出现在任何包中，无需提前准备
	// 因此，也就无需在每一个实现上增加新的接口名称，这样同时也鼓励了明确的接口定义

	// 后接test5.go
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

