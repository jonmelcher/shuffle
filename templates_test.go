package main

import "testing"

func ShuffleWithEmptySlice(t *testing.T, shuffle Shuffle) {
	arr := []int{}
	arr = shuffle(arr)
	if len(arr) != 0 {
		t.Errorf("Shuffle was incorrect, started with an empty slice and ended up with more values: %v", arr)
	}
}

func ShuffleWithSingleElementSlice(t *testing.T, shuffle Shuffle) {
	arr := []int{0}
	arr = shuffle(arr)
	if len(arr) != 1 {
		t.Errorf("Shuffle was incorrect, started with a single element slice and ended up with more values: %v", arr)
	}
	if arr[0] != 0 {
		t.Errorf("Shuffle was incorrect, started with a 0 and ended up with a different value: %d", arr[0])
	}
}

// UniformDistributionTestConfig serves as a container with named properties for more clarity when calling VerifyShuffleIsUniformlyDistributed
type UniformDistributionTestConfig struct {
	testState  *testing.T
	shuffle    Shuffle
	runs       int
	elements   int
	errorDelta int
}

// VerifyShuffleIsUniformlyDistributed runs a shuffle algorithm multiple times, aggregating frequencies of each element at each index
// finally verifying that each frequency falls within expected bounds of the mathematically expected frequency for uniform distribution
func VerifyShuffleIsUniformlyDistributed(config *UniformDistributionTestConfig) {
	expectedFrequency := config.runs / config.elements
	expectedUpperBound := expectedFrequency + config.errorDelta
	expectedLowerBound := expectedFrequency - config.errorDelta

	distributions := make([]map[int]int, config.elements)
	arr := make([]int, config.elements)
	for i := 0; i < config.elements; i++ {
		distributions[i] = make(map[int]int)
		arr[i] = i
	}

	// shuffle and record frequency of each value appearing at each index
	for i := 0; i < config.runs; i++ {
		tmp := arr
		tmp = config.shuffle(tmp)
		for j := 0; j < len(tmp); j++ {
			value := tmp[j]
			distributions[value][j] = distributions[value][j] + 1
		}
	}

	// for each value, verify its frequency at any index in the array is within expected bounds
	for value := 0; value < len(distributions); value++ {
		for index, frequency := range distributions[value] {
			if frequency < expectedLowerBound || frequency > expectedUpperBound {
				config.testState.Logf("Frequency %d of %d at index %d fell outside of expected bounds [%d, %d]", frequency, value, index, expectedLowerBound, expectedUpperBound)
				config.testState.Fail()
			}
		}
	}
}
