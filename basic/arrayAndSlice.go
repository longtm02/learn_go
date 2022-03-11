package basic

import "fmt"

type Student struct {
	id   int
	name string
}

func ArrayAndSlice() {
	array := []int{1, 12, 1}
	inArray(array)

	array1 := []int{0, 12, 1}
	inArray(array1)

	array2 := []int{0, 12, 0}
	inArray(array2)
}

func inArray(array []int) {
	slice1 := make([]int, len(array))
	copy(slice1, array)
	for index1, _ := range array {
		for index2, _ := range array {
			if index1 != index2 {
				slice1[index1] = slice1[index1] * slice1[index2]
			}
		}
	}

	fmt.Println(array, slice1)
}
