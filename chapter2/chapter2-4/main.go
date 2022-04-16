package main

// 抽象策略
type Stragety interface {
	Algorithm()
}

// 具体策略A类
type ConcreteStragetyA struct{}

func (csa *ConcreteStragetyA) Algorithm() {
	print("算法A")
}

// 具体策略B类
type ConcreteStragetyB struct{}

func (csb *ConcreteStragetyB) Algorithm() {
	print("算法B")
}

// 上下文
type Context struct {
	stragety Stragety
}

// 上下文初始化时传入具体的策略对象
func NewContext(stragety Stragety) *Context {
	return &Context{stragety}
}

// 上下文接口
func (c *Context) ContextInterface() {
	c.stragety.Algorithm()
}

// 客户端代码
func main() {
	var context *Context

	context = NewContext(new(ConcreteStragetyA))
	context.ContextInterface()

	context = NewContext(new(ConcreteStragetyB))
	context.ContextInterface()
}
