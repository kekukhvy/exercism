package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

func Keep[T any](in []T, filterFunc func(T) bool) []T {
	var out []T

	for _, v := range in {
		if filterFunc(v) {
			out = append(out, v)
		}
	}
	return out
}

func Discard[T any](in []T, filterFunc func(T) bool) []T {
	var out []T

	for _, v := range in {
		if !filterFunc(v) {
			out = append(out, v)
		}
	}
	return out
}
