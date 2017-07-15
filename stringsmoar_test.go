package stringsmoar

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRuneFrequency(t *testing.T) {
	testCases := []runeFrequencyTestObject{
		{s: "a", expected: map[rune]int{'a': 1}},
		{s: "a猫猫", expected: map[rune]int{'a': 1, '猫': 2}},
		{s: "猫猫猫bcc", expected: map[rune]int{'猫': 3, 'b': 1, 'c': 2}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%#v has distinct rune counts", tc.s), func(t *testing.T) {
			assertMapsRuneIntEqual(t, RuneFrequency(tc.s), tc.expected)
		})
	}
}

type runeFrequencyTestObject struct {
	s        string
	expected map[rune]int
}

func TestSet(t *testing.T) {
	testCases := []stringStringTestObject{
		{s: "a", expected: "a"},
		{s: "a猫猫", expected: "a猫"},
		{s: "猫猫猫bcc", expected: "猫bc"},
		{s: "7Z猫猫猫Zcc猫ZZ", expected: "7Z猫c"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%#v can be reduced to distinct rune counts", tc.s), func(t *testing.T) {
			result := Set(tc.s)
			if tc.expected != result {
				t.Error("\nExpected:", tc.expected, "\nReceived: ", result)
			}
		})
	}
}

func TestSorted(t *testing.T) {
	testCases := []stringStringTestObject{
		{s: "a", expected: "a"},
		{s: "zaZA", expected: "AZaz"},
		{s: "a猫猫", expected: "a猫猫"},
		{s: "猫猫猫b9cC0", expected: "09Cbc猫猫猫"},
		{s: "猫b猫-c猫C7猫", expected: "-7Cbc猫猫猫猫"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%#v runes can be sorted", tc.s), func(t *testing.T) {
			result := Sorted(tc.s)
			if tc.expected != result {
				t.Error("\nExpected:", tc.expected, "\nReceived: ", result)
			}
		})
	}
}

func TestRemoveNthRune(t *testing.T) {
	testCases := []removeNthRuneTestObject{
		{s: "abc", i: 1, expected: "ac"},
		{s: "", i: 0, expected: ""},
		{s: "", i: 1, expected: ""},
		{s: "a", i: 0, expected: ""},
		{s: "ab", i: 0, expected: "b"},
		{s: "abc", i: 2, expected: "ab"},
		{s: "abc", i: 3, expected: "abc"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%#v removal of location %#v", tc.s, tc.i), func(t *testing.T) {
			result := RemoveNthRune(tc.s, tc.i)
			if tc.expected != result {
				t.Error("\nExpected:", tc.expected, "\nReceived: ", result)
			}
		})
	}
}

type removeNthRuneTestObject struct {
	s        string
	i        int
	expected string
}

func TestReplaceNthRuneHappyPath(t *testing.T) {
	testCases := []replaceNthRuneTestObject{
		{s: "a", i: 0, r: 'a', expected: "a"},
		{s: "a", i: 0, r: 'Z', expected: "Z"},
		{s: "abc", i: 0, r: 'b', expected: "bbc"},
		{s: "abc", i: 1, r: 'b', expected: "abc"},
		{s: "abc", i: 2, r: 'b', expected: "abb"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%#v replacement of location %#v with %v", tc.s, tc.i, tc.r), func(t *testing.T) {
			result, _ := replaceNthRune(tc.s, tc.i, tc.r)
			if tc.expected != result {
				t.Error("\nExpected:", tc.expected, "\nReceived: ", result)
			}
		})
	}
}

func TestReplaceNthRuneErrors(t *testing.T) {
	testCases := []replaceNthRuneTestObject{
		{s: "a", i: 0, r: -1, expected: "", expectedError: "Invalid Rune"},
		{s: "a", i: -1, r: 'b', expected: "a", expectedError: ""},
		{s: "a", i: 2, r: 'b', expected: "a", expectedError: ""},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%#v replacement of location %#v with %v", tc.s, tc.i, tc.r), func(t *testing.T) {
			result, err := replaceNthRune(tc.s, tc.i, tc.r)
			if tc.expected != result {
				t.Error("\nExpected:", tc.expected, "\nReceived: ", result)
			}
			if err != nil && tc.expectedError != err.Error() {
				t.Error("\nExpected Error:", tc.expectedError, "\nReceived: ", err)
			}
		})
	}
}

