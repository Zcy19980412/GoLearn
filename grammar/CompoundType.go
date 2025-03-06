package main

import "fmt"

func main() {

	//Array
	testArray()

}

func testArray() {
	//define and init
	var numbers1 []int
	fmt.Println(numbers1)

	var numbers2 [3]int
	fmt.Println(numbers2)

	var numbers3 [3]int = [3]int{1, 2, 3}
	fmt.Println(numbers3)

	var numbers4 [3]int = [3]int{1, 2}
	fmt.Println(numbers4)

	numbers5 := [3]int{1, 2, 3}
	fmt.Println(numbers5)

	//foreach
	numbers := numbers5
	for index, value := range numbers {
		fmt.Println(index, value)
	}

	//testArrayFunction
	testArrayPassValue(numbers)
	fmt.Println(numbers)

	//two-dimensional array
	var numbers6 [2][3]int = [2][3]int{1: {3, 4}}
	fmt.Println(numbers6)

}

func testArrayPassValue(arr [3]int) {
	arr[0] = 200
}
