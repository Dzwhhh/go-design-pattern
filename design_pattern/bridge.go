package design_pattern

import "fmt"

type printer interface {
	SetBrand(Brand)
	GetBrand() string
	PrintFile()
}

type LaserPrinter struct {
	brand Brand
}

func (l *LaserPrinter) SetBrand(brand Brand) {
	l.brand = brand
}
func (l *LaserPrinter) GetBrand() string {
	return l.brand.getName()
}

func (l *LaserPrinter) PrintFile() {
	fmt.Printf("print by %s \n", l.GetBrand())
}

type InkPrinter struct {
	brand Brand
}

func (i *InkPrinter) SetBrand(brand Brand) {
	i.brand = brand
}
func (i *InkPrinter) GetBrand() string {
	return i.brand.getName()
}

func (i *InkPrinter) PrintFile() {
	fmt.Printf("print by %s \n", i.GetBrand())
}

type Brand interface {
	getName() string
}

type Epson struct {
	name string
}

func (e *Epson) getName() string {
	e.name = "epson"
	return e.name
}

type Hp struct {
	name string
}

func (h *Hp) getName() string {
	h.name = "hp"
	return h.name
}
