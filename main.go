package main

import (
	"fmt"
	"strconv"
)

type codes struct {
	answer     string
	elementIdx int
}

var m = map[string]codes{
	"000": {"No errors, correct code:", 0},
	"001": {"Error in 4 bit, r3:", 0},
	"010": {"Error in 2 bit, r2:", 0},
	"011": {"Error in 6 bit, i3:", 5},
	"100": {"Error in 1 bit, r1:", 0},
	"101": {"Error in 5 bit, i2:", 4},
	"110": {"Error in 3 bit, i1:", 2},
	"111": {"Error in 7 bit, i4:", 6},
}

func adder(x1 int, x2 int, x3 int, x4 int, symbols string) int {
	result := 0
	printInfo := "S = "
	for i := 0; i < 7; i++ {
		s, _ := strconv.Atoi(string(symbols[i]))

		if i == x1 || i == x2 || i == x3 || i == x4 {
			result += s
			printInfo += strconv.Itoa(s) + " + "
		}
	}
	fmt.Println(printInfo[:len(printInfo)-3] + " = " + strconv.Itoa(result%2))
	return result % 2
}

func main() {
	var symbols string
	fmt.Printf("Введите набор символов: ")
	_, err := fmt.Scanf("%s", &symbols)
	if err != nil || len(symbols) != 7 {
		fmt.Println("Ошибка ввода!\nВведите валидное значение...")
		return
	}
	s1 := adder(0, 2, 4, 6, symbols)
	s2 := adder(1, 2, 5, 6, symbols)
	s3 := adder(3, 4, 5, 6, symbols)
	fmt.Println("S1:", s1, "S2:", s2, "S3:", s3)
	element := m[strconv.Itoa(s1)+strconv.Itoa(s2)+strconv.Itoa(s3)]
	if element.elementIdx != 0 {
		s, _ := strconv.Atoi(string(symbols[element.elementIdx]))
		if element.elementIdx != 6 {
			symbols = symbols[:element.elementIdx] + strconv.Itoa(1-s) + symbols[element.elementIdx+1:]
		} else {
			symbols = symbols[:element.elementIdx] + strconv.Itoa(1-s)
		}
	}
	fmt.Println("Result:", element.answer, symbols)
}
