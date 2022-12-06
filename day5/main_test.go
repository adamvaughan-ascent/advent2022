package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	stack := Stack{}
	assert.Equal(t, 0, stack.Len())

	stack = stack.Push("hi")
	stack = stack.Push("hello")
	assert.Equal(t, 2, stack.Len())

	stack, v := stack.Pop()
	assert.Equal(t, "hello", v)
	assert.Equal(t, 1, stack.Len())

	stack, v = stack.Pop()
	assert.Equal(t, "hi", v)
	assert.Equal(t, 0, stack.Len())

	stack = stack.PushMany([]interface{}{"a", "b", "c"})
	assert.Equal(t, 3, stack.Len())

	stack, v = stack.PopMany(2)
	assert.Equal(t, []interface{}{"b", "c"}, v)
	assert.Equal(t, 1, stack.Len())

	stack, v = stack.Pop()
	assert.Equal(t, "a", v)

	assert.Panics(t, func() { stack.Pop() })
}

func TestParseInput(t *testing.T) {
	input := parseInput([]string{
		"        [C] [B] [H]                ",
		"move 5 from 4 to 7",
	})
	assert.Equal(t, 0, input.Stacks[0].Len())
	assert.Equal(t, 0, input.Stacks[1].Len())
	assert.Equal(t, 1, input.Stacks[2].Len())
	assert.Equal(t, 1, input.Stacks[3].Len())
	assert.Equal(t, 1, input.Stacks[4].Len())
	assert.Equal(t, 0, input.Stacks[5].Len())
	assert.Equal(t, 0, input.Stacks[6].Len())
	assert.Equal(t, 0, input.Stacks[7].Len())
	assert.Equal(t, 0, input.Stacks[8].Len())

	assert.Len(t, input.Moves, 1)
	assert.Equal(t, 5, input.Moves[0].Count)
	assert.Equal(t, 3, input.Moves[0].From)
	assert.Equal(t, 6, input.Moves[0].To)
}

func TestExecuteMoves(t *testing.T) {
	lines := []string{
		"    [D]                            ",
		"[N] [C]                            ",
		"[Z] [M] [P]                        ",
		" 1   2   3   4   5   6   7   8   9 ",
		"",
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}

	input := parseInput(lines)
	input = executeMoves(input)

	assert.Equal(t, "M", input.Stacks[0].Peek())
	assert.Equal(t, "C", input.Stacks[1].Peek())
	assert.Equal(t, "D", input.Stacks[2].Peek())
}
