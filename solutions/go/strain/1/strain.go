package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

func Keep[T any](in []T, cond func(T) bool) []T {
	var keepArr []T

	for i := range in {
		if cond(in[i]) {
			keepArr = append(keepArr, in[i])
		}
	}
	return keepArr
}

func Discard[T any](in []T, filterFunc func(T) bool) []T {
	var diskArr []T

	for i := range in {
		if !filterFunc(in[i]) {
			diskArr = append(diskArr, in[i])
		}
	}
	return diskArr
}
