package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func init() {
	// we are assuming that seeding with time is safe enough for the purposes of this program
	// there are risks with doing this, where it is important that someone cannot determine the seed
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {

	if len(os.Args) < 2 {
		panic(fmt.Sprintf("N was not passed in when executing program."))
	}

	input := os.Args[1]
	N, err := strconv.ParseInt(input, 10, 0)

	if err != nil {
		panic(fmt.Sprintf("N: %s is unable to be parsed into an integer.", input))
	}

	if N < 0 {
		panic(fmt.Sprintf("N: %s is not a positive integer.", input))
	}

	output := GeneratePopulatedSlice(int(N))
	output = FisherYatesShuffle(output)
	fmt.Println(output)
}
