package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	aoc "aoc"
)

type Food struct {
	ingredients []string
	allergens   []string
}

func main() {
	var expectedResult1 int64 = 2493
	var expectedResult2 string = "kqv,jxx,zzt,dklgl,pmvfzk,tsnkknk,qdlpbt,tlgrhdh"
	day := "21"

	lines, err := aoc.ReadLines("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var foods []Food
	for _, line := range lines {
		s := strings.Split(line, " (contains ")
		f := Food{strings.Split(s[0], " "), strings.Split(strings.TrimRight(s[1], ")"), ", ")}
		foods = append(foods, f)
	}

	startPart1 := time.Now()
	resultPartOne := PartOne(foods)
	fmt.Printf("\nDay_%s Part 1 result: %d in %s\n", day, resultPartOne, time.Since(startPart1))
	startPart2 := time.Now()
	resultPartTwo := PartTwo()
	fmt.Printf("\nDay_%s Part 2 result: %s in %s\n", day, resultPartTwo, time.Since(startPart2))

	if resultPartOne != expectedResult1 || resultPartTwo != expectedResult2 {
		fmt.Println("Incorrect result")
	} else {
		fmt.Println("Success")
	}
}

var allergens map[string][]string

func PartOne(foods []Food) int64 {
	var tally int64 = 0

	allergens = make(map[string][]string)
	for _, food := range foods {
		for _, alg := range food.allergens {
			val, exists := allergens[alg]
			if !exists {
				allergens[alg] = food.ingredients
				continue
			}
			allergens[alg] = aoc.Intersects(food.ingredients, val)
		}
	}

	resolved := false
	for !resolved {
		resolved = true
		for i, ing1 := range allergens {
			for j, ings := range allergens {
				if i == j || len(ing1) != 1 {
					continue
				}
				idx := slices.Index(ings, ing1[0])
				if idx == -1 {
					continue
				}
				allergens[j] = slices.Concat(allergens[j][:idx], allergens[j][idx+1:])
				resolved = false
			}
		}
	}

	ingredients := make(map[string]string)
	for k, v := range allergens {
		ingredients[v[0]] = k
	}

	for _, food := range foods {
		for _, ing := range food.ingredients {
			if len(ingredients[ing]) == 0 {
				tally++
			}
		}
	}

	return tally
}

func PartTwo() string {
	var algList []string
	for alg := range allergens {
		algList = append(algList, alg)
	}
	slices.Sort(algList)

	var dangerous []string
	for _, alg := range algList {
		dangerous = append(dangerous, allergens[alg][0])
	}

	return strings.Join(dangerous, ",")
}
