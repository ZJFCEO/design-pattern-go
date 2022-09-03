package main

// 每个星期行为
type Week interface {
	Today()
	Next(*DayContext)
}

// 星期数据结构
type DayContext struct {
	today Week
}

func NewDayContext() *DayContext {
	return &DayContext{&Sunday{}}
}
func (dc *DayContext) Today() {
	dc.today.Today() //今天
}
func (dc *DayContext) Next() {
	dc.today.Next(dc) //m明天
}
