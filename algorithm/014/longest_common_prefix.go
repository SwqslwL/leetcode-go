package main

func longestCommonPrefix(strs []string) string {
	common := ""
	for i := range strs {
		if i == 0 {
			common = strs[0]
			continue
		}
		for j := range common {
			if j == len(strs[i]) || common[j] != strs[i][j] {
				common = common[:j]
				break
			}
		}
	}
	return common
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	longestCommonPrefix(strs)
}
