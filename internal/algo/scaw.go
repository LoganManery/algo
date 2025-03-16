package algo

import ()

// substring with Concatenation of All Words

func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return []int{}
	}

	result := []int{}
	wordLen := len(words[0])
	wordsCount := len(words)
	totalLen := wordLen * wordsCount

	// if the total length of all words is greater than s, return empty
	if totalLen > len(s) {
		return []int{}
	}

	// Create a map to store frequency of each word
	wordFreq := make(map[string]int)
	for _, word := range words {
		wordFreq[word]++
	}

	// Iterate through possible starting positions
	for start := 0; start <= len(s) - totalLen; start++ {
		// Create a copy of the word frequency map for each starting position
		seen := make(map[string]int)
		valid := true

		// Check each word-sized chunk in the current window
		for i := 0; i < wordsCount; i++ {
			// Extract the next word-sized chunk
			nextWord := s[start+i*wordLen : start+(i+1)*wordLen]

			// check if the chunk is a valid word in out words array
			if count, exists := wordFreq[nextWord]; !exists {
				valid = false
				break
			} else {
				// increment seen count for this word
				seen[nextWord]++

				// Check if we've seen this word more times than it appears in words array
				if seen[nextWord] > count {
					valid = false
					break
				}
			}
		}

		// If valid concatenation found, add the starting index to result
		if valid {
			result = append(result, start)
		}
	}
	return result
}