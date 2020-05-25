package main

import (
	"bufio"
	"github.com/euskadi31/go-tokenizer"
	"io"
	"sort"
)

type Counter struct {
	wordToCount map[string]uint64
}

type WordCount struct {
	Word  string
	Count uint64
}

func NewCounter() *Counter {
	return &Counter{
		wordToCount: make(map[string]uint64),
	}
}

func (c *Counter) Count(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	t := tokenizer.New()
	for scanner.Scan() {
		line := scanner.Text()
		tokens := t.Tokenize(line)
		c.update(tokens)
	}

	err := scanner.Err()
	return err
}

func (c *Counter) update(words []string) {
	for _, word := range words {
		c.wordToCount[word]++
	}
}

// Return a slice of the n most common elements and their counts
// from the most common to the least.
// If n < 0, returns all elements in the counter.
func (c *Counter) GetMostCommon(n int) []WordCount {
	result := make([]WordCount, 0)
	for word, count := range c.wordToCount {
		result = append(result, WordCount{
			Word:  word,
			Count: count,
		})
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].Count < result[j].Count {
			return false
		}
		if result[i].Count == result[j].Count && result[i].Word > result[j].Word {
			return false
		}
		return true
	})

	if n >= 0 {
		return result[:n]
	}
	return result
}
