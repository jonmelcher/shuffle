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
