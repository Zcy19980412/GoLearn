package main

import "fmt"

func main() {

	nums := []int{1, 3, 4, 65, 2, 1, 4, 6, 78, 3}
	//bubbleSort(nums)
	selectionSort(nums)
	fmt.Println(nums)
}

func bubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func selectionSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}
}
