package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Move string
type Outcome string

const (
	Rock     Move = "A"
	Paper    Move = "B"
	Scissors Move = "C"
)

const (
	Win  Outcome = "Z"
	Lose Outcome = "X"
	Draw Outcome = "Y"
)

func (m1 Move) Beats(m2 Move) bool {
	return (m1 == Rock && m2 == Scissors) || (m1 == Paper && m2 == Rock) || (m1 == Scissors && m2 == Paper)
}

func (m Move) Points() int {
	if m == Rock {
		return 1
	}

	if m == Paper {
		return 2
	}

	return 3
}

type Round struct {
	OpponentMove Move
	YourMove     Move
	Outcome      Outcome
}

func readInput() ([]Round, error) {
	var file *os.File
	var rounds []Round
	var err error

	if file, err = os.Open("input.txt"); err != nil {
		return rounds, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")

		rounds = append(rounds, Round{
			OpponentMove: Move(parts[0]),
			Outcome:      Outcome(parts[1]),
		})
	}

	return rounds, scanner.Err()
}

func getMove(opponentMove Move, outcome Outcome) Move {
	if outcome == Draw {
		return opponentMove
	}

	if outcome == Win {
		if opponentMove == Rock {
			return Paper
		}

		if opponentMove == Paper {
			return Scissors
		}

		return Rock
	}

	if opponentMove == Rock {
		return Scissors
	}

	if opponentMove == Paper {
		return Rock
	}

	return Paper
}

func computeScore(rounds []Round) int {
	var score int

	for _, round := range rounds {
		round.YourMove = getMove(round.OpponentMove, round.Outcome)

		if round.Outcome == Draw {
			score += 3 + round.YourMove.Points()
		} else if round.Outcome == Win {
			score += 6 + round.YourMove.Points()
		} else {
			score += round.YourMove.Points()
		}
	}

	return score
}

func main() {
	var rounds []Round
	var score int
	var err error

	if rounds, err = readInput(); err != nil {
		log.Fatal(err)
	}

	score = computeScore(rounds)
	fmt.Printf("Final score: %d\n", score)
}
