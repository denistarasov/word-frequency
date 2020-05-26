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

func benchmarkHelper(b *testing.B, isBaseline bool) {
	text := strings.Repeat(strings.Repeat("a b c d ", 50)+"\n", 10000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c := NewCounter()
		if isBaseline {
			_ = c.countBaseline(strings.NewReader(text))
		} else {
			_ = c.countGoroutines(strings.NewReader(text))
		}
	}
}

func BenchmarkCounter_CountBaseline(b *testing.B) {
	benchmarkHelper(b, true)
}

func BenchmarkCounter_CountGoroutines(b *testing.B) {
	benchmarkHelper(b, false)
}
