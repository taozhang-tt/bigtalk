package main

import "fmt"

// 实现一个计算器控制台程序，要求输入两个数和运算符号，得到结果。
func main() {
	var num1 int
	fmt.Print("请输入第1个数字：")
	fmt.Scanln(&num1)
	var operator string
	fmt.Print("请输入操作符号：")
	fmt.Scanln(&operator)
	var num2 int
	fmt.Print("请输入第2个数字：")
	fmt.Scanln(&num2)
	var ret int
	switch operator {
	case "+":
		ret = num1 + num2
	case "-":
		ret = num1 - num2
	case "*":
		ret = num1 * num2
	case "/":
		ret = num1 / num2
	}
	fmt.Printf("结果是：%v\n", ret)
}
