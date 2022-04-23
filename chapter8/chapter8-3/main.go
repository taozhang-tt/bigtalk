package main

// 工厂方法

import "fmt"

type IFactory interface {
	CreateOperation() Operation
}

type AddFactory struct{}

func (af *AddFactory) CreateOperation() Operation { return new(OperationAdd) }

type SubFactory struct{}

func (sf *SubFactory) CreateOperation() Operation { return new(OperationSub) }

type MulFactory struct{}

func (mf *MulFactory) CreateOperation() Operation { return new(OperationMul) }

type DivFactory struct{}

func (df *DivFactory) CreateOperation() Operation { return new(OperationDiv) }

type Operation interface {
	GetResult(num1, num2 float64) float64
}

// 加法类
type OperationAdd struct{}

func (o *OperationAdd) GetResult(num1, num2 float64) float64 {
	return num1 + num2
}

// 减法类
type OperationSub struct{}

func (o *OperationSub) GetResult(num1, num2 float64) float64 {
	return num1 - num2
}

// 乘法类
type OperationMul struct{}

func (o *OperationMul) GetResult(num1, num2 float64) float64 {
	return num1 * num2
}

// 除法类
type OperationDiv struct{}

func (o *OperationDiv) GetResult(num1, num2 float64) float64 {
	return num1 / num2
}

// 客户端
func main() {
	fmt.Println("请输入数字A：")
	var num1 float64
	fmt.Scanln(&num1)

	fmt.Println("请输入数字B：")
	var num2 float64
	fmt.Scanln(&num2)

	factory := new(AddFactory)
	operator := factory.CreateOperation()
	fmt.Printf("结果是：%v\n", operator.GetResult(num1, num2))
}
