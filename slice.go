package goutils

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Merge(s1, s2 []string) []string {
	return append(s1, s2...)
}
