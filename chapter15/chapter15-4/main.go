package main

import "fmt"

// 用了抽象工厂模式的数据访问程序

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

type Department struct {
	id   int
	name string
}

type IDepartment interface {
	Insert(department Department)
	Get(id int)
}

type SqlserverDepartment struct{}

func (sd *SqlserverDepartment) Insert(department Department) {
	fmt.Println("在SQL Server中给Department表增加一条记录")
}

func (sd *SqlserverDepartment) Get(id int) {
	fmt.Println("在SQL Server中根据ID得到Department表的一条记录")
}

type AccessDepartment struct{}

func (ad *AccessDepartment) Insert(department Department) {
	fmt.Println("在Access中给Department表增加一条记录")
}

func (ad *AccessDepartment) Get(id int) {
	fmt.Println("在Access中根据ID得到Department表的一条记录")
}

type IFactory interface {
	CreateUser() IUser
	CreateDepartment() IDepartment
}

type SqlserverFactory struct{}

func (sf *SqlserverFactory) CreateUser() IUser {
	return new(SqlserverUser)
}

func (sf *SqlserverFactory) CreateDepartment() IDepartment {
	return new(AccessDepartment)
}

type AccessFactory struct{}

func (af *AccessFactory) CreateUser() IUser {
	return new(AccessUser)
}

func (af *AccessFactory) CreateDepartment() IDepartment {
	return new(AccessDepartment)
}

// 客户端
func main() {
	var user User = User{}
	var department Department = Department{}

	var factory IFactory = new(SqlserverFactory)

	var iu IUser = factory.CreateUser()
	iu.Insert(user)
	iu.Get(1)

	var id IDepartment = factory.CreateDepartment()
	id.Insert(department)
	id.Get(1)
}
