package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLine(t *testing.T) {
	pair, err := parseLine("2-4,6-8")
	assert.Nil(t, err)
	assert.Equal(t, 2, pair.First.Start)
	assert.Equal(t, 4, pair.First.End)
}

func TestAssignmentContains(t *testing.T) {
	pair := Pair{
		First:  Assignment{Start: 2, End: 8},
		Second: Assignment{Start: 3, End: 7},
	}

	assert.True(t, pair.First.Contains(pair.Second))

	pair = Pair{
		First:  Assignment{Start: 6, End: 6},
		Second: Assignment{Start: 4, End: 6},
	}

	assert.True(t, pair.First.Contains(pair.Second))

	pair = Pair{
		First:  Assignment{Start: 6, End: 8},
		Second: Assignment{Start: 7, End: 9},
	}

	assert.False(t, pair.First.Contains(pair.Second))
}

func TestAssignmentOverlaps(t *testing.T) {
	pair := Pair{
		First:  Assignment{Start: 2, End: 8},
		Second: Assignment{Start: 3, End: 9},
	}

	assert.True(t, pair.First.Overlaps(pair.Second))

	pair = Pair{
		First:  Assignment{Start: 6, End: 6},
		Second: Assignment{Start: 4, End: 6},
	}

	assert.True(t, pair.First.Contains(pair.Second))

	pair = Pair{
		First:  Assignment{Start: 6, End: 8},
		Second: Assignment{Start: 4, End: 5},
	}

	assert.False(t, pair.First.Contains(pair.Second))
}
