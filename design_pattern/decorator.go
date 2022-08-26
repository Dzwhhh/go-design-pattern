package design_pattern

import "fmt"

type pizza interface {
	GetPrice() int
}

type PizzaA struct{}

func (p *PizzaA) GetPrice() int {
	return 20
}

type PizzaB struct{}

func (p *PizzaB) GetPrice() int {
	return 18
}

type ChessTopping struct {
	Pizza pizza
}

func (c *ChessTopping) GetPrice() int {
	fmt.Println("pizza with chess")
	pizzaPrice := c.Pizza.GetPrice()
	return pizzaPrice + 8
}

type TomatoTopping struct {
	Pizza pizza
}

func (c *TomatoTopping) GetPrice() int {
	fmt.Println("pizza with tomato")
	pizzaPrice := c.Pizza.GetPrice()
	return pizzaPrice + 5
}
