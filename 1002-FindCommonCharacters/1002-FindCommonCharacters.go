package main

import (
	"fmt"
	"sort"
)

/*
Given a string array words, return an array of all characters that show up in all strings within the words (including duplicates).
You may return the answer in any order.



Example 1:

Input: words = ["bella","label","roller"]
Output: ["e","l","l"]
Example 2:

Input: words = ["cool","lock","cook"]
Output: ["c","o"]


Constraints:

1 <= words.length <= 100
1 <= words[i].length <= 100
words[i] consists of lowercase English letters.
*/
func commonChars(words []string) []string {
	count_map := make(map[byte]int)
	tmp_map := make(map[byte]int)
	for i, word := range words {
		for j := 0; j < len(word); j++ {
			if i == 0 {
				if val, ok := count_map[word[j]]; ok {
					count_map[word[j]] = val + 1
				} else {
					count_map[word[j]] = 1
				}
			} else {
				if val, ok := tmp_map[word[j]]; ok {
					tmp_map[word[j]] = val + 1
				} else {
					tmp_map[word[j]] = 1
				}

			}

		}
		if i != 0 {
			new_count_map := make(map[byte]int)
			for tmp_val, num := range tmp_map {
				if count_num, ok := count_map[tmp_val]; ok {
					new_count_map[tmp_val] = IntMin(count_num, num)
				}
			}
			count_map = new_count_map
			tmp_map = make(map[byte]int)
		}

	}
	var ans []string
	for char, num := range count_map {
		// fmt.Println(string(char))
		for i := 0; i < num; i++ {
			ans = append(ans, string(char))
		}

	}
	sort.Strings(ans)
	// fmt.Println(count_map)
	return ans
}

func IntMin(x int, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
func main() {
	// var words []string ["cool","lock","cook"]
	words := []string{"cool", "lock", "cook"}
	ans := commonChars(words)
	fmt.Println(ans)
}
