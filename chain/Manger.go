package main

// 主管  经理  总经理  董事长
// 50 100   150   200
type Manger interface {
	HaveRight(money int) bool                     //判断权限
	HandleFeeRequest(name string, money int) bool //向后传递
}
