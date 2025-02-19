package aoc

import (
	"bufio"
	"os"
	"strings"
)

func ReadLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		str := scanner.Text()
		if str != "" {
			lines = append(lines, scanner.Text())
		}
	}
	return lines, scanner.Err()
}

func ReadFileSplitBy(path string, sep string) ([]string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var lines []string
	content := string(b)

	for _, str := range strings.Split(content, sep) {
		if str != "" {
			lines = append(lines, str)
		}
	}
	return lines, nil
}
