package work

import (
	"errors"
	"strconv"
	"strings"
)

type Bigint struct {
	Value string
}

var (
	ErrorNumber = errors.New("this is not number")
)

// * VALIDATION NUMBER
func validation(num string) error {
	allowed := "0123456789"
	start := 0
	if string(num[0]) == "-" || string(num[0]) == "+" {
		start = 1
	}
	str := strings.Split(num, "")
	for _, v := range str[start:] {
		if !strings.Contains(allowed, v) {
			return ErrorNumber
		}
	}
	return nil
}

//* CLEAN NUMBER

func cleanUp(num string) string {
	for strings.HasPrefix(num, "0") {
		num = num[1:]
	}
	return num
}

func NewInt(num string) (Bigint, error) {
	err := validation(num)
	if err != nil {
		return Bigint{Value: ""}, err
	}

	num = cleanUp(num)
	return Bigint{
		Value: num,
	}, nil

}

// // * SET METHODS UPDATES VALUE
// func (x *Bigint) Set(num string) error {
// 	//
// 	//TODO
// 	//
// 	return ErrorNumber
// }

// * ADD NUMBERS
func Add(num1, num2 Bigint) Bigint {
	resp := Bigint{}
	first := strings.Split(num1.Value, "")
	second := strings.Split(num2.Value, "")
	isRest := false

	for i, j := len(first)-1, len(second)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		n1 := 0
		n2 := 0

		if i >= 0 {
			n1, _ = strconv.Atoi(first[i])
		}
		if j >= 0 {
			n2, _ = strconv.Atoi(second[j])
		}

		if isRest {
			n1 = n1 + 1
			isRest = false
		}

		if n1+n2 >= 10 { //* 17, 14
			resp.Value = strconv.Itoa(n1+n2-10) + resp.Value
			isRest = true
		} else {
			resp.Value = strconv.Itoa(n1+n2) + resp.Value
			isRest = false
		}
	}
	if isRest {
		resp.Value = "1" + resp.Value
	}
	return resp
}

// * ==========================================================================================
// * SUB NUMBERS
func Sub(num1, num2 Bigint) Bigint {
	resp := Bigint{}
	first := strings.Split(num1.Value, "")
	second := strings.Split(num2.Value, "")
	isNeed := false

	for i, j := len(first)-1, len(second)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		n1 := 0
		n2 := 0

		if i >= 0 {
			n1, _ = strconv.Atoi(first[i])
		}
		if j >= 0 {
			n2, _ = strconv.Atoi(second[j])
		}

		if isNeed {
			n1 = n1 - 1
			isNeed = false
		}

		if n1-n2 < 0 { //*
			resp.Value = strconv.Itoa(n1-n2+10) + resp.Value
			isNeed = true
		} else {
			resp.Value = strconv.Itoa(n1-n2) + resp.Value
			isNeed = false
		}
	}
	if isNeed {
		resp.Value = "1" + resp.Value
	}
	return resp
}
// * ==========================================================================================
// * MULTIPLE NUMBERS
func Multiples(num1, num2 Bigint) Bigint {
	bag := Bigint{}
	first, _ := strconv.Atoi(num1.Value)
	second, _ := strconv.Atoi(num2.Value)
	bag.Value = strconv.Itoa(first * second)
	return bag
}
// * ==========================================================================================
// * DIVISION NUMBERS
func Division(num1, num2 Bigint) Bigint {
	bag := Bigint{}
	first, _ := strconv.Atoi(num1.Value)
	second, _ := strconv.Atoi(num2.Value)
	bag.Value = strconv.Itoa(first / second)
	return bag
}
// * ==========================================================================================
// * MOD NUMBERS
func Mod(num1, num2 Bigint) Bigint {
	bag := Bigint{}
	first, _ := strconv.Atoi(num1.Value)
	second, _ := strconv.Atoi(num2.Value)
	bag.Value = strconv.Itoa(first % second)
	return bag
}
// * ==========================================================================================
// * Abs NUMBERS
func Abs(num Bigint) Bigint {
	g := Bigint{}
	bag, _ := strconv.Atoi(num.Value[1:])
	g.Value = strconv.Itoa(bag)
	return g
}