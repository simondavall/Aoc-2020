package aoc

import (
	"cmp"
	"slices"
)

func Filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func Transpose[T any](slice [][]T) [][]T {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]T, xl)
	for i := range result {
		result[i] = make([]T, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func Intersects[T cmp.Ordered](first []T, second []T) []T {
	slices.Sort(first)
	slices.Sort(second)
	commonValues := make([]T, 0)
	for i, j := 0, 0; i < len(first) && j < len(second); {
		if first[i] == second[j] {
			commonValues = append(commonValues, second[j])
			i++
			j++
		} else if first[i] < second[j] {
			i++
		} else {
			j++
		}
	}
	return commonValues
}
