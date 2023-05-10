package utils

import "strings"

func SplitAndRemoveSpaces(s string) []string {
	split := strings.Split(s, ",")
	ret := make([]string, 0)
	for _, s := range split {
		n := strings.TrimSpace(s)
		if n != "" {
			ret = append(ret, n)
		}
	}
	return ret
}
