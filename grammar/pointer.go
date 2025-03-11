package main

import "fmt"

func main() {
	//testBasic()

	testHeap()
}

func testBasic() {
	var a int = 1
	var p *int = &a
	fmt.Println(p)
	fmt.Println(&p)

	testPassBasic(p)
}

func testPassBasic(a *int) {
	fmt.Println(a)
	fmt.Println(&a)
}

func testHeap() {
	//new : 在heap上申请一段内存空间
	var p = new(int)
	fmt.Println(*p)
	fmt.Println(p)
	fmt.Println(&p)

}
