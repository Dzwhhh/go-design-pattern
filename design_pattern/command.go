package design_pattern

import "fmt"

type Command interface {
	Execute()
}

type Button struct {
	Command
}

func (b *Button) Press() {
	b.Command.Execute()
}

type Device interface {
	on()
	off()
}

type TV struct{}

func (t *TV) on() {
	fmt.Println("TV is on")
}

func (t *TV) off() {
	fmt.Println("TV is off")
}

type Lights struct{}

func (t *Lights) on() {
	fmt.Println("Lights on")
}

func (t *Lights) off() {
	fmt.Println("Lights off")
}

type OnCommand struct {
	Device Device
}

func (on *OnCommand) Execute() {
	on.Device.on()
}

type OffCommand struct {
	Device Device
}

func (off *OffCommand) Execute() {
	off.Device.off()
}
