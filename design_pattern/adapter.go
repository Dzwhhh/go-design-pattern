package design_pattern

import "fmt"

type Computer interface {
	InsertSquareUsb()
}

type Mac struct{}

func (m *Mac) InsertSquareUsb() {
	fmt.Println("insert success with square usb")
}

type Windows struct{}

func (w *Windows) InsertCirclePort() {
	fmt.Println("insert success with circle port")
}

type WindowsAdaptor struct {
	Win *Windows
}

func (w *WindowsAdaptor) InsertSquareUsb() {
	fmt.Println("port not match, use adaptor")
	w.Win.InsertCirclePort()
}
