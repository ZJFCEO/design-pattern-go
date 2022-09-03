package main

import "fmt"

type MySQLMainDAO struct {
}

func (*MySQLMainDAO) SaveOrderMain() {
	fmt.Println("Mysql  main save ")
}

type MySQLDetailDAO struct {
}

func (*MySQLDetailDAO) SaveOrderDetail() {
	fmt.Println("mysql detail save ")
}
