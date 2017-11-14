package main

import (
	"fmt"
	"math/big"
)

func main() {
	s := big.NewInt(0)
	s.Exp(big.NewInt(2), big.NewInt(7830457), nil).
		Mul(s, big.NewInt(28433)).
		Add(s, big.NewInt(1)).
		Mod(s, big.NewInt(10000000000))
	fmt.Println(s)
}
