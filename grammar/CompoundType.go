package main

import "fmt"

func main() {

	//Array
	//testArray()

	//slice
	testSlice()

}

func testSlice() {
	//init
	var slice1 []int
	var slice2 []int = []int{1, 2, 3}
	var slice3 []int = make([]int, 3)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(len(slice3), cap(slice3))

	//update
	slice1 = append(slice1, slice2...)
	slice2 = append(slice2, slice3...)
	slice3 = append(slice3, slice2...)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(len(slice3), cap(slice3))

	//truncation
	s1 := []int{1, 2, 3}
	// [)
	s2 := s1[0:1]
	s2 = s1[:]
	s2 = s1[0:]
	s2 = s1[:2]
	fmt.Println(s2)
	//same point!!!
	s2[0] = 100
	fmt.Println(s1)

	//copy
	copy1 := []int{1, 2, 3}
	var copy2 []int = make([]int, 1, 3)
	copy(copy2, copy1)
	fmt.Println(copy1, copy2)

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
