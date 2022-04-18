package main

// 客户端代码
func main() {
	context := &Context{new(ConcreteStateA)}

	context.Request()
	context.Request()
	context.Request()
	context.Request()
}

type Context struct {
	State State
}

// 对请求做处理，并设置下一状态
func (c *Context) Request() {
	c.State.Handle(c)
}

// State，定义一个接口以封装与Context的一个特定状态相关的行为
type State interface {
	Handle(context *Context)
}

// ConcreteState类，具体状态，
// 每一个类实现一个与Context的一个状态相关的行为

type ConcreteStateA struct{}

// 设置 ContreteStateA 的下一状态是 ConcreteStateB
func (csa *ConcreteStateA) Handle(context *Context) {
	context.State = new(ConcreteStateB)
}

type ConcreteStateB struct{}

// 设置 ContreteStateB 的下一状态是 ConcreteStateA
func (csb *ConcreteStateB) Handle(context *Context) {
	context.State = new(ConcreteStateA)
}
