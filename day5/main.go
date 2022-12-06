package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Stack struct {
	values []interface{}
}

func (s Stack) Peek() interface{} {
	return s.values[len(s.values)-1]
}

func (s Stack) Push(v interface{}) Stack {
	s.values = append(s.values, v)
	return s
}

func (s Stack) PushMany(v []interface{}) Stack {
	s.values = append(s.values, v...)
	return s
}

func (s Stack) Pop() (Stack, interface{}) {
	if len(s.values) == 0 {
		panic("popping empty stack")
	}

	v := s.values[len(s.values)-1]
	s.values = s.values[0 : len(s.values)-1]
	return s, v
}

func (s Stack) PopMany(count int) (Stack, []interface{}) {
	fmt.Printf("popping %d from %v\n", count, s.values)
	if len(s.values) < count {
		panic("not enough items on stack")
	}

	v := s.values[len(s.values)-count:]
	s.values = s.values[0 : len(s.values)-count]
	return s, v
}

func (s Stack) Reverse() Stack {
	for i, j := 0, len(s.values)-1; i < j; i, j = i+1, j-1 {
		s.values[i], s.values[j] = s.values[j], s.values[i]
	}

	return s
}

func (s Stack) Len() int {
	return len(s.values)
}

type Move struct {
	Count int
	From  int
	To    int
}

type Input struct {
	Stacks []Stack
	Moves  []Move
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

func parseInput(lines []string) Input {
	input := Input{
		Stacks: make([]Stack, 9),
	}

	stackRegex, _ := regexp.Compile(`([\[\]A-Z ]{3}) ([\[\]A-Z ]{3}) ([\[\]A-Z ]{3}) ([\[\]A-Z ]{3}) ([\[\]A-Z ]{3}) ([\[\]A-Z ]{3}) ([\[\]A-Z ]{3}) ([\[\]A-Z ]{3}) ([\[\]A-Z ]{3})`)
	moveRegex, _ := regexp.Compile(`move (\d+) from (\d+) to (\d+)`)

	for _, line := range lines {
		if match := stackRegex.FindStringSubmatch(line); match != nil {
			for i, m := range match[1:] {
				if m != "   " {
					input.Stacks[i] = input.Stacks[i].Push(string([]rune(m)[1]))
				}
			}
		} else if match := moveRegex.FindStringSubmatch(line); match != nil {
			count, _ := strconv.Atoi(match[1])
			from, _ := strconv.Atoi(match[2])
			to, _ := strconv.Atoi(match[3])

			input.Moves = append(input.Moves, Move{Count: count, From: from - 1, To: to - 1})
		}
	}

	for i := 0; i < len(input.Stacks); i += 1 {
		input.Stacks[i].Reverse()
	}

	return input
}

func executeMoves(input Input) Input {
	for _, move := range input.Moves {
		fmt.Printf("Moving %d from %d to %d\n", move.Count, move.From, move.To)

		from := input.Stacks[move.From]
		to := input.Stacks[move.To]

		fmt.Printf("   Before From  %v\n", from)
		fmt.Printf("   Before To    %v\n", to)

		s, v := input.Stacks[move.From].PopMany(move.Count)
		input.Stacks[move.From] = s

		s = input.Stacks[move.To].PushMany(v)
		input.Stacks[move.To] = s

		from = input.Stacks[move.From]
		to = input.Stacks[move.To]

		fmt.Printf("   After From  %v\n", from)
		fmt.Printf("   After To    %v\n", to)
	}

	return input
}

func main() {
	lines, err := readInput()

	if err != nil {
		log.Fatal(err)
	}

	input := parseInput(lines)

	input = executeMoves(input)

	for i := 0; i < len(input.Stacks); i += 1 {
		fmt.Printf("%s", input.Stacks[i].Peek())
	}
}
