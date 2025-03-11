package main

import "fmt"

func main() {

	//Array
	//testArray()

	//slice
	//testSlice()
	testSliceFunc()

	//map
	//testMap()

	//struct
	//testStruct()

	//point
	//testPoint()

}

func testPoint() {
	var a int = 2
	var p *int = &a
	fmt.Println(&a)
	fmt.Println(p)
	fmt.Println(&p)
	fmt.Println(*p)
	*p = 222
	fmt.Println(a)

	//arr
	var arr1 [3]int = [3]int{1, 2, 3}
	var pArr1 *[3]int = &arr1
	fmt.Println(pArr1)
	(*pArr1)[0] = 100
	fmt.Println(arr1)

	//[]point
	p1 := new(int)
	p2 := new(int)
	p3 := new(int)
	fmt.Println(p1, *p2, p3)
	*p1 = 1
	*p2 = 2
	*p3 = 3
	var numsPoint [3]*int = [3]*int{p1, p2, p3}
	fmt.Println(numsPoint)
	for _, value := range numsPoint {
		fmt.Println(*value)
	}

	//struct

	//map
	var map1 map[int]string = map[int]string{1: "a", 2: "b", 3: "c"}
	map1[1] = "4"
	fmt.Println(map1)

}

type Student struct {
	name string
	age  int
}

func testStruct() {
	//init
	var student Student
	student.name = "jack"
	student.age = 23
	fmt.Println(student)
	var student2 Student = Student{name: "calvin", age: 23}
	fmt.Println(student2)

	//slice crud
	var students []Student = []Student{{name: "jack", age: 23}, {name: "calvin", age: 23}}
	fmt.Println(students)
	students = append(students, student2)
	fmt.Println(students)
	students[2] = Student{}
	fmt.Println(students)
	students[1].age = 26
	fmt.Println(students)

	testChangeSliceStructValue(students)
	fmt.Println(students)

	//map crud
	var studentMap map[int]Student = make(map[int]Student)
	studentMap[1] = Student{name: "jack", age: 23}
	studentMap[2] = Student{name: "calvin", age: 23}
	fmt.Println(studentMap)
	delete(studentMap, 1)
	fmt.Println(studentMap)
	//point
	//studentMap[2].age = 22
	fmt.Println(studentMap)
}

func testMap() {
	//init
	var m1 = map[string]int{"java": 12, "python": 34, "Go": 23}
	m2 := map[string]int{"java": 12, "python": 34, "Go": 23}
	m3 := make(map[string]int)
	fmt.Println(m1, m2, m3)

	//set and get
	m3["java"] = 12
	m3["python"] = 34
	v, exist := m1["java"]
	fmt.Println(v, exist)

	//for range
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	//delete
	delete(m1, "java")
	fmt.Println(m1)

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

func testSliceFunc() {
	//var slice []int = make([]int, 5, 10)
	//fmt.Println(slice)
	//fmt.Println("slice' len", len(slice))
	//fmt.Println("slice' cap", cap(slice))
	//

	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	//var arr = []int{1, 2, 3, 4, 5}

	//[low:high:max]  len = high-low  cap = max - low
	slice := arr[1:3]
	fmt.Println(slice)
	fmt.Println("len:", len(slice))
	fmt.Println("cap:", cap(slice))

	//test append
	//sliceAppend := append(slice, 1, 2, 3, 4, 5, 5, 6, 7)
	sliceAppend := append(slice, 1)
	fmt.Println(slice)
	fmt.Println(sliceAppend)
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

// !!!struct has nothing with pointer
func testChangeSliceStructValue(slice []Student) {
	//won't change
	//for i := 0; i < len(slice); i++ {
	//	slice[i].name = "testChangeSliceValue"
	//}
	//will change
	for _, value := range slice {
		value.name = "testChangeSliceValue"
	}
}
