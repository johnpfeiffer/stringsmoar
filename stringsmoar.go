// Package stringsmoar contains moar utility functions for working with strings.
// Extending the helper functions from https://golang.org/pkg/strings and https://golang.org/src/strings/strings.go
package stringsmoar

import (
	"bytes"
	"errors"
	"sort"
	"unicode/utf8"
)

// RuneFrequency returns a map of the count of each rune in the string
func RuneFrequency(s string) map[rune]int {
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}
	return m
}

// Set returns a string where each rune from the original string only occurs once, in the order that they first appear
func Set(s string) string {
	var uniques string
	m := make(map[rune]struct{})
	for _, r := range s {
		_, ok := m[r]
		if !ok {
			m[r] = struct{}{}
			// TODO: benchmark to compare performance of various string creation techniques
			uniques = uniques + string(r)
		}
	}
	return uniques
}

// Sorted returns a string where each rune from the original is now sorted
func Sorted(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, k int) bool { return runes[i] < runes[k] })
	// TODO: benchmark an alternative of strings.Split() -> sort.Strings -> strings.Join()
	return string(runes)
}

// Permutations generates all the permutations of the runes in a string https://en.wikipedia.org/wiki/Permutation
func Permutations(s string) []string {
	if utf8.RuneCountInString(s) <= 1 {
		return []string{s}
	}
	var result []string
	for i, v := range s {
		p := Permutations(RemoveNthRune(s, i))
		for _, c := range p {
			result = append(result, string(v)+c)
		}
	}
	return result
}

// PermutePick generates the permutations when only subset N of S is picked
func PermutePick(s string, n int) []string {
	return permutePickInternal(s, n, 0)
}

func permutePickInternal(s string, n int, current int) []string {
	if utf8.RuneCountInString(s) <= 1 {
		return []string{s}
	}
	var result []string
	current++
	for i, v := range s {
		if current == n {
			result = append(result, string(v))
		} else {
			p := permutePickInternal(RemoveNthRune(s, i), n, current)
			for _, c := range p {
				result = append(result, string(v)+c)
			}
		}
	}
	return result
}

// Combinations chooses the subset N of S without regards to order , https://en.wikipedia.org/wiki/Combination
func Combinations(s string, n int) []string {
	if utf8.RuneCountInString(s) <= n {
		return []string{s}
	}
	var result []string
	p := permutePickInternal(s, n, 0)
	result = DeduplicateRuneCombinations(p)
	return result
}

// DeduplicateRuneCombinations returns a slice of strings where each one is a unique (sorted) rune combination (aka a set)
func DeduplicateRuneCombinations(strings []string) []string {
	var uniques []string
	m := make(map[string]struct{})
	for _, s := range strings {
		sSorted := Sorted(s)
		_, ok := m[sSorted]
		if !ok {
			m[sSorted] = struct{}{}
			uniques = append(uniques, sSorted)
		}
	}
	return uniques
}

// RemoveNthRune removes a specific rune from the string by it's index location
func RemoveNthRune(s string, n int) string {
	buffer := bytes.NewBuffer(nil)
	for i, r := range s {
		if i != n {
			buffer.WriteRune(r)
		}
	}
	return buffer.String()
	/* // below works but seems like it has an extra allocation
	runes := make([]rune, len(s))
	copy(runes, []rune(s))
	string(append(runes[:i], runes[i+1:]...))
	*/
}

func replaceNthRune(s string, n int, newR rune) (string, error) {
	if !utf8.ValidRune(newR) {
		return "", errors.New("Invalid Rune")
	}
	buffer := bytes.NewBuffer(nil)
	for i, r := range s {
		if i == n {
			buffer.WriteRune(newR)
		} else {
			buffer.WriteRune(r)
		}
	}
	return buffer.String(), nil
}
