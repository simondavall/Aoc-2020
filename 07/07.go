package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	aoc "aoc"
)

type Bag struct {
	name     string
	children []BagChild
}

type BagChild struct {
	name   string
	amount int
}

var (
	bags       = make(map[string]Bag)
	goldCache  = make(map[string]bool)
	countCache = make(map[string]int)
)

func main() {
	var expectedResult1 int64 = 252
	var expectedResult2 int64 = 35487
	day := "07"

	lines, err := aoc.ReadLines("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	startPart1 := time.Now()
	resultPartOne := PartOne(lines)
	fmt.Printf("\nDay_%s Part 1 result: %d in %s\n", day, resultPartOne, time.Since(startPart1))
	startPart2 := time.Now()
	resultPartTwo := PartTwo()
	fmt.Printf("\nDay_%s Part 2 result: %d in %s\n", day, resultPartTwo, time.Since(startPart2))

	if resultPartOne != expectedResult1 || resultPartTwo != expectedResult2 {
		fmt.Println("Incorrect result")
	} else {
		fmt.Println("Success")
	}
}

func PartOne(lines []string) int64 {
	loadBags(lines)

	var tally int64 = 0
	for name := range bags {
		if CanContainGoldBag(name) {
			tally++
		}
	}

	return tally - 1
}

func PartTwo() int64 {
	var tally int = 0

	shinyGold := bags["shiny gold"]
	for _, child := range shinyGold.children {
		tally += GetBagCount(child)
	}

	return int64(tally)
}

func GetBagCount(bagChild BagChild) int {
	if val, ok := countCache[bagChild.name]; ok {
		return bagChild.amount + (val * bagChild.amount)
	}

	var current Bag = bags[bagChild.name]
	childrenCount := 0

	for _, next := range current.children {
		nextcount := GetBagCount(next)
		childrenCount += nextcount
	}

	countCache[bagChild.name] = childrenCount
	return bagChild.amount + (bagChild.amount * childrenCount)
}

func CanContainGoldBag(bagName string) bool {
	if bagName == "shiny gold" {
		return true
	}

	if hasGold, ok := goldCache[bagName]; ok {
		return hasGold
	}

	bag := bags[bagName]
	if bag.children == nil {
		return false
	}

	for _, next := range bag.children {
		if hasGold, ok := goldCache[next.name]; ok {
			if hasGold {
				return true
			}
			continue
		}

		hasGold := CanContainGoldBag(next.name)

		goldCache[next.name] = hasGold
		if hasGold {
			goldCache[bagName] = true
			return true
		}
	}

	goldCache[bagName] = false
	return false
}

func loadBags(lines []string) {
	pattern1 := " (\\d+) ([\\w\\s]+) bags?[,.]"

	for _, line := range lines {
		s := strings.Split(line, " bags contain")
		r := regexp.MustCompile(pattern1)
		matches := r.FindAllStringSubmatch(s[1], -1)
		var children []BagChild
		for _, match := range matches {
			name := match[2]
			amount, _ := strconv.Atoi(match[1])
			children = append(children, BagChild{name, amount})
		}
		var bag Bag = Bag{s[0], children}
		bags[s[0]] = bag
	}
}
