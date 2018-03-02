// https://tour.go-zh.org/welcome/1
// https://tour.go-zh.org/basics/1

package main

// import "fmt"

// “分组”形式的导入语句
import (
	"fmt"
	"math"
	"math/rand"
	"math/cmplx"
)



var c, python, java bool // var 语句可以出现在包或函数级别
var i, j int = 1, 2 //变量声明可以包含初始值，每个变量对应一个
// 如果初始化值已存在，则可以省略类型；变量会从初始值中获得类型

// 函数外的每个语句都必须以关键字开始（ var 、 func 等等）， 因此 := 结构不能在函数外使用
var mm = 1.2

var ( // 同导入语句一样，变量声明也可以“分组”成一个语法块。
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const Pi = 3.14 // 常量不能用 := 语法声明

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100 //2进制10……0
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99 //2进制首位向右移动99位
)



func add(x int, y int) int {
	return x + y
}

func mul(x, y int) int {
	return x * y
}

func swap(x, y string) (string, string) {
	return y, x // 函数可以返回任意数量的返回值
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // 没有参数的 return 语句返回已命名的返回值：x、y
	// Go 的返回值可被命名，它们会被视作定义在函数顶部的变量
	// 返回值的名称应当具有一定的意义，它可以作为文档使用
	// 直接返回语句应当仅用在下面这样的短函数中。在长的函数中它们会影响代码的可读性
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}



func main() {
	fmt.Println("Hello, world!")

	fmt.Println("My favorite number is", rand.Intn(10))
	rand.Seed(5) // 要得到不同的数字，需为生成器提供不同的种子数，参见 rand.Seed
	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Printf("Now you have %g problems.", math.Sqrt(7))

	fmt.Println("")
	fmt.Println(math.Pi) // not math.pi

	fmt.Println(add(42, 13))
	fmt.Println(mul(42, 13))

	a, b := swap("hello", "world!")
	fmt.Println(a, b)

	fmt.Println(split(17))

	var h int
	fmt.Println(h, c, python, java) // 未赋值时的初始值
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

	// 在函数中，简洁赋值语句 := 可在类型明确的地方代替 var 声明
	k := 3
	c2, python2, java2 := false, true, "yes! PPG！！"
	fmt.Println(i, j, k, c2, python2, java2)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var ii int // 没有明确初始值的变量声明会被赋予它们的零值
	var ff float64
	var bb bool
	var ss string
	fmt.Printf("%v %v %v %q\n", ii, ff, bb, ss)

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f) // Go 在不同类型的项之间赋值时需要显式转换
	fmt.Println(x, y, z)

	v := 42 // 修改这里！
	fmt.Printf("v is of type %T\n", v)
	v2 := 42.1234 // 修改这里！
	fmt.Printf("v2 is of type %T\n", v2)
	v3 := 0.867 + 0.5i // complex128
	fmt.Printf("v3 is of type %T\n", v3)

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	// fmt.Println(needInt(Big)) 错误：overflows int
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}





