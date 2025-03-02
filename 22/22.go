package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	aoc "aoc"
)

func main() {
	var expectedResult1 int64 = 32179
	var expectedResult2 int64 = 30498
	day := "22"

	blocks, err := aoc.ReadFileSplitBy("input.txt", "\n\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	startPart1 := time.Now()
	resultPartOne := PartOne(blocks)
	fmt.Printf("\nDay_%s Part 1 result: %d in %s\n", day, resultPartOne, time.Since(startPart1))
	startPart2 := time.Now()
	resultPartTwo := PartTwo(blocks)
	fmt.Printf("\nDay_%s Part 2 result: %d in %s\n", day, resultPartTwo, time.Since(startPart2))

	if resultPartOne != expectedResult1 || resultPartTwo != expectedResult2 {
		fmt.Println("Incorrect result")
	} else {
		fmt.Println("Success")
	}
}

func PartOne(blocks []string) int64 {
	player := make([]aoc.Queue[int], 2)
	player[0] = getDeck(blocks[0])
	player[1] = getDeck(blocks[1])

	winner, _ := playNewGame(player)
	return getScore(winner)
}

func PartTwo(blocks []string) int64 {
	player := make([]aoc.Queue[int], 2)
	player[0] = getDeck(blocks[0])
	player[1] = getDeck(blocks[1])

	winner, _ := playNewRecursiveGame(player, 1)
	return getScore(winner)
}

func playNewGame(player []aoc.Queue[int]) (aoc.Queue[int], int) {
	for !player[0].IsEmpty() && !player[1].IsEmpty() {
		card1, _ := player[0].Dequeue()
		card2, _ := player[1].Dequeue()

		if card1 > card2 {
			player[0].Enqueue(card1)
			player[0].Enqueue(card2)
		} else {
			player[1].Enqueue(card2)
			player[1].Enqueue(card1)
		}
	}
	if player[0].Count() > 0 {
		return player[0], 0
	} else {
		return player[1], 1
	}
}

func roundSeenBefore(player []aoc.Queue[int], cache []map[string]bool) bool {
	key0 := getKey(player[0])
	key1 := getKey(player[1])

	if cache[0][key0] || cache[1][key1] {
		return true
	}

	cache[0][key0] = true
	cache[1][key1] = true

	return false
}

func playNewRecursiveGame(player []aoc.Queue[int], game int) (aoc.Queue[int], int) {
	cache := make([]map[string]bool, 2)
	cache[0] = make(map[string]bool)
	cache[1] = make(map[string]bool)
	round := 0
	for !player[0].IsEmpty() && !player[1].IsEmpty() {
		round++

		if roundSeenBefore(player, cache) {
			return nil, 0
		}

		card1, _ := player[0].Dequeue()
		card2, _ := player[1].Dequeue()

		var winner int
		if player[0].Count() >= card1 && player[1].Count() >= card2 {
			newPlayer := make([]aoc.Queue[int], 2)
			newPlayer[0] = slices.Clone(player[0][:card1])
			newPlayer[1] = slices.Clone(player[1][:card2])

			_, winner = playNewRecursiveGame(newPlayer, game+1)
		} else {
			if card1 > card2 {
				winner = 0
			} else {
				winner = 1
			}
		}
		if winner == 0 {
			player[0].Enqueue(card1)
			player[0].Enqueue(card2)
		} else {
			player[1].Enqueue(card2)
			player[1].Enqueue(card1)
		}
	}

	if player[0].Count() > 0 {
		return player[0], 0
	} else {
		return player[1], 1
	}
}

func getDeck(cards string) aoc.Queue[int] {
	var q aoc.Queue[int]
	for _, card := range strings.Split(cards, "\n")[1:] {
		c, err := strconv.Atoi(card)
		if err != nil {
			panic(fmt.Sprintf("Error convering card to int: card:%s, err:%s", card, err))
		}
		q.Enqueue(c)
	}
	return q
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func getKey(q aoc.Queue[int]) string {
	return arrayToString(q, ",")
}

func getScore(q aoc.Queue[int]) int64 {
	var score int64 = 0
	multiplier := q.Count()
	for !q.IsEmpty() {
		card, _ := q.Dequeue()
		score += int64(card * multiplier)
		multiplier--
	}
	return score
}