type replaceNthRuneTestObject struct {
	s             string
	i             int
	r             rune
	expected      string
	expectedError string
}

func TestPermutations(t *testing.T) {
	testCases := []stringStringSliceTestObject{
		{s: "a", expected: []string{"a"}},
		{s: "猫", expected: []string{"猫"}},
		{s: "猫b", expected: []string{"猫b", "b猫"}},
		{s: "ba", expected: []string{"ba", "ab"}},
		{s: "abc", expected: []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
		{s: "abcd", expected: []string{
			"abcd", "abdc", "acbd", "acdb", "adbc", "adcb",
			"bacd", "badc", "bcad", "bcda", "bdac", "bdca",
			"cabd", "cadb", "cbad", "cbda", "cdab", "cdba",
			"dabc", "dacb", "dbac", "dbca", "dcab", "dcba"}},
		{s: "猫咪", expected: []string{"猫咪", "咪猫"}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%#v permutations", tc.s), func(t *testing.T) {
			assertSlicesEqual(t, tc.expected, Permutations(tc.s))
		})
	}
}

func TestPermutePick(t *testing.T) {
	testCases := []stringIntStringSliceTestObject{
		{s: "a", n: 1, expected: []string{"a"}},
		{s: "ab", n: 2, expected: []string{"ab", "ba"}},
		{s: "ba", n: 2, expected: []string{"ba", "ab"}},
		{s: "abc", n: 3, expected: []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
		{s: "abc", n: 1, expected: []string{"a", "b", "c"}},
		{s: "abcd", n: 1, expected: []string{"a", "b", "c", "d"}},
		{s: "abcd", n: 2, expected: []string{
			"ab", "ac", "ad",
			"ba", "bc", "bd",
			"ca", "cb", "cd",
			"da", "db", "dc"}},
		{s: "Туч", n: 2, expected: []string{
			"Ту", "Тч", "уТ", "уч", "чТ", "чу"}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%#v permutations", tc.s), func(t *testing.T) {
			assertSlicesEqual(t, tc.expected, PermutePick(tc.s, tc.n))
		})
	}
}

func TestCombinations(t *testing.T) {
	testCases := []stringIntStringSliceTestObject{
		{s: "a", n: 1, expected: []string{"a"}},
		{s: "9猫", n: 2, expected: []string{"9猫"}},
		{s: "ab", n: 1, expected: []string{"a", "b"}},
		{s: "猫bc", n: 1, expected: []string{"猫", "b", "c"}},
		{s: "abc", n: 2, expected: []string{"ab", "ac", "bc"}},
		{s: "ab猫", n: 2, expected: []string{"ab", "a猫", "b猫"}},
		{s: "ab猫9", n: 2, expected: []string{"ab", "a猫", "9a", "b猫", "9b", "9猫"}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%#v combinations", tc.s), func(t *testing.T) {
			assertSlicesEqual(t, tc.expected, Combinations(tc.s, tc.n))
		})
	}
}

func TestDeduplicateRuneCombinations(t *testing.T) {
	testCases := []stringSliceStringSliceTestObject{
		{s: []string{"a"}, expected: []string{"a"}},
		{s: []string{"a", "a"}, expected: []string{"a"}},
		{s: []string{"ab", "ba"}, expected: []string{"ab"}},
		{s: []string{"猫b", "b猫", "猫"}, expected: []string{"b猫", "猫"}},
		{s: []string{"猫b", "b猫", "b猫7", "b7猫", "猫", "7b猫"}, expected: []string{"b猫", "7b猫", "猫"}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%#v deduplicating", tc.s), func(t *testing.T) {
			assertSlicesEqual(t, tc.expected, DeduplicateRuneCombinations(tc.s))
		})
	}
}

type stringStringTestObject struct {
	s        string
	expected string
}

type stringStringSliceTestObject struct {
	s        string
	expected []string
}

type stringSliceStringSliceTestObject struct {
	s        []string
	expected []string
}

type stringIntStringSliceTestObject struct {
	s        string
	n        int
	expected []string
}

func assertMapsRuneIntEqual(t *testing.T, expected map[rune]int, result map[rune]int) {
	if !reflect.DeepEqual(expected, result) {
		t.Error("\nExpected:", expected, "\nReceived: ", result)
	}
}

func assertSlicesEqual(t *testing.T, expected []string, result []string) {
	if !reflect.DeepEqual(expected, result) {
		t.Error("\nExpected:", expected, "\nReceived: ", result)
	}
}
