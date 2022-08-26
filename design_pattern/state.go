package design_pattern

import (
	"fmt"
)

type State interface {
	SelectItem() error
	AddItem(int) error
	InsertMoney(int) error
	DispenseItem() error
}

type NoItemState struct {
	v *VendingMachine
}

func (n *NoItemState) SelectItem() error {
	return fmt.Errorf("item out of stock")
}

func (n *NoItemState) AddItem(count int) error {
	n.v.increaseItem(count)
	n.v.setState(n.v.hasItemState)
	return nil
}

func (n *NoItemState) InsertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}

func (n *NoItemState) DispenseItem() error {
	return fmt.Errorf("item out of stock")
}

type HasItemState struct {
	v *VendingMachine
}

func (h *HasItemState) SelectItem() error {
	if h.v.itemCount == 0 {
		h.v.setState(h.v.noItemState)
		return fmt.Errorf("no item present")
	}
	fmt.Println("item requested")
	h.v.setState(h.v.itemRequestedState)
	return nil
}

func (h *HasItemState) AddItem(count int) error {
	h.v.increaseItem(count)
	return nil
}

func (h *HasItemState) InsertMoney(money int) error {
	return fmt.Errorf("please select item first")
}

func (h *HasItemState) DispenseItem() error {
	return fmt.Errorf("please select item first")
}

type ItemRequestedState struct {
	v *VendingMachine
}

func (i *ItemRequestedState) SelectItem() error {
	return fmt.Errorf("item already requested")
}

func (i *ItemRequestedState) AddItem(count int) error {
	return fmt.Errorf("item dispense in progress")
}

func (i *ItemRequestedState) InsertMoney(money int) error {
	if money >= i.v.itemPrice {
		fmt.Println("money enter is ok")
		i.v.setState(i.v.hasMoneyState)
		return nil
	}
	return fmt.Errorf("insert money is not enough, please retry")
}

func (i *ItemRequestedState) DispenseItem() error {
	return fmt.Errorf("please insert money first")
}

type HasMoneyState struct {
	v *VendingMachine
}

func (h *HasMoneyState) SelectItem() error {
	return fmt.Errorf("item dispense in progress")
}

func (h *HasMoneyState) AddItem(count int) error {
	return fmt.Errorf("item dispense in progress")
}

func (h *HasMoneyState) InsertMoney(money int) error {
	return fmt.Errorf("money has already inserted")
}

func (h *HasMoneyState) DispenseItem() error {
	fmt.Println("dispensing item")
	h.v.itemCount -= 1
	if h.v.itemCount <= 0 {
		h.v.setState(h.v.noItemState)
	} else {
		h.v.setState(h.v.hasItemState)
	}
	return nil
}

type VendingMachine struct {
	currentState       State
	noItemState        State
	hasItemState       State
	itemRequestedState State
	hasMoneyState      State

	itemCount int
	itemPrice int
}

func NewVendingMachine(initCount int) *VendingMachine {
	v := &VendingMachine{
		itemCount: 0,
		itemPrice: 100,
	}
	v.noItemState = &NoItemState{v}
	v.hasItemState = &HasItemState{v}
	v.itemRequestedState = &ItemRequestedState{v}
	v.hasMoneyState = &HasMoneyState{v}

	v.itemCount = initCount
	if v.itemCount <= 0 {
		v.setState(v.noItemState)
	} else {
		v.setState(v.hasItemState)
	}
	return v
}

func (v *VendingMachine) AddItem(count int) {
	err := v.currentState.AddItem(count)
	if err != nil {
		panic(err)
	}
}

func (v *VendingMachine) SelectItem() {
	err := v.currentState.SelectItem()
	if err != nil {
		panic(err)
	}
}

func (v *VendingMachine) InsertMoney(money int) {
	err := v.currentState.InsertMoney(money)
	if err != nil {
		panic(err)
	}
}

func (v *VendingMachine) DispenseItem() {
	err := v.currentState.DispenseItem()
	if err != nil {
		panic(err)
	}
}

func (v *VendingMachine) setState(newState State) {
	v.currentState = newState
}

func (v *VendingMachine) increaseItem(count int) {
	fmt.Printf("add %d item success\n", count)
	v.itemCount += count
}
