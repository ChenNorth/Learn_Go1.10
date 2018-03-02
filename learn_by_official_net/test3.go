// https://tour.go-zh.org/moretypes/1
// 14到23节未学习

package main

import (
	"fmt"
	"math"
)



type Vertex struct {
	// 一个结构体（struct）就是一个字段的集合
	X int
	Y int
}



var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)



func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}



func main() {
	// 指针：类型 *T 是指向 T 类型值的指针。其零值为 nil；
	// & 操作符会生成一个指向其操作数的指针
	// * 操作符表示指针指向的底层值，也就是通常所说的“间接引用”或“重定向”
	i, j := 42, 2701
	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i
	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	// 结构体字段可以通过结构体指针来访问。
	p2 := &v
	p2.X = 1e9
	fmt.Println(v)

	// 结构体文法通过直接列出字段的值来新分配一个结构体
	// 使用 Name: 语法可以仅列出部分字段，字段名的顺序无关
	// 特殊的前缀 & 返回一个指向结构体的指针
	fmt.Println(v1, p, v2, v3)

	var a [2]string // 类型 [n]T 表示拥有 n 个 T 类型的值的数组
	// 数组的长度是其类型的一部分，因此数组不能改变大小
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// 每个数组的大小都是固定的。 而切片则为数组元素提供动态大小的、灵活的视角
	// 切片通过两个下标来界定，即一个上界和一个下界，二者以冒号分隔
	// 它会选择一个半开区间，包括第一个元素，但排除最后一个元素
	var s []int = primes[1:4]
	fmt.Println(s)

	// 切片并不存储任何数据， 它只是描述了底层数组中的一段
	// 更改切片的元素会修改其底层数组中对应的元素
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a2 := names[0:2]
	b2 := names[1:3]
	fmt.Println(a2, b2)

	b2[0] = "XXX"
	fmt.Println(a2, b2)
	fmt.Println(names)

	// 切片文法类似于没有长度的数组文法
	// 会创建一个和上面相同的数组，然后构建一个引用了它的切片
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)
	s2 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s2)

	// 对于数组var a [10]int来说，以下切片是等价的：
	// a[0:10]、a[:10]、a[0:]、a[:]
	s3 := []int{2, 3, 5, 7, 11, 13}
	s3 = s3[1:4]
	fmt.Println(s3)
	s3 = s3[:2]
	fmt.Println(s3)
	s3 = s3[1:]
	fmt.Println(s3)

	// 切片拥有 长度 和 容量，长度就是它所包含的元素个数
	// 切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数
	// 切片 s 的长度和容量可通过表达式 len(s) 和 cap(s) 来获取
	s4 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s4)
	// Slice the slice to give it zero length.
	s4 = s4[:0]
	printSlice(s4)
	// Extend its length.
	s4 = s4[:4]
	printSlice(s4)
	// Drop its first two values.
	s4 = s4[2:]
	printSlice(s4)

	// 切片的零值是 nil，nil 切片的长度和容量为 0 且没有底层数组
	var s5 []int
	fmt.Println(len(s5), cap(s5), s5)
	if s5 == nil {
		fmt.Println("nil!")
	}

	// 切面可以用内建函数 make 来创建，这也是创建动态数组的方式
	// make 函数会分配一个元素为零值的数组并返回一个引用了它的切片：
	// a := make([]int, 5)  // len(a)=5
	// 要指定它的容量，需向 make 传入第三个参数：
	// b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	// b = b[:cap(b)] // len(b)=5, cap(b)=5
	// b = b[1:]      // len(b)=4, cap(b)=4
	a3 := make([]int, 5)
	printSlice2("a3", a3)
	b3 := make([]int, 0, 5)
	printSlice2("b3", b3)
	c3 := b3[:2]
	printSlice2("c3", c3)
	d3 := c3[2:5]
	printSlice2("d3", d3)

	// 14未学习
	// 23未学习

	// 函数也是值。它们可以像其它值一样传递。
	// 函数值可以用作函数的参数或返回值
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	// Go 函数可以是一个闭包。
	// 闭包是一个函数值，它引用了其函数体之外的变量。
	// 该函数可以访问并赋予其引用的变量的值，
	// 换句话说，该函数被“绑定”在了这些变量上
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}