// https://tour.go-zh.org/flowcontrol/1


package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)



func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// 同 for 一样， if 语句可以在条件表达式前执行一个简单的语句
	// 该语句声明的变量作用域仅在 if 之内
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// 这里开始就不能使用 v 了
	return lim
}



func main() {
	// defer 语句会将函数推迟到外层函数返回之后执行
	// 推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用
	// 即，本函数执行完标注“defer”的行后才执行
	defer fmt.Println("world")

	// Go 只有一种循环结构： for 循环
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum2 := 1
	for ; sum2 < 1000; { // 初始化语句和后置语句是可选的
		sum2 += sum2
		// fmt.Println(sum2)
	}
	fmt.Println(sum2)
	// 或者
	sum3 := 1
	for sum3 < 1000 {
		sum3 += sum3
	}
	fmt.Println(sum3)

	fmt.Println(sqrt(2), sqrt(-4)) // if

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)

	// switch 是编写一连串 if - else 语句的简便方法。
	// 它运行第一个值等于条件表达式的 case 语句。
	// 实际上，Go 自动提供了在这些语言中每个 case 后面所需的 break 语句。
	// 除非以 fallthrough 语句结束，否则分支会自动终止。
	// Go 的另一点重要的不同在于 switch 的 case 无需为常量，且取值不必为整数。
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}

	fmt.Println("")
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today - 1:
		fmt.Println("Yestaday.")
	case today + 2:
		fmt.Println("In two days.")
	case today - 2:
		fmt.Println("Two days ago.")
	default:
		fmt.Println("Too far away.")
	}

	// 没有条件的 switch
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	fmt.Println("hello") // defer

	// defer 栈：推迟的函数调用会被压入一个栈中。
	// 当外层函数返回时，被推迟的函数会按照后进先出的顺序调用
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")

	
}
