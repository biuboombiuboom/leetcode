package problem

import "fmt"

func Print() {
	begin := "hit"
	end := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Printf("%v", ladderLength(begin, end, wordList))
}
