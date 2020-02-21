package main

// Shuffle shuffles a slice of ints, returning a slice containing the values (now shuffled) of the original slice
type Shuffle func([]int) []int

// GenerateSlice creates an int slice of length n and populates the indices in order from 1..n
func GenerateSlice(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = i + 1
	}
	return slice
}
