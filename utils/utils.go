package utils

import (
	"strconv"
	"strings"
)

func ToIntegerArray(source string) ([]int, error) {
	sourceArr := strings.Fields(source)

	arr := make([]int, len(sourceArr))
	for i, src := range sourceArr {
		num, err := strconv.Atoi(src)
		if err != nil {
			return nil, err
		}
		arr[i] = num
	}
	return arr, nil
}

func ToInteger(source string) (num int, err error) {
	num, err = strconv.Atoi(source)
	return
}

func ToStringArray(source ...int) (result []string) {
	result = make([]string, len(source))
	for i, num := range source {
		result[i] = strconv.Itoa(num)
	}
	return
}
