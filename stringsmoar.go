package stringsmoar

import (
	"bytes"
	"errors"
	// "fmt"
	"unicode/utf8"
)

/*
func main() {
	s := "abcd"
	p := permute(s)
	fmt.Println(s, "permutations:", p)
	pPick := permutePick(s, 2)
	fmt.Println(s, "permutations pick 2:", pPick)
}
*/

// permutePick generates the permutations when only subset N of S is picked
func permutePick(s string, n int) []string {
	return permutePickInternal(s, n, 0)
}

func permutePickInternal(s string, n int, current int) []string {
	if len(s) <= 1 {
		return []string{s}
	}
	var result []string
	current++
	for i, v := range s {
		if current == n {
			result = append(result, string(v))
		} else {
			p := permutePickInternal(removeNthRune(s, i), n, current)
			for _, c := range p {
				result = append(result, string(v)+c)
			}
		}
	}
	return result
}

// permute generates all the permutations of the runes in a string
func permute(s string) []string {
	if len(s) <= 1 {
		return []string{s}
	}
	var result []string
	for i, v := range s {
		p := permute(removeNthRune(s, i))
		for _, c := range p {
			result = append(result, string(v)+c)
		}
	}
	return result
}

func removeNthRune(s string, n int) string {
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
