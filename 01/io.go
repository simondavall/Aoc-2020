package main

import (
	"bufio"
	"os"
)

func readLines(path string) ([]string, error) {
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
