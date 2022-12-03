package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type Group struct {
	Rucksacks  []Rucksack
	SharedItem rune
}

func (g Group) Priority() int {
	if int(g.SharedItem) > 90 {
		return int(g.SharedItem) - 96
	}

	return int(g.SharedItem) - 38
}

type Rucksack struct {
	Contents     string
	Compartments []string
}

func readInput() ([]Group, error) {
	var file *os.File
	var groups []Group
	var group Group
	var err error

	if file, err = os.Open("input.txt"); err != nil {
		return groups, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		group.Rucksacks = append(group.Rucksacks, parseInput(line))

		if len(group.Rucksacks) == 3 {
			group.SharedItem = findCommonItem(group)
			groups = append(groups, group)
			group = Group{}
		}
	}

	return groups, scanner.Err()
}

func parseInput(line string) Rucksack {
	i := len(line) / 2

	return Rucksack{
		Contents:     line,
		Compartments: []string{line[0:i], line[i:]},
	}
}

func findCommonItem(group Group) rune {
	indexes := make([]int, len(group.Rucksacks))
	runes := make([][]rune, len(group.Rucksacks))

	for i, rucksack := range group.Rucksacks {
		r := []rune(rucksack.Contents)
		sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
		runes[i] = r
	}

	for {
		if indexes[0] >= len(runes[0]) || indexes[1] >= len(runes[1]) || indexes[2] >= len(runes[2]) {
			break
		}

		a := runes[0][indexes[0]]
		b := runes[1][indexes[1]]
		c := runes[2][indexes[2]]

		if a == b && a == c {
			return a
		}

		if a < b || a < c {
			indexes[0] += 1
		}

		if b < a || b < c {
			indexes[1] += 1
		}

		if c < a || c < b {
			indexes[2] += 1
		}
	}

	panic("didn't find a common item")
}

func sumOfPriorities(groups []Group) int {
	sum := 0

	for _, group := range groups {
		sum += group.Priority()
	}

	return sum
}

func main() {
	groups, err := readInput()

	if err != nil {
		log.Fatal(err)
	}

	sum := sumOfPriorities(groups)
	fmt.Printf("The sum of priorities is %d\n", sum)
}
