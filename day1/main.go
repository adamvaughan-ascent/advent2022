package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
	Calories      []int
	TotalCalories int
}

func readInput() ([]Elf, error) {
	var file *os.File
	var elves []Elf
	var elf Elf
	var err error

	if file, err = os.Open("input.txt"); err != nil {
		return elves, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if elf.TotalCalories > 0 {
				elves = append(elves, elf)
			}

			elf = Elf{}
			continue
		}

		i, err := strconv.Atoi(line)

		if err != nil {
			log.Fatal(err)
		}

		elf.Calories = append(elf.Calories, i)
		elf.TotalCalories += i
	}

	if elf.TotalCalories > 0 {
		elves = append(elves, elf)
	}

	return elves, scanner.Err()
}

func main() {
	var elves []Elf
	var err error

	if elves, err = readInput(); err != nil {
		log.Fatal(err)
	}

	sort.SliceStable(elves, func(i, j int) bool {
		return elves[i].TotalCalories > elves[j].TotalCalories
	})

	fmt.Printf("The elf with the most calories is carrying %d calories.\n", elves[0].TotalCalories)

	topThree := elves[0].TotalCalories + elves[1].TotalCalories + elves[2].TotalCalories
	fmt.Printf("The top three elves are carrying %d calories.\n", topThree)
}
