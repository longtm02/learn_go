package basic

import (
	"fmt"
	"strconv"
	"strings"
)

// bài tập : aabbcc => a2b2c2

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func TestCreatePackage2(params string) {
	array := strings.Split(params, "")

	result := ""

	for i := 0; i < len(params); i++ {
		number := 0
		for j := 0; j < len(array); j++ {
			if array[i] == array[j] && i > j {
				break
			}
			if array[i] == array[j] && i <= j {
				number++
			}

		}

		if number != 0 {
			result = result + array[i] + strconv.Itoa(number)
		}

	}

	fmt.Println("result", params, "--------->>>>", result, RemoveIndex(array, 5))
}
