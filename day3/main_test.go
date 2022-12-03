package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	r := parseInput("vJrwpWtwJgWrhcsFMMfFFhFp")
	assert.Equal(t, "vJrwpWtwJgWrhcsFMMfFFhFp", r.Contents)
	assert.Len(t, r.Compartments, 2)
	assert.Equal(t, "vJrwpWtwJgWr", r.Compartments[0])
	assert.Equal(t, "hcsFMMfFFhFp", r.Compartments[1])

	r = parseInput("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL")
	assert.Equal(t, "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", r.Contents)
	assert.Len(t, r.Compartments, 2)
	assert.Equal(t, "jqHRNqRjqzjGDLGL", r.Compartments[0])
	assert.Equal(t, "rsFMfFZSrLrFZsSL", r.Compartments[1])

	r = parseInput("PmmdzqPrVvPwwTWBwg")
	assert.Equal(t, "PmmdzqPrVvPwwTWBwg", r.Contents)
	assert.Len(t, r.Compartments, 2)
	assert.Equal(t, "PmmdzqPrV", r.Compartments[0])
	assert.Equal(t, "vPwwTWBwg", r.Compartments[1])
}

func TestFindCommonItem(t *testing.T) {
	g := Group{
		Rucksacks: []Rucksack{
			{Contents: "vJrwpWtwJgWrhcsFMMfFFhFp"},
			{Contents: "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"},
			{Contents: "PmmdzqPrVvPwwTWBwg"},
		},
	}

	assert.Equal(t, 'r', findCommonItem(g))

	g = Group{
		Rucksacks: []Rucksack{
			{Contents: "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"},
			{Contents: "ttgJtRGJQctTZtZT"},
			{Contents: "CrZsJsPPZsGzwwsLwLmpwMDw"},
		},
	}

	assert.Equal(t, 'Z', findCommonItem(g))
}

func TestSumOfPriorities(t *testing.T) {
	groups := []Group{
		{SharedItem: 'r'},
		{SharedItem: 'Z'},
	}

	assert.Equal(t, 70, sumOfPriorities(groups))
}
