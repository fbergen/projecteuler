package main

import "fmt"

// Get all prime factors of a given number n
func primeFactors(n int) (pfs []int) {
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}

// return p^i
func power(p, i int) int {
	result := 1
	for j := 0; j < i; j++ {
		result *= p
	}
	return result
}

// formula comes from https://math.stackexchange.com/a/22723
func sumOfProperDivisors(n int) int {
	if n == 0 {
		return 0
	}
	pfs := primeFactors(n)

	// key: prime
	// value: prime exponents
	m := make(map[int]int)
	for _, prime := range pfs {
		_, ok := m[prime]
		if ok {
			m[prime] += 1
		} else {
			m[prime] = 1
		}
	}

	sumOfAllFactors := 1
	for prime, exponents := range m {
		sumOfAllFactors *= (power(prime, exponents+1) - 1) / (prime - 1)
	}
	return sumOfAllFactors - n
}

func next(i int) int {
	s := sumOfProperDivisors(i)
	return s
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	longest_chain := 0
	smallest_member := 1000000
	gseen := make(map[int]bool)
	for i := 2; i < 1000000; i++ {
		if i%10000 == 0 {
			fmt.Println("i ", i)
		}
		if gseen[i] {
			continue
		}
		gseen[i] = true
		seen := make(map[int]bool)
		seen[i] = true
		curr := next(i)
		for ; gseen[curr] == false && curr < 1000000; curr = next(curr) {
			gseen[curr] = true
			seen[curr] = true
		}
		// we found a chain at curr
		if seen[curr] {
			start := curr
			length := 0
			smallest := start
			for curr = next(curr); curr != start; curr = next(curr) {
				length++
				smallest = min(smallest, curr)
			}
			if length > longest_chain {
				longest_chain = length
				smallest_member = smallest
				fmt.Println("smallest_member", smallest_member, " longest_chain", longest_chain, " i ", i)
			} else if length == longest_chain {
				smallest_member = max(smallest, smallest_member)
				fmt.Println("smallest_member", smallest_member, " longest_chain", longest_chain, " i ", i)
			}
		}
	}
	fmt.Println("smallest_member", smallest_member, " longest_chain", longest_chain)
	//fmt.Println("Cache hit!", m[in])
}
