package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Assignment struct {
	Start int
	End   int
}

func (a Assignment) Contains(b Assignment) bool {
	if a.Start <= b.Start {
		if a.End >= b.End {
			return true
		}
	}

	if b.Start <= a.Start {
		if b.End >= a.End {
			return true
		}
	}

	return false
}

func (a Assignment) Overlaps(b Assignment) bool {
	if a.Start <= b.Start && a.End >= b.Start {
		return true
	}

	if b.Start <= a.Start && b.End >= a.Start {
		return true
	}

	return false
}

type Pair struct {
	First  Assignment
	Second Assignment
}

func readInput() ([]string, error) {
	var file *os.File
	var lines []string
	var err error

	if file, err = os.Open("input.txt"); err != nil {
		return lines, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			lines = append(lines, line)
		}
	}

	return lines, scanner.Err()
}

func parseLines(lines []string) ([]Pair, error) {
	pairs := make([]Pair, 0, len(lines))

	for _, line := range lines {
		pair, err := parseLine(line)

		if err != nil {
			return pairs, err
		}

		pairs = append(pairs, pair)
	}

	return pairs, nil
}

func parseLine(line string) (Pair, error) {
	parts := strings.Split(line, ",")
	firstParts := strings.Split(parts[0], "-")
	secondParts := strings.Split(parts[1], "-")

	var pair Pair
	var err error

	if pair.First.Start, err = strconv.Atoi(firstParts[0]); err != nil {
		return pair, err
	}

	if pair.First.End, err = strconv.Atoi(firstParts[1]); err != nil {
		return pair, err
	}

	if pair.Second.Start, err = strconv.Atoi(secondParts[0]); err != nil {
		return pair, err
	}

	if pair.Second.End, err = strconv.Atoi(secondParts[1]); err != nil {
		return pair, err
	}

	return pair, nil
}

func main() {
	lines, err := readInput()

	if err != nil {
		log.Fatal(err)
	}

	pairs, err := parseLines(lines)

	if err != nil {
		log.Fatal(err)
	}

	count := 0

	for _, pair := range pairs {
		if pair.First.Contains(pair.Second) || pair.Second.Contains(pair.First) {
			count += 1
		}
	}

	fmt.Printf("Found %d pairs where one range fully contains the other.\n", count)

	count = 0

	for _, pair := range pairs {
		if pair.First.Overlaps(pair.Second) {
			count += 1
		}
	}

	fmt.Printf("Found %d pairs where one range overlaps the other.\n", count)
}
