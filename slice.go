package goutils

//type Comparable interface {
//
//}

func Contains[T comparable](s []T, e T) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Merge[T any](s1, s2 []T) []T {
	return append(s1, s2...)
}
