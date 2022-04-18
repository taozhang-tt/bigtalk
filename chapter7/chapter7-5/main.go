package main

import "fmt"

// Subject类，定义了RealSubject和Proxy的共用接口，这样就在任何使用RealSubject的地方都可以使用Proxy
type Subject interface {
	Request()
}

// RealSubject类，定义Proxy所代表的真实实体
type RealSubject struct{}

func (rs *RealSubject) Request() {
	fmt.Println("真实的请求")
}

// Proxy类，保存一个引用使得代理可以访问实体，并提供一个与Subject的接口相同的接口，这样代理就可以用来替代实体
type Proxy struct {
	realsubject *RealSubject
}

func (p *Proxy) Request() {
	if p.realsubject == nil {
		p.realsubject = new(RealSubject)
	}
	p.realsubject.Request()
}

func main() {
	proxy := new(Proxy)
	proxy.Request()
}
