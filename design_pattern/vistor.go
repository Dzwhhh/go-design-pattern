package design_pattern

import "fmt"

type visitor interface {
	Visit(*Cube)
}

type Cube struct {
	length int
	width  int
	height int
}

func NewCube(l, w, h int) *Cube {
	return &Cube{
		length: l,
		width:  w,
		height: h,
	}
}

func (c *Cube) Accept(visitor visitor) {
	visitor.Visit(c)
}

type AreaVisitor struct{}

func (a *AreaVisitor) Visit(cube *Cube) {
	totalArea := 2 * (cube.length*cube.width + cube.length*cube.height + cube.width*cube.height)
	fmt.Println("cube surface area:", totalArea)
}

type VolumeVisitor struct{}

func (v *VolumeVisitor) Visit(cube *Cube) {
	volume := cube.height * cube.length * cube.width
	fmt.Println("cube volume:", volume)
}

type ConnectClient struct{}

func (c *ConnectClient) InsertUsb(computer Computer) {
	computer.InsertSquareUsb()
}
