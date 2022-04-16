package main

// 把加减乘除等运算分离，修改其中一个不影响另外的几个，增加运算算法也不影响其他代码

func main() {

}

type Operation struct {
	Num1 float64
	Num2 float64
}

func (o *Operation) GetResult() float64 {
	return 0
}

// 加法类
type OperationAdd struct {
	Operation
}

func (o *OperationAdd) GetResult() float64 {
	return o.Num1 + o.Num2
}

// 减法类
type OperationSub struct {
	Operation
}

func (o *OperationSub) GetResult() float64 {
	return o.Num1 - o.Num2
}

// 乘法类
type OperationMul struct {
	Operation
}

func (o *OperationMul) GetResult() float64 {
	return o.Num1 * o.Num2
}

// 除法类
type OperationDiv struct {
	Operation
}

func (o *OperationDiv) GetResult() float64 {
	return o.Num1 / o.Num2
}
