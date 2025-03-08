package main

import "fmt"

func main() {

	//testPolymorphism
	testPolymorphism()

}

func testPolymorphism() {
	var keyBoard KeyBoard
	var mouse Mouse
	DeviceConnect(&keyBoard)
	DeviceConnect(&mouse)
}

// ------------------- device demo -------------------
type Device interface {
	Read()
	Write()
}

func DeviceConnect(d Device) {
	d.Read()
	d.Write()
}

type Mouse struct {
}

type KeyBoard struct {
}

func (k *KeyBoard) Read() {
	fmt.Println("keyBoard read")
}

func (k *KeyBoard) Write() {
	fmt.Println("keyBoard write")
}

func (m *Mouse) Read() {
	fmt.Println("mouse read")
}

func (m *Mouse) Write() {
	fmt.Println("mouse write")
}
