package main

import "math/rand"

// Shuffle shuffles a slice of ints, returning a slice containing the values (now shuffled) of the original slice
type Shuffle func([]int) []int

// GeneratePopulatedSlice creates an int slice of length n and populates the indices in order from 1..n
func GeneratePopulatedSlice(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = i + 1
	}
	return slice
}

// FisherYatesShuffle shuffles a slice in-place based on the Fisher-Yates shuffle algorithm, returning the original slice
// for reference: https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle
func FisherYatesShuffle(arr []int) []int {
	length := len(arr)
	// with a small enough array there will be no swapping
	// moreover, this protects against cases where we might have been invoking rand.Intn(0)
	// it also guards against nil slices (length 0)
	if length < 2 {
		return arr
	}
	// we stop before the last element as it will have no other element to swap with
	// this also guards against invoking rand.Intn(0)
	// not swapping with previously set elements is what helps achieve a uniform distribution
	for i := 0; i < length-1; i++ {
		swapIdx := i + rand.Intn(length-i)
		arr[i], arr[swapIdx] = arr[swapIdx], arr[i]
	}

	return arr
}
