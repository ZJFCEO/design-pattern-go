package main

func main() {
	ctx := NewDayContext()
	todayAndNext := func() {
		ctx.Today()
		ctx.Next()
	}
	for i := 0; i < 98; i++ {
		todayAndNext()
	}

}
