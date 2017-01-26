package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	fmt.Printf("string to test: `%v`\n", s)
	occurrence := make([]int, 1024)
	indexMap := make([]int, 1024)
	longest := 0
	currentLength := 0
	pos := 0
	for {
		idx := s[pos] % 97
		if occurrence[idx] != 1 {
			occurrence[idx] = 1
			indexMap[idx] = pos
			currentLength += 1
			pos++
		} else {
			backToPos := indexMap[idx] + 1
			fmt.Printf("repeats found, current pos: %d, repeating char: %c, back to %d\n", pos, s[pos], backToPos)
			// clear the occurrence map
			occurrence = make([]int, 1024)
			pos = backToPos
			indexMap = make([]int, 1024)

			if currentLength >= longest {
				longest = currentLength
			}
			currentLength = 0
		}

		if pos == len(s) {
			break
		}
	}

	if currentLength >= longest {
		longest = currentLength
		currentLength = 0
	}

	return longest
}

func main() {
	fmt.Printf("%d\n", lengthOfLongestSubstring("pwwkew"))
	fmt.Printf("%d\n", lengthOfLongestSubstring("abcabcbb"))
	fmt.Printf("%d\n", lengthOfLongestSubstring("bbbbb"))
	fmt.Printf("%d\n", lengthOfLongestSubstring("abcad"))
	fmt.Printf("%d\n", lengthOfLongestSubstring("ggububgvfk"))
}
