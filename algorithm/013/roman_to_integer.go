package main

import "fmt"

func romanToInt(s string) int {
	dict := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
	num := 0
	for i := range s {
		if (string(s[i]) == "I" || string(s[i]) == "X" || string(s[i]) == "C") && i+1 < len(s) {
			if _, ok := dict[s[i:i+2]]; ok {
				num -= dict[s[i:i+1]]
				continue
			}
		}
		if _, ok := dict[s[i:i+1]]; ok {
			num += dict[s[i:i+1]]
		}
		i += 1
	}
	return num
}

func main() {
	fmt.Println(romanToInt("IIII"))
}
