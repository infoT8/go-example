package main

import "fmt"

// 变量的初始化
// var 变量名 类型 = 表达式
func main() {
	// 指定类型
	var a int = 1
	fmt.Println(a)

	// 不指定类型
	var b = true
	fmt.Println(b)

	// 先声明变量 在赋值
	var c int
	c = 1
	fmt.Println(c)

}
