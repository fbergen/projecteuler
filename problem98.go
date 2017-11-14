package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

func main() {
	f, err := ioutil.ReadFile("p098_words_lines.txt")
	if err != nil {
		panic(err)
	}

	all_words := strings.Split(string(f), "\n")

	anagrams := getAnagrams(all_words)

	//anagrams = map[string][]string{"AECR": []string{"CARE", "RACE"}}
	square_num := 0
	for _, words := range anagrams {
		s := getAnagramSquare(words)
		if s > 0 {
			fmt.Println(s)
		}
		square_num = max(square_num, s)
	}
	fmt.Println(square_num)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getAnagramSquareForMapping(mapping map[rune]int, words []string) int {
	squares_found := 0
	max_sum := 0
	for _, word := range words {
		sum := 0
		for _, c := range word {
			sum = sum*10 + mapping[c]
		}
		if float64(sum) >= math.Pow(float64(10), float64(len(word)-1)) && isSquare(sum) {
			max_sum = max(max_sum, sum)
			squares_found++
		}
	}
	if squares_found >= 2 {
		return max_sum
	}
	return 0
}

func getAnagramSquare(words []string) int {
	key := getWordKey(words[0])
	perms := getAllPermutations(len(key), make([]int, 0))

	global_max := 0
	for _, perm := range perms {
		mapping := make(map[rune]int)
		for i, c := range key {
			mapping[c] = perm[i]
		}
		global_max = max(global_max, getAnagramSquareForMapping(mapping, words))

	}
	return global_max
}

func getAllPermutations(length int, perm []int) [][]int {
	if length <= 0 {
		return append(make([][]int, 0), perm)
	}

	ret := make([][]int, 0)
	for j := 0; j < 10; j++ {
		if !contains(perm, j) {
			perm_copy := append(perm, j)
			for _, child_perm := range getAllPermutations(length-1, perm_copy) {
				ret = append(ret, child_perm)
			}
		}
	}
	return ret
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func isSquare(num int) bool {
	sqrt := int(math.Sqrt(float64(num)))
	return sqrt*sqrt == num

}

func getWordKey(word string) string {
	s := strings.Split(word, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func getAnagrams(words []string) map[string][]string {
	var anagrams map[string][]string = make(map[string][]string, 0)
	for _, word := range words {
		key := getWordKey(word)
		anagrams[key] = append(anagrams[key], word)
	}

	// Delete single words
	for key, value := range anagrams {
		if len(value) == 1 {
			delete(anagrams, key)
		}
	}
	fmt.Println(anagrams)

	return anagrams
}
