package stringsmoar

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRunes(t *testing.T) {
	testCases := []stringRuneSliceTestObject{
		{s: "a", expected: []rune{'a'}},
		{s: "猫b", expected: []rune{'猫', 'b'}},
		{s: "猫bчч", expected: []rune{'猫', 'b', 'ч', 'ч'}},
		{s: "b 猫-7", expected: []rune{'b', ' ', '猫', '-', '7'}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%#v into runes", tc.s), func(t *testing.T) {
			assertRuneSlicesEqual(t, Runes(tc.s), tc.expected)
		})
	}
}

func TestRuneFrequency(t *testing.T) {
	testCases := []stringMapRuneIntTestObject{
		{s: "", expected: map[rune]int{}},
		{s: "a", expected: map[rune]int{'a': 1}},
		{s: "a猫猫", expected: map[rune]int{'a': 1, '猫': 2}},
		{s: "猫猫猫bccч", expected: map[rune]int{'猫': 3, 'b': 1, 'c': 2, 'ч': 1}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%#v has distinct rune counts", tc.s), func(t *testing.T) {
			assertMapsRuneIntEqual(t, RuneFrequency(tc.s), tc.expected)
		})
	}
}

func TestSet(t *testing.T) {
	testCases := []stringStringTestObject{
		{s: "", expected: ""},
		{s: "a", expected: "a"},
		{s: "a猫猫", expected: "a猫"},
		{s: "猫猫猫bcc", expected: "猫bc"},
		{s: "7Z猫猫猫Zcc猫ZZ", expected: "7Z猫c"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%#v can be reduced to distinct rune counts", tc.s), func(t *testing.T) {
			assertStringsEqual(t, tc.expected, Set(tc.s))
		})
	}
}

func TestExclusive(t *testing.T) {
	testCases := []stringMapRuneBoolStringTestObject{
		{s: "a", runes: map[rune]bool{'a': true}, expected: "a"},
		{s: "a", runes: map[rune]bool{}, expected: ""},
		{s: "猫a", runes: map[rune]bool{'猫': true}, expected: "猫"},
		{s: "猫a", runes: map[rune]bool{'猫': true}, expected: "猫"},
		{s: "чч9猫ччч", runes: map[rune]bool{'ч': true, '9': true}, expected: "чч9ччч"},
		{s: "zyxabc9猫", runes: map[rune]bool{}, expected: ""},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%#v with only runes %#v", tc.s, tc.runes), func(t *testing.T) {
			assertStringsEqual(t, tc.expected, Exclusive(tc.s, tc.runes))
		})
	}
}

func TestRemoveWhenAdjacentRunes(t *testing.T) {
	testCases := []stringStringTestObject{
		{s: "a", expected: "a"},
		{s: "aab", expected: "b"},
		{s: "猫a", expected: "猫a"},
		{s: "foo9bar猫猫", expected: "f9bar"},
		{s: "чччfoo9bar猫猫", expected: "f9bar"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%#v with any adjacent-repeated runes removed", tc.s), func(t *testing.T) {
			assertStringsEqual(t, tc.expected, removeWhenAdjacentRunes(tc.s))
		})
	}
}

// TODO: TestGetAdjacentRunes

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
			assertStringsEqual(t, tc.expected, Sorted(tc.s))
		})
	}
}

func TestRemoveNthRune(t *testing.T) {
	var testCases = []struct {
		s        string
		i        int
		expected string
	}{
		{s: "", i: 0, expected: ""},
		{s: "", i: 1, expected: ""},
		{s: "a", i: 0, expected: ""},
		{s: "abc", i: 3, expected: "abc"},
		{s: "abc", i: 1, expected: "ac"},
		{s: "猫b", i: 0, expected: "b"},       // runeValue, width := utf8.DecodeRuneInString(s[i:]) , 3 bytes
		{s: "aчc", i: 3, expected: "aч"},     // a = 1 byte, ч = 2 bytes
		{s: "猫猫猫9ч", i: 6, expected: "猫猫9ч"}, // 猫 is 3 bytes
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%#v removal of location %#v", tc.s, tc.i), func(t *testing.T) {
			assertStringsEqual(t, tc.expected, RemoveNthRune(tc.s, tc.i))
		})
	}
}

