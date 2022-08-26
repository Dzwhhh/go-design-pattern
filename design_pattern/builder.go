package design_pattern

import "fmt"

type absBuilder interface {
	setStep1()
	setStep2()
	getResult()
}

type BuilderA struct {
	head string
	body string
}

func (a *BuilderA) setStep1() {
	a.head = "aaa-h"
	fmt.Println("A: build step 1")
}
func (a *BuilderA) setStep2() {
	a.body = "aaa-b"
	fmt.Println("A: build step 2")
}

func (a *BuilderA) getResult() {
	fmt.Printf("head:%s, body:%s\n", a.head, a.body)
}

type BuilderB struct {
	leftHand  string
	rightHand string
}

func (b *BuilderB) setStep1() {
	b.leftHand = "left"
	fmt.Println("B: build step 1")
}

func (b *BuilderB) setStep2() {
	b.rightHand = "right"
	fmt.Println("B: build step 2")
}

func (b *BuilderB) getResult() {
	fmt.Printf("left:%s, right:%s\n", b.leftHand, b.rightHand)
}

type Director struct {
	Builder absBuilder
}

func (d *Director) NewBuilder(builder absBuilder) {
	d.Builder = builder
}

func (d *Director) Build() {
	d.Builder.setStep1()
	d.Builder.setStep2()
}

func (d *Director) Show() {
	d.Builder.getResult()
}
