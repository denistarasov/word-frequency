package main

import (
	"bufio"
	"github.com/euskadi31/go-tokenizer"
	"io"
	"sort"
	"sync"
)

type Counter struct {
	wordToCount map[string]uint64
	lines       chan string
	wg          sync.WaitGroup
	m           sync.Mutex
}

type WordCount struct {
	Word  string
	Count uint64
}

func NewCounter() *Counter {
	return &Counter{
		wordToCount: make(map[string]uint64),
		lines:       make(chan string),
		wg:          sync.WaitGroup{},
		m:           sync.Mutex{},
	}
}

func (c *Counter) Count(r io.Reader) error {
	return c.countGoroutines(r)
}

func (c *Counter) countBaseline(r io.Reader) error {
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

func (c *Counter) countGoroutines(r io.Reader) error {
	for i := 0; i != 3; i++ {
		c.wg.Add(1)
		go func() {
			t := tokenizer.New()
			for line := range c.lines {
				tokens := t.Tokenize(line)
				c.update(tokens)
			}
			c.wg.Done()
		}()
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		c.lines <- line
	}
	close(c.lines)

	err := scanner.Err()
	c.wg.Wait()
	return err
}

func (c *Counter) update(words []string) {
	c.m.Lock()
	defer c.m.Unlock()
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