func TestRemoveNthItem(t *testing.T) {
	var testCases = []struct {
		a        []string
		target   int
		expected []string
	}{
		{a: []string{"a", "b"}, target: 0, expected: []string{"b"}},
		{a: []string{"a", "b"}, target: 1, expected: []string{"a"}},
		{a: []string{"猫", "b", "c"}, target: 0, expected: []string{"b", "c"}},
		{a: []string{"a", "b", "猫"}, target: 1, expected: []string{"a", "猫"}},
		// negative cases
		{a: []string{}, target: 0, expected: []string{}},
		{a: []string{"aчc"}, target: 0, expected: []string{}},
		{a: []string{"a"}, target: 2, expected: []string{"a"}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%#v removal of index %#v", tc.a, tc.target), func(t *testing.T) {
			assertSlicesEqual(t, tc.expected, RemoveNthItemSlow(tc.a, tc.target))
			assertSlicesEqual(t, tc.expected, RemoveNthItem(tc.a, tc.target))

		})
	}
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

func TestPermutationsSlices(t *testing.T) {
	var testCases = []struct {
		a        []string
		expected [][]string
	}{
		{a: []string{"a"}, expected: generateSliceOfStringSlices([]string{"a"})},
		{a: []string{"a", "b"},
			expected: generateSliceOfStringSlices([]string{"a", "b"}, []string{"b", "a"})},
		{a: []string{"a", "b", "猫咪"},
			expected: generateSliceOfStringSlices(
				[]string{"a", "b", "猫咪"}, []string{"a", "猫咪", "b"},
				[]string{"b", "a", "猫咪"}, []string{"b", "猫咪", "a"},
				[]string{"猫咪", "a", "b"}, []string{"猫咪", "b", "a"})},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%#v permutation slices", tc.a), func(t *testing.T) {
			result := PermutationsSlices(tc.a)
			if !reflect.DeepEqual(tc.expected, result) {
				t.Error("\nExpected:", tc.expected, "\nReceived: ", result)
			}
		})
	}
}

func generateSliceOfStringSlices(a ...[]string) [][]string {
	return a
}

func TestPermutePick(t *testing.T) {
	testCases := []stringIntStringSliceTestObject{
		{s: "a", n: 1, expected: []string{"a"}},
		{s: "猫", n: 1, expected: []string{"猫"}},
		{s: "ab", n: 2, expected: []string{"ab", "ba"}},
		{s: "ba", n: 2, expected: []string{"ba", "ab"}},
		{s: "ab猫", n: 3, expected: []string{"ab猫", "a猫b", "ba猫", "b猫a", "猫ab", "猫ba"}},
		{s: "猫bc", n: 1, expected: []string{"猫", "b", "c"}},
		{s: "abcd", n: 1, expected: []string{"a", "b", "c", "d"}},
		{s: "abcd", n: 2, expected: []string{
			"ab", "ac", "ad",
			"ba", "bc", "bd",
			"ca", "cb", "cd",
			"da", "db", "dc"}},
		{s: "Туч", n: 2, expected: []string{"Ту", "Тч", "уТ", "уч", "чТ", "чу"}},
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

type stringRuneSliceTestObject struct {
	s        string
	expected []rune
}

type stringMapRuneIntTestObject struct {
	s        string
	expected map[rune]int
}

type stringMapRuneBoolStringTestObject struct {
	s        string
	runes    map[rune]bool
	expected string
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

func assertStringsEqual(t *testing.T, expected string, result string) {
	// t.Helper()
	if expected != result {
		t.Error("\nExpected:", expected, "\nReceived: ", result)
	}
}

func assertRuneSlicesEqual(t *testing.T, expected []rune, result []rune) {
	// t.Helper()
	if !reflect.DeepEqual(expected, result) {
		t.Error("\nExpected:", expected, "\nReceived: ", result)
	}
}

func assertSlicesEqual(t *testing.T, expected []string, result []string) {
	if !reflect.DeepEqual(expected, result) {
		t.Error("\nExpected:", expected, "\nReceived: ", result)
	}
}

func assertMapsRuneIntEqual(t *testing.T, expected map[rune]int, result map[rune]int) {
	if !reflect.DeepEqual(expected, result) {
		t.Error("\nExpected:", expected, "\nReceived: ", result)
	}
}

func BenchmarkRemoveNthItemSlowShort(b *testing.B) {
	a := []string{"a", "b", "猫"}
	for i := 0; i < b.N; i++ {
		RemoveNthItemSlow(a, 0)
	}
}

func BenchmarkRemoveNthItemSlowMany(b *testing.B) {
	a := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "猫"}
	for i := 0; i < b.N; i++ {
		RemoveNthItemSlow(a, 14)
	}
}

func BenchmarkRemoveNthItemShort(b *testing.B) {
	a := []string{"a", "b", "猫"}
	for i := 0; i < b.N; i++ {
		RemoveNthItem(a, 0)
	}
}

func BenchmarkRemoveNthItemMany(b *testing.B) {
	a := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "猫"}
	for i := 0; i < b.N; i++ {
		RemoveNthItem(a, 14)
	}
}
