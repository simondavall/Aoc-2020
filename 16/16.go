package main

import (
	"cmp"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"

	aoc "aoc"
)

func main() {
	var expectedResult1 int64 = 32835
	var expectedResult2 int64 = 514662805187
	day := "16"

	blocks, err := aoc.ReadFileSplitBy("input.txt", "\n\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	input, err := parseInput(blocks)
	if err != nil {
		fmt.Println(err)
		return
	}

	startPart1 := time.Now()
	resultPartOne := PartOne(input)
	fmt.Printf("\nDay_%s Part 1 result: %d in %s\n", day, resultPartOne, time.Since(startPart1))
	startPart2 := time.Now()
	resultPartTwo := PartTwo(input)
	fmt.Printf("\nDay_%s Part 2 result: %d in %s\n", day, resultPartTwo, time.Since(startPart2))

	if resultPartOne != expectedResult1 || resultPartTwo != expectedResult2 {
		fmt.Println("Incorrect result")
	} else {
		fmt.Println("Success")
	}
}

type Range struct {
	lower int
	upper int
}

type Field struct {
	id     int
	name   string
	ranges []Range
}

type Input struct {
	fields    []Field
	my_ticket []int
	nearby    [][]int
}

func PartOne(input Input) int64 {
	var tally int64 = 0

	for _, ticket := range input.nearby {
		invalidNumber := getInvalidField(ticket, input.fields)
		if invalidNumber > 0 {
			tally += int64(invalidNumber)
		}
	}

	return tally
}

func PartTwo(input Input) int64 {
	whereValid := func(ticket []int) bool { return isTicketValid(ticket, input.fields) }
	validTickets := aoc.Filter(input.nearby, whereValid)

	var validFields [][][]string
	for _, ticket := range validTickets {
		var validIdsForTicket [][]string
		for _, n := range ticket {
			var validIdsForN []string
			fieldsForValue := validFieldsForValue(n, input.fields)
			for _, f := range fieldsForValue {
				validIdsForN = append(validIdsForN, f.name)
			}
			sort.Strings(validIdsForN)
			validIdsForTicket = append(validIdsForTicket, validIdsForN)
		}
		validFields = append(validFields, validIdsForTicket)
	}

	transposed := aoc.Transpose(validFields)

	type Result struct {
		col_id int
		result []string
	}
	var result []Result
	for idx, vals := range transposed {
		ticket := vals[0]
		for _, items := range vals[1:] {
			ticket = intersects(ticket, items)
		}
		result = append(result, Result{idx, ticket})
	}
	lenCmp := func(a, b Result) int {
		return cmp.Compare(len(a.result), len(b.result))
	}
	slices.SortFunc(result, lenCmp)

	for i := 0; i < len(result)-1; i++ {
		for j := i + 1; j < len(result); j++ {
			result[j].result = removeValue(result[j].result, result[i].result[0])
		}
	}

	var tally int64 = 1
	for _, res := range result {
		if strings.HasPrefix(res.result[0], "departure") {
			tally *= int64(input.my_ticket[res.col_id])
		}
	}

	return tally
}

func intersects(first []string, second []string) []string {
	commonValues := make([]string, 0)
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

func getInvalidField(ticket []int, fields []Field) int {
	for _, n := range ticket {
		isValid := false
		for _, field := range fields {
			if isValueValid(n, field) {
				isValid = true
				break
			}
		}
		if !isValid {
			return n
		}
	}
	return -1
}

func isValueValid(value int, field Field) bool {
	for _, rng := range field.ranges {
		if rng.lower <= value && value <= rng.upper {
			return true
		}
	}
	return false
}

func isTicketValid(ticket []int, fields []Field) bool {
	return getInvalidField(ticket, fields) < 0
}

func validFieldsForValue(value int, fields []Field) []Field {
	whereValid := func(field Field) bool {
		is_valid := false
		for _, rng := range field.ranges {
			if rng.lower <= value && value <= rng.upper {
				is_valid = true
				break
			}
		}
		return is_valid
	}
	return aoc.Filter(fields, whereValid)
}

func removeValue(s []string, str string) []string {
	newArray := make([]string, len(s)-1)
	idx := 0
	for _, val := range s {
		if val == str {
			continue
		}
		newArray[idx] = val
		idx++
	}
	return newArray
}

func parseInput(blocks []string) (Input, error) {
	var input Input

	rawFields := strings.Split(blocks[0], "\n")
	for idx, rawField := range rawFields {
		fields := strings.Split(rawField, ": ")
		var fieldRanges []Range
		for _, rawRanges := range strings.Split(fields[1], " or ") {
			items := strings.Split(rawRanges, "-")
			lower, err := strconv.Atoi(items[0])
			if err != nil {
				return input, err
			}
			upper, err := strconv.Atoi(items[1])
			if err != nil {
				return input, err
			}
			fieldRanges = append(fieldRanges, Range{lower, upper})
		}
		input.fields = append(input.fields, Field{idx, fields[0], fieldRanges})
	}

	rawMyTicket := strings.Split(blocks[1], "\n")
	for _, val := range strings.Split(rawMyTicket[1], ",") {
		ticket, err := strconv.Atoi(val)
		if err != nil {
			return input, err
		}
		input.my_ticket = append(input.my_ticket, ticket)
	}

	rawNearby := strings.Split(blocks[2], "\n")
	for _, nearby := range rawNearby[1:] {
		if nearby == "" {
			continue
		}
		numbers := strings.Split(nearby, ",")
		var ticketNumbers []int
		for _, strN := range numbers {
			n, err := strconv.Atoi(strN)
			if err != nil {
				return input, err
			}
			ticketNumbers = append(ticketNumbers, n)
		}
		input.nearby = append(input.nearby, ticketNumbers)
	}
	return input, nil
}
