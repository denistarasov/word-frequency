package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCounter_simple(t *testing.T) {
	c := NewCounter()
	err := c.Count(strings.NewReader("b a c c b"))
	assert.NoError(t, err)

	wordCounts := c.GetMostCommon(-1)

	assert.Equal(t, 3, len(wordCounts))

	assert.Equal(t, "b", wordCounts[0].Word)
	assert.Equal(t, uint64(2), wordCounts[0].Count)

	assert.Equal(t, "c", wordCounts[1].Word)
	assert.Equal(t, uint64(2), wordCounts[1].Count)

	assert.Equal(t, "a", wordCounts[2].Word)
	assert.Equal(t, uint64(1), wordCounts[2].Count)
}

func TestCounter_GetMostCommon(t *testing.T) {
	c := NewCounter()

	wordCounts := c.GetMostCommon(-1)
	assert.Equal(t, 0, len(wordCounts))

	err := c.Count(strings.NewReader("b a c c b"))
	assert.NoError(t, err)

	wordCounts = c.GetMostCommon(1)
	assert.Equal(t, 1, len(wordCounts))
	assert.Equal(t, "b", wordCounts[0].Word)
	assert.Equal(t, uint64(2), wordCounts[0].Count)
}
