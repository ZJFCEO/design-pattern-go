package main

type OrderMainDAO interface { //订单记录
	SaveOrderMain()
}

type OrderDetailDAO interface { //订单详情
	SaveOrderDetail()
}

type DAOFactory interface { //抽象工厂接口
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}
