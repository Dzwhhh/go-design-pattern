package design_pattern

import "fmt"

type Goods interface {
	RegisterObserver(Observer)
	DeRegisterObserver(Observer)
	NotifyAll()
}

type Observer interface {
	Update(*Item)
	GetId() string
}

type Item struct {
	Price     int
	observers []Observer
}

func (i *Item) RegisterObserver(observer Observer) {
	i.observers = append(i.observers, observer)
}

func (i *Item) DeRegisterObserver(observer Observer) {
	for idx, o := range i.observers {
		if o.GetId() == observer.GetId() {
			i.observers[idx], i.observers[len(i.observers)-1] = i.observers[len(i.observers)-1], i.observers[idx]
			i.observers = i.observers[:len(i.observers)-1]
		}
	}
}

func (i *Item) UpdatePrice(newPrice int) {
	i.Price = newPrice
	i.NotifyAll()
}

func (i *Item) NotifyAll() {
	for _, o := range i.observers {
		o.Update(i)
	}
}

type Buyer struct {
	Id          string
	ExpectPrice int
	canBuy      bool
}

func (b *Buyer) Update(item *Item) {
	if item.Price <= b.ExpectPrice {
		b.canBuy = true
		fmt.Println(b.Id + ": OK, I can buy it")
	} else {
		b.canBuy = false
		fmt.Println(b.Id + ": The price is still too high")
	}
}

func (b *Buyer) GetId() string {
	return b.Id
}
