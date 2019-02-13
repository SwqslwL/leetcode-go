package main

import "fmt"

func strStr(haystack string, needle string) int {
	length := len(needle)
	if needle == "" {
		return 0
	}

	for i := range haystack {
		if i+length > len(haystack) {
			break
		}

		if haystack[i:i+length] == needle {
			return i
		}
	}
	return -1
}

func main() {
	haystack := "a"
	needle := "a"
	fmt.Println(strStr(haystack, needle))
}
