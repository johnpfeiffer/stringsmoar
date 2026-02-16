# Description

**stringsmoar** is a lot of small helper functions for strings that are not in the Go standard library (or even in x packages). <https://pkg.go.dev/std>

I fully understand the need to keep the Go standard library lean and maintainable, but after using Python it can be frustrating having to constantly rewrite simple things.

I hope this open source code is helpful (at the very least as an example of how to approach some of the common problems).

updated to do ContinuousIntegration builds with:

- <https://app.circleci.com/pipelines/github/johnpfeiffer/stringsmoar> (see the current build status)
- <https://circleci.com/developer/images/image/cimg/go>
- <https://circleci.com/docs/2.0/collect-test-data/>


*<https://blog.john-pfeiffer.com/golang-testing-benchmark-profiling-subtests-fuzz-testing/> is a longer blog post on Golang testing*

# Features

- `Runes` - string to rune slice
- `RuneFrequency` - count of each rune in a string
- `Set` - deduplicate runes, preserving first-occurrence order
- `Exclusive` - keep only runes present in a given map
- `ExcludeRunesWithAdjacentDuplicates` - remove runes entirely when they repeat adjacently
- `FindAdjacentDuplicateRunes` - find which runes have adjacent duplicates
- `Sorted` - sort runes in a string
- `RemoveNthRune` - remove a rune by its byte-index position
- `RemoveNthItem` / `RemoveNthItemSlow` - remove an item from a string slice by index
- `ConsecutiveIndex` - find the end of a consecutive rune run
- `Permutations` - all rune permutations of a string
- `PermutationsSlices` - all permutations of a string slice
- `PermutePick` - permutations when picking a subset of N runes
- `Combinations` - order-independent subsets of N runes
- `DeduplicateRuneCombinations` - deduplicate sorted rune combinations

# Tests

While 100% coverage is a dubious achievement, having some mterics and data can still be helpful:

`go test ./...`

`go test -cover ./...`

`go test -covermode=count -coverprofile=count.out .`

`go tool cover -html=count.out`

> cool heat map , <https://go.dev/blog/cover>

