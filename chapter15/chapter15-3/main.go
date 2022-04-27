package main

import "fmt"

// 用了工厂方法模式的数据访问程序

// 用户类
type User struct {
	id   int
	name string
}

type IUser interface {
	Insert(user User)
	Get(id int)
}

type SqlserverUser struct{}

func (su *SqlserverUser) Insert(user User) {
	fmt.Println("在SQL Server中给User表增加一条记录")
}

func (su *SqlserverUser) Get(id int) {
	fmt.Println("在SQL Server中根据ID得到User表的一条记录")
}

type AccessUser struct{}

func (su *AccessUser) Insert(user User) {
	fmt.Println("在Access中给User表增加一条记录")
}

func (su *AccessUser) Get(id int) {
	fmt.Println("在Access中根据ID得到User表的一条记录")
}

type IFactory interface {
	CreateUser() IUser
}

type SqlserverFactory struct{}

func (sf *SqlserverFactory) CreateUser() IUser {
	return new(SqlserverUser)
}

type AccessFactory struct{}

func (sf *AccessFactory) CreateUser() IUser {
	return new(AccessUser)
}

// 客户端
func main() {
	var user User = User{}
	var factory IFactory = new(SqlserverFactory)
	var iu IUser = factory.CreateUser()

	iu.Insert(user)
	iu.Get(1)
}
