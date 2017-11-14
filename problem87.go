package main

import (
	"fmt"
)

func main() {
	primePowerTriples := make(map[int]bool)
	primes := getPrimesTo(7071)
	fmt.Println("primes", len(primes))
	for _, i := range primes {
		for _, j := range primes {
			if j <= 368 {
				for _, k := range primes {
					if k <= 85 {
						num := i*i + j*j*j + k*k*k*k
						if num <= 50000000 {
							primePowerTriples[num] = true
						}
					}
				}
			}
		}
	}
	fmt.Println(len(primePowerTriples))
}

func getPrimesTo(x int) []int {
	primes := make([]int, 0)
	primes = append(primes, 2)
	for i := 2; i < 7072; i++ {
		is_prime := true
		for _, prime := range primes {
			if i%prime == 0 {
				is_prime = false
				break
			}
		}
		if is_prime {
			primes = append(primes, i)
		}
	}
	return primes
}
