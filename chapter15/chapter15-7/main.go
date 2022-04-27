package main

import "fmt"

// 用简单工厂来改进抽象工厂模式

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

type DataAccess struct {
	db string
}

func (da *DataAccess) CreateUser() IUser {
	var user IUser
	switch da.db {
	case "Sqlserver":
		user = new(SqlserverUser)
	case "Access":
		user = new(AccessUser)
	}
	return user
}

func (da *DataAccess) CreateDepartment() IDepartment {
	var department IDepartment
	switch da.db {
	case "Sqlserver":
		department = new(SqlserverDepartment)
	case "Access":
		department = new(AccessDepartment)
	}
	return department
}

// 客户端
func main() {
	var user User = User{}
	var department Department = Department{}

	da := &DataAccess{"Sqlserver"}

	iu := da.CreateUser()
	iu.Insert(user)
	iu.Get(1)

	id := da.CreateDepartment()
	id.Insert(department)
	id.Get(1)
}
