package main

import "fmt"

func main() {

	//testPolymorphism
	testPolymorphism()

	//testAny
	testAny()

}

func testPolymorphism() {
	var keyBoard KeyBoard
	var mouse Mouse
	DeviceConnect(&keyBoard)
	DeviceConnect(&mouse)
}

func testAny() {
	// any is alias of interface{}
	var i interface{}
	i = []any{1, "1", 4, true}
	fmt.Println(i)
	value, ok := i.([]any)
	fmt.Println(value, ok)
	value1, ok1 := value[1].(string)
	fmt.Println(value1, ok1)
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
