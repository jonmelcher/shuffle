package main

import (
	"testing"
)

func TestFisherYatesShuffleWithAnEmptySlice(t *testing.T) {
	ShuffleWithEmptySlice(t, FisherYatesShuffle)
}

func TestFisherYatesShuffleWithASingleElementSlice(t *testing.T) {
	ShuffleWithSingleElementSlice(t, FisherYatesShuffle)
}

// expected frequency - 2000000 (5 possible slots over 10000000 runs)
// boundaries - +- 5000
// test asserts distribution is uniform with a 0.5% error margin
func TestFisherYatesShuffleDistributionWithFiveElementsNative(t *testing.T) {
	VerifyShuffleIsUniformlyDistributed(&UniformDistributionTestConfig{testState: t, shuffle: FisherYatesShuffle, runs: 10000000, elements: 5, errorDelta: 5000})
}

// expected frequency - 1000000 (10 possible slots over 10000000 runs)
// boundaries - += 3000
// test asserts distribution is uniform with a 0.6% error margin
func TestFisherYatesShuffleDistributionWithTenElementsNative(t *testing.T) {
	VerifyShuffleIsUniformlyDistributed(&UniformDistributionTestConfig{testState: t, shuffle: FisherYatesShuffle, runs: 10000000, elements: 10, errorDelta: 3000})
}

// expected frequency - 120000 (50 possible slots over 5000000 runs)
// boundaries - += 1500
// test asserts distribution is uniform with a 2.5% error margin
func TestFisherYatesShuffleDistributionWithFiftyElementsNative(t *testing.T) {
	VerifyShuffleIsUniformlyDistributed(&UniformDistributionTestConfig{testState: t, shuffle: FisherYatesShuffle, runs: 6000000, elements: 50, errorDelta: 1500})
}
