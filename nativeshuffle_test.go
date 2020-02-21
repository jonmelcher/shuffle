// this file was originally used to run and test the unit tests
package main

import (
	"math/rand"
	"testing"
)

// NativeShuffle satisfies the Shuffle function signature
// it is a wrapper over the math/rand native Shuffle function
func NativeShuffle(arr []int) []int {
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	return arr
}

func TestNativeShuffleShufflesAnEmptySlice(t *testing.T) {
	ShuffleWithEmptySlice(t, NativeShuffle)
}

func TestNativeShuffleShufflesASingleElementSlice(t *testing.T) {
	ShuffleWithSingleElementSlice(t, NativeShuffle)
}

// expected frequency - 2000000 (5 possible slots over 10000000 runs)
// boundaries - +- 5000
// test asserts distribution is uniform with a 0.5% error margin
func TestShuffleDistributionWithFiveElementsNative(t *testing.T) {
	VerifyShuffleIsUniformlyDistributed(&UniformDistributionTestConfig{testState: t, shuffle: NativeShuffle, runs: 10000000, elements: 5, errorDelta: 5000})
}

// expected frequency - 1000000 (10 possible slots over 10000000 runs)
// boundaries - += 4000
// test asserts distribution is uniform with a 0.8% error margin
func TestShuffleDistributionWithTenElementsNative(t *testing.T) {
	VerifyShuffleIsUniformlyDistributed(&UniformDistributionTestConfig{testState: t, shuffle: NativeShuffle, runs: 10000000, elements: 10, errorDelta: 4000})
}

// expected frequency - 120000 (50 possible slots over 5000000 runs)
// boundaries - += 1500
// test asserts distribution is uniform with a 2.5% error margin
func TestShuffleDistributionWithFiftyElementsNative(t *testing.T) {
	VerifyShuffleIsUniformlyDistributed(&UniformDistributionTestConfig{testState: t, shuffle: NativeShuffle, runs: 6000000, elements: 50, errorDelta: 1500})
}
