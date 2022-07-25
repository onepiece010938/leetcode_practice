package main

import (
	"fmt"
	"reflect"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// solution1

	eq := reflect.DeepEqual(convert2map(s), convert2map(t))
	if eq {
		fmt.Println("They're equal.")
		return true
	} else {
		fmt.Println("They're unequal.")
		return false
	}

	// solution2

	// if fmt.Sprint(convert2map(s)) == fmt.Sprint(convert2map(t)) {
	// 	fmt.Print(fmt.Sprintln(convert2map(s)))
	// 	return true
	// } else {
	// 	return false
	// }

	// solution3 (use array is faster)
	//     var counter [26]int
	//     for i := 0; i < len(s); i++ {
	//         counter[s[i] - 'a']++
	//         counter[t[i] - 'a']--
	//     }

	//     for _, v := range counter {
	//         if v != 0 {
	//             return false
	//         }
	//     }
	//     return true

	// solution4 (use map)
	// m := map[byte]int{}

	// for i := 0; i < len(s); i++ {
	//     m[s[i]] = m[s[i]] + 1
	//     m[t[i]] = m[t[i]] - 1
	//   }

	// for _, v := range m {
	//     if v != 0 {
	//       return false
	//     }
	//   }

	//   return true

}

func convert2map(s string) (result map[byte]int) {
	count_map := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if val, ok := count_map[s[i]]; ok {
			count_map[s[i]] = val + 1
		} else {
			count_map[s[i]] = 1
		}
	}
	return count_map
}

func main() {
	s := "anagrama"
	t := "nagaram"
	println(isAnagram(s, t))
}
