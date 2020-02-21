package main

import (
	"math/rand"
	"testing"
)

// nativeShuffle satisfies the Shuffle function signature
// it is a wrapper over the math/rand native Shuffle function
func nativeShuffle(arr []int) []int {
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	return arr
}

func TestNativeShuffleShufflesAnEmptySlice(t *testing.T) {
	ShuffleWithEmptySlice(t, nativeShuffle)
}

func TestNativeShuffleShufflesASingleElementSlice(t *testing.T) {
	ShuffleWithSingleElementSlice(t, nativeShuffle)
}

// expected frequency - 2000000 (5 possible slots over 10000000 runs)
// boundaries - +- 5000
// test asserts distribution is uniform with a 0.5% error margin
func TestShuffleDistributionWithFiveElementsNative(t *testing.T) {
	VerifyShuffleIsUniformlyDistributed(&UniformDistributionTestConfig{testState: t, shuffle: nativeShuffle, runs: 10000000, elements: 5, errorDelta: 5000})
}

// expected frequency - 1000000 (10 possible slots over 10000000 runs)
// boundaries - += 3000
// test asserts distribution is uniform with a 0.6% error margin
func TestShuffleDistributionWithTenElementsNative(t *testing.T) {
	VerifyShuffleIsUniformlyDistributed(&UniformDistributionTestConfig{testState: t, shuffle: nativeShuffle, runs: 10000000, elements: 10, errorDelta: 3000})
}

// expected frequency - 120000 (50 possible slots over 5000000 runs)
// boundaries - += 1500
// test asserts distribution is uniform with a 2.5% error margin
func TestShuffleDistributionWithFiftyElementsNative(t *testing.T) {
	VerifyShuffleIsUniformlyDistributed(&UniformDistributionTestConfig{testState: t, shuffle: nativeShuffle, runs: 6000000, elements: 50, errorDelta: 1500})
}
