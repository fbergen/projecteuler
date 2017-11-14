package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	var sum int64 = 0
	count := 0
	//for i := int64(3); i < 333333333; i += 2 {
	for i := int64(3); i < 93687; i += 2 {
		//for i := int64(235000835); i < 235000836; i += 2 {
		if i%10000000 == 0 {
		}
		if isAlmostEquilateral(i, i-1) {
			count++
			sum += i*3 - 1
		}
		if isAlmostEquilateral(i, i+1) {
			count++
			sum += i*3 + 1
		}
	}

	fmt.Println(sum, count)
}

func isAlmostEquilateral(a, b int64) bool {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}

	if diff > 1 {
		fmt.Println("HEHEHEHEHE", diff)
		return false
	}

	x := big.NewFloat(0.0)
	biga := big.NewFloat(float64(a))
	biga.Mul(biga, biga)
	bigb := big.NewFloat(float64(b))
	//fmt.Printf("%6.4f\n", bigb)
	bigb.Mul(bigb, bigb)
	//fmt.Printf("%6.4f\n", bigb)
	x.Sub(biga.Mul(biga, big.NewFloat(4)), bigb)
	x.Mul(bigb, x)
	if !x.IsInt() {
		//fmt.Println("HERE1")
		return false
	}

	e, _ := x.Int(nil)
	m := big.NewInt(0)
	e.DivMod(e, big.NewInt(16), m)
	if m.Cmp(big.NewInt(0)) != 0 {
		//fmt.Println("HERE2", m)
		return false
	}

	//if !q.IsInt64() {
	//	fmt.Printf("HERE22")
	//	return false
	//}

	f := float64(e.UInt64())
	//fmt.Println(e)
	//fmt.Println(e.Int64())
	//fmt.Println(f)

	area := math.Sqrt(f)
	if area*area == f {
		fmt.Println("HERE3", a, b)
		return true
	}
	//fmt.Println("HERE4")
	return false
	//
	//	//fmt.Printf("%6.4f\n", bigb)
	//	bf, _ := bigb.Float64()
	//	//fmt.Printf("%6.4f\n", bf)
	//	h := math.Sqrt(bf)
	//	//fmt.Printf("%6.10f\n", h)
	//	//fmt.Printf("%6.10f\n", h*float64(b))
	//
	//	//h := math.Sqrt(math.Pow(af, 2) - math.Pow(bf/2, 2))
	//	area := float64(b) * h / 2
	//	if area == math.Floor(area) {
	//		//fmt.Printf("%d, %6.10f, %6.10f, %6.10f\n", a, h, area, math.Floor(area))
	//		return true
	//	}
	//	return false
}
