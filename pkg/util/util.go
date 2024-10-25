package util

import "sort"

func Includes(ss []string, s string) bool {
	sort.Strings(ss)
	index := sort.SearchStrings(ss, s)
	if index < len(ss) && ss[index] == s {
		return true
	}
	return false
}
