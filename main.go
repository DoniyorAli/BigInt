package main

import (
	"myBigint/work"
	"fmt"
)

func main() {

	n1, err := work.NewInt("1795")
	if err != nil {
		panic(err)
	}
	n2, err := work.NewInt("678")
	if err != nil {
		panic(err)
	}
	addNums := work.Add(n1, n2)
	fmt.Println(addNums)
//*===========================================
	n3, err := work.NewInt("1587")
	if err != nil {
		panic(err)
	}
	n4, err := work.NewInt("587")
	if err != nil {
		panic(err)
	}
	subNums := work.Sub(n3, n4)
	fmt.Println(subNums)
//*===========================================
	n5, err := work.NewInt("11")
	if err != nil {
		panic(err)
	}
	n6, err := work.NewInt("11")
	if err != nil {
		panic(err)
	}
	multiNums := work.Multiples(n5, n6)
	fmt.Println(multiNums)
//*===========================================
	n7, err := work.NewInt("1587")
	if err != nil {
		panic(err)
	}
	n8, err := work.NewInt("587")
	if err != nil {
		panic(err)
	}
	divNums := work.Division(n7, n8)
	fmt.Println(divNums)
//*===========================================
	n9, err := work.NewInt("1587")
	if err != nil {
		panic(err)
	}
	n10, err := work.NewInt("587")
	if err != nil {
		panic(err)
	}
	modNums := work.Mod(n9, n10)
	fmt.Println(modNums)
//*===========================================
	n11 := work.Bigint{Value: "-123"}

	absNums := work.Abs(n11)
	fmt.Println(absNums)
}