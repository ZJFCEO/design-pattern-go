package main

func main() {
	fac := MySQLFactory{}
	fac.CreateOrderMainDAO().SaveOrderMain()
}
