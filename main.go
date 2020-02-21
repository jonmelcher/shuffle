package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 2 {
		panic(fmt.Sprintf("N not passed in when executing program."))
	}

	input := os.Args[1]
	N, err := strconv.ParseInt(input, 10, 0)

	if err != nil {
		panic(fmt.Sprintf("N: %s is unable to be parsed into an integer.", input))
	}

	if N < 0 {
		panic(fmt.Sprintf("N: %s is not a positive integer.", input))
	}

	output := GenerateSlice(int(N))

	for _, val := range output {
		fmt.Println(val)
	}
}
