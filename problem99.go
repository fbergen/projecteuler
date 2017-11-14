package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

//  a^x > b^y = log(a^x) > log(b^y) = x log(a) > y log(b)
func main() {
	f, err := ioutil.ReadFile("p099_base_exp.txt")
	if err != nil {
		panic(err)
	}

	base_exps := strings.Split(string(f), "\n")
	max_z := float64(0)
	biggest_line := 0
	for linenr, base_exp := range base_exps {
		pair := strings.Split(base_exp, ",")
		base_int, _ := strconv.Atoi(pair[0])
		exp_int, _ := strconv.Atoi(pair[1])
		z := math.Log(float64(base_int)) * float64(exp_int)
		if z > max_z {
			max_z = z
			biggest_line = linenr + 1 // Line starts from 0
		}
	}

	fmt.Println(biggest_line)
}
