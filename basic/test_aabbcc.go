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
	number := 0
	currentItem := array[0]

	for i := 0; i < len(params); i++ {

		if array[i] == currentItem {
			number++
			continue
		}
		result = result + currentItem + strconv.Itoa(number)
		currentItem = array[i]
		number = 1
	}

	fmt.Println("result", params, "--------->>>>", result+currentItem+strconv.Itoa(number), params[0])
}
