package design_pattern

import "fmt"

type service interface {
	Execute()
}

type OrderService struct{}

func (o *OrderService) Execute() {
	fmt.Println("正在下订单，准备跳转到支付流程...")
}

type PayService struct{}

func (p *PayService) Execute() {
	fmt.Println("正在支付...支付成功！")
}

type LogisticService struct{}

func (l *LogisticService) Execute() {
	fmt.Println("进入物流系统，准备出库...出库成功！")
}

type BuyFacade struct{}

func (b *BuyFacade) BuySomething() {
	new(OrderService).Execute()
	new(PayService).Execute()
	new(LogisticService).Execute()
}
