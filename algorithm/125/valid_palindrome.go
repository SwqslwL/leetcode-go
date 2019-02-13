package main

import "fmt"

func isPalindrome(s string) bool {
	for head, tail := 0, len(s)-1; head <= tail; {
		if s[head] < 48 || s[head] > 57 && s[head] < 65 ||
			s[head] > 90 && s[head] < 97 || s[head] > 122 {
			head += 1
			continue
		}
		if s[tail] < 48 || s[tail] > 57 && s[tail] < 65 ||
			s[tail] > 90 && s[tail] < 97 || s[tail] > 122 {
			tail -= 1
			continue
		}
		fmt.Println(string(s[head]), string(s[tail]))
		if s[head] == s[tail] {
			head += 1
			tail -= 1
		} else {
			//low case
			if s[head] >= 65 && s[head] <= 90 {
				if s[head]+32 == s[tail] {
					head += 1
					tail -= 1
				} else {
					return false
				}
				//high case
			} else if s[head] >= 97 && s[head] <= 122 {
				if s[head]-32 == s[tail] {
					head += 1
					tail -= 1
				} else {
					return false
				}
			}
		}
	}
	return true
}

func isParlindrome2(s string) bool {

}
func main() {
	s := " "
	fmt.Println(isPalindrome(s))
}
