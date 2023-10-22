package two_pointers

import (
	"strings"
)

func IsPalindrome(s string) bool {

	return false
}

func IsPalindrome1(s string) bool {
	str := strings.ToLower(s)
	str = RemoveNonAlphanumeric(str)
	if len(str) == 1 {
		return true
	}
	if str == ReverseString(str) {
		return true
	}

	return false
}

func ReverseString(s string) string {
	str := ""
	for i := len(s) - 1; i >= 0; i-- {
		str += string(s[i])
	}
	return str
}

func RemoveNonAlphanumeric(s string) string {
	str := ""
	for i := range s {
		if (48 <= s[i] && s[i] <= 57) || (97 <= s[i] && s[i] <= 122) {
			str += string(s[i])
		}
	}
	return str
}
