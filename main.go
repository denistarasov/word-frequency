package main

import (
	"fmt"
	"os"
)

func main() {
	counter := NewCounter()
	err := counter.Count(os.Stdin)
	if err != nil {
		panic(err)
	}

	for _, wordCount := range counter.GetMostCommon(-1) {
		fmt.Println(wordCount.Word, wordCount.Count)
	}
}
