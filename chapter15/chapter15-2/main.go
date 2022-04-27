package main

import "fmt"

// 最基本的数据访问程序

// 用户类
type User struct {
	id   int
	name string
}

// SqlserverUser类——用于操作User表
type SqlserverUser struct{}

func (su *SqlserverUser) Insert(user User) {
	fmt.Println("在SQL Server中给User表增加一条记录")
}

func (su *SqlserverUser) Get(id int) {
	fmt.Println("在SQL Server中根据ID得到User表的一条记录")
}

// 客户端
func main() {
	var user User = User{}
	var su *SqlserverUser = new(SqlserverUser) // 与 SQL Server 耦合

	su.Insert(user)
	su.Get(1)
}
