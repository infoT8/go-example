package main

import "fmt"

// 批量定义变量
var (
	e int
	f string
	j bool
	h float64
)
// 变量的声明
func main() {
	// 声明变量用var关键字
	// 格式  var  变量名  变量类型
	var a int
	var b string
	var c float32
	var d bool
	// go 定义变量 默认都有初始值 int(0) string("") float(0.0) bool(false)
	fmt.Printf("a=%v b=%#v c=%v  d=%v\n", a,b,c,d) 
	fmt.Printf("e=%v f=%#v j=%v  h=%v\n", e,f,j,h) 

	// 第二种方式定义变量 不用var关键字  go语言会默认推断类型
	i := 1
	fmt.Printf("i=%v type=", i)
	fmt.Printf("%T\n", i)

	// 多变量声明
	a1, a2 := 1 ,true
	fmt.Printf("a1=%v a2=%v\n", a1,a2)

	// 因为简洁和灵活的特点，简短变量声明被广泛用于大部分的局部变量的声明和初始化。
	// var 形式的声明语句往往是用于需要显式指定变量类型地方，
}