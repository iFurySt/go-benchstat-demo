package util

func Includes[T comparable](list []T, ele T) bool {
	for _, v := range list {
		if v == ele {
			return true
		}
	}
	return false
}
