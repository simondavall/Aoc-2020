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
	var expectedResult1 int64 = 206
	var expectedResult2 int64 = 123
	day := "04"

	lines, err := aoc.ReadFileSplitBy("input.txt", "\n\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	startPart1 := time.Now()
	resultPartOne := PartOne(lines)
	fmt.Printf("\nDay_%s Part 1 result: %d in %s\n", day, resultPartOne, time.Since(startPart1))
	startPart2 := time.Now()
	resultPartTwo := PartTwo(lines)
	fmt.Printf("\nDay_%s Part 2 result: %d in %s\n", day, resultPartTwo, time.Since(startPart2))

	if resultPartOne != expectedResult1 || resultPartTwo != expectedResult2 {
		fmt.Println("Incorrect result")
	} else {
		fmt.Println("Success")
	}
}

type Passport struct {
	BirthYear  int
	IssueYear  int
	ExpYear    int
	Height     string
	HairColour string
	EyeColour  string
	PassportId string
	CountryId  int
}

func PartOne(input []string) int64 {
	var tally int64 = 0

	passports, err := getPassports(input)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	for _, p := range passports {
		if isPassportValid(p) {
			tally++
		}
	}

	return tally
}

func PartTwo(input []string) int64 {
	var tally int64 = 0

	passports, err := getPassports(input)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	for _, p := range passports {
		if isPassportValidStrict(p) {
			tally++
		}
	}

	return tally
}

func isPassportValid(p Passport) bool {
	if p.BirthYear == 0 || p.EyeColour == "" || p.HairColour == "" ||
		p.ExpYear == 0 || p.Height == "" || p.IssueYear == 0 || p.PassportId == "" {
		return false
	}

	return true
}

var (
	validColourChars []string = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	validEyecolours  []string = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	validMetrics     []string = []string{"cm", "in"}
)

func isPassportValidStrict(p Passport) bool {
	if p.BirthYear < 1920 || p.BirthYear > 2002 {
		return false
	}
	if p.IssueYear < 2010 || p.IssueYear > 2020 {
		return false
	}
	if p.ExpYear < 2020 || p.ExpYear > 2030 {
		return false
	}

	if len(p.Height) < 4 {
		return false
	}

	metric := p.Height[(len(p.Height) - 2):]
	if !slices.Contains(validMetrics, metric) {
		return false
	}
	height, err := strconv.Atoi(p.Height[:len(p.Height)-2])
	if err != nil {
		return false
	}
	if metric == "cm" && (height < 150 || height > 193) {
		return false
	}
	if metric == "in" && (height < 59 || height > 76) {
		return false
	}

	if p.HairColour == "" || p.HairColour[0] != '#' {
		return false
	}
	for _, ch := range p.HairColour[1:] {
		if !slices.Contains(validColourChars, string(ch)) {
			return false
		}
	}

	if p.EyeColour == "" || !slices.Contains(validEyecolours, p.EyeColour) {
		return false
	}

	if p.PassportId == "" {
		return false
	}

	_, err1 := strconv.Atoi(p.PassportId)
	if err1 != nil || len(p.PassportId) != 9 {
		return false
	}

	return true
}

func getPassports(input []string) ([]Passport, error) {
	var passports []Passport
	for _, str := range input {
		lines := strings.Split(str, "\n")
		var passport Passport
		for _, line := range lines {
			items := strings.Split(line, " ")
			for _, item := range items {
				entry := strings.Split(item, ":")
				err := populate(entry[0], entry[1], &passport)
				if err != nil {
					return nil, err
				}
			}
		}
		passports = append(passports, passport)
	}
	return passports, nil
}

func populate(k string, v string, passport *Passport) error {
	switch k {
	case "byr":
		byr, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		passport.BirthYear = byr
	case "iyr":
		iyr, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		passport.IssueYear = iyr
	case "eyr":
		eyr, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		passport.ExpYear = eyr
	case "hgt":
		passport.Height = v
	case "hcl":
		passport.HairColour = v
	case "ecl":
		passport.EyeColour = v
	case "pid":
		passport.PassportId = v
	case "cid":
		cid, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		passport.CountryId = cid
	default:
		panic("Ooops found invalid key: " + k)

	}
	return nil
}
