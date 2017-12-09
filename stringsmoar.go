// Package stringsmoar contains moar utility functions for working with strings.
// Extending the helper functions from https://golang.org/pkg/strings and https://golang.org/src/strings/strings.go
package stringsmoar

import (
	"bytes"
	"errors"
	"sort"
	"strings"
	"unicode/utf8"
)

// Runes returns a slice of runes from a string
func Runes(s string) []rune {
	var runes []rune
	for _, r := range s {
		runes = append(runes, r)
	}
	return runes
}

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

// Exclusive returns a string which contains only the runes that are in the map
func Exclusive(s string, runes map[rune]bool) string {
	var result string
	for _, v := range s {
		if runes[v] {
			result += string(v)
		}
	}
	return result
}

// removeWhenAdjacentRunes will remove runes that repeat, i.e. aaabccd will become bd
func removeWhenAdjacentRunes(s string) string {
	if utf8.RuneCountInString(s) < 2 {
		return s
	}
	runes := Runes(s)
	duplicates := getAdjacentRunes(runes)
	reduced := s
	for _, r := range duplicates {
		reduced = strings.Replace(reduced, string(r), "", -1)
	}
	return reduced
}

func getAdjacentRunes(runes []rune) []rune {
	var duplicates []rune
	for i := 1; i < len(runes); i++ {
		if runes[i-1] == runes[i] {
			duplicates = append(duplicates, runes[i])
		}
	}
	return duplicates
}

// RemoveNthRune removes a specific rune from the string by it's index location (i.e. the value returned by range s)
func RemoveNthRune(s string, n int) string {
	if s == "" {
		return s
	}
	buffer := bytes.NewBuffer(nil)
	for i, r := range s {
		if i != n {
			buffer.WriteRune(r)
		}
	}
	return buffer.String()
}

// RemoveNthItem returns a completey new copy of the slice with a specific item (by index location) removed
func RemoveNthItem(a []string, target int) []string {
	result := []string{}
	for i := 0; i < len(a); i++ {
		if i != target {
			result = append(result, a[i])
		}
	}
	return result
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
			smaller := RemoveNthRune(s, i)
			p := permutePickInternal(smaller, n, current)
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
