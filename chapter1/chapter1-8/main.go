package main

import "fmt"

// 业务封装，让业务逻辑和界面逻辑分开，降低耦合度。

// 客户端代码
func main() {
	fmt.Println("请输入数字A：")
	var num1 float64
	fmt.Scanln(&num1)

	fmt.Println("请输入运算符号：")
	var operator string
	fmt.Scanln(&operator)

	fmt.Println("请输入数字B：")
	var num2 float64
	fmt.Scanln(&num2)

	operation := new(Operation)
	fmt.Printf("结果是：%v\n", operation.GetResult(num1, num2, operator))
}

// Operation 计算类
type Operation struct{}

func (o *Operation) GetResult(num1, num2 float64, operator string) float64 {
	var ret float64
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
	return ret
}
