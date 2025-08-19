package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	r := make(map[string]int)
	for k, v := range in {
		for _, c := range v {
			r[strings.ToLower(c)] = k
		}
	}

	return r
}
