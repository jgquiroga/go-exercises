package main

import "fmt"

// Calculates the frequency of words in an array
func frequency(words []string) map[string]int {
	results := make(map[string]int)
	wordsLen := len(words)

	for i := 0; i < wordsLen; i++ {
		results[words[i]] = results[words[i]] + 1
	}

	return results
}

func main() {
	words := []string{
		"This", "is", "a", "list", "of", "words", ".", "some", "words",
		"can", "appear", "more", "than", "once", ".", "this",
		"is", "just", "a", "test",
	}

	frequencyResult := frequency(words)

	for key, value := range frequencyResult {
		fmt.Println("key: ", key)
		fmt.Println("value: ", value)
		fmt.Println("----")
	}
}
