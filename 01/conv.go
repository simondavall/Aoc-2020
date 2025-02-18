package main

import "strconv"

func toInt32Array(strArray []string) ([]int32, error) {
	var intArray []int32
	for _, str := range strArray {
		n, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		intArray = append(intArray, int32(n))
	}
	return intArray, nil
}
