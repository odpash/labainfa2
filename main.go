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
	"000": {"Ошибок нет, правильный код: ", 0},
	"001": {"Ошибка в 4 бите, r3: ", 0},
	"010": {"Ошибка во 2 бите, r2: ", 0},
	"011": {"Ошибка в 6 бите, i3: ", 5},
	"100": {"Ошибка в 1 бите, r1: ", 0},
	"101": {"Ошибка в 5 бите, i2: ", 4},
	"110": {"Ошибка в 3 бите, i1: ", 2},
	"111": {"Ошибка в 7 бите, i4: ", 6},
}

func adder(x1 int, x2 int, x3 int, x4 int, symbols string) int {
	result := 0
	for i := 0; i < 7; i++ {
		s, _ := strconv.Atoi(string(symbols[i]))

		if i == x1 || i == x2 || i == x3 || i == x4 {
			result += s
		}
	}
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
	element := m[strconv.Itoa(s1)+strconv.Itoa(s2)+strconv.Itoa(s3)]
	if element.elementIdx != 0 {
		s, _ := strconv.Atoi(string(symbols[element.elementIdx]))
		if element.elementIdx != 6 {
			symbols = symbols[:element.elementIdx] + strconv.Itoa(1-s) + symbols[element.elementIdx+1:]
		} else {
			symbols = symbols[:element.elementIdx] + strconv.Itoa(1-s)
		}
	}
	fmt.Println(element.answer, symbols)
}
