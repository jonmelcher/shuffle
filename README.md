# shuffle

Using the language of your choice implement the following program named Shuffle:
- Input a single integer N.
- Output the integers 1..N in random order.
- Your algorithm should optimize for speed.
- Feel free to ask us any questions.

## implementation
- chosen language is Go
- chosen algorithm is the Fisher-Yates Shuffle algorithm
    - uniform distribution (shuffled elements are equally likely to appear in any index)
    - O(1) space complexity
    - O(n) time complexity as we must pass over the slice of integers at least once, and computing a random number is O(1)
    - random number generator will be the global math/rand generator, seeded with the time
        - if security were a concern then we could reconsider how we seed, and consider using the global crypto/rand random number generator
- program is a binary which can be executed with a single integer argument `N`
    - `N` cannot be negative, or a value unable to be parsed to an integer
    - `0` is an acceptable value for `N`, no integers will be outputted

## testing
- run `go test` in the root project directory
- the distribution tests take some time and are running for both native / FisherYates algorithms, can specify test files or test cases with as arguments to `go test`

## building / installing
- run `go build` in the root project directory to generate the `shuffle` binary in the root directory
- alternatively, set `GOBIN` to point to the project's bin directory, and run `go install` to generate the `shuffle` binary there
- a `shuffle` binary has been committed for convenience

## running the program
- execute the `shuffle` binary with a single integer argument `N`, eg. `./bin/shuffle 5` from the root project directory

## output
- each integer will be outputted to stdout on an individual line
- in the case where there are no integers, no output will occur
