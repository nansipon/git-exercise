package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ROMANS = map[string]int{
	"M": 1000, "D": 500, "C": 100,
	"L": 50, "X": 10, "V": 5, "I": 1}

func sum(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func multy(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func findArg(line string) (string, error) {
	switch {
	case strings.Contains(line, "+"):
		return "+", nil
	case strings.Contains(line, "-"):
		return "-", nil
	case strings.Contains(line, "*"):
		return "*", nil
	case strings.Contains(line, "/"):
		return "/", nil
	default:
		return "", fmt.Errorf("can't find operator")
	}
}

func calculation(a, b int, op string) (num int, err error) {
	switch op {
	case "+":
		num = sum(a, b)
	case "-":
		num = sub(a, b)
	case "*":
		num = multy(a, b)
	case "/":
		num = div(a, b)
	default:
		err = fmt.Errorf("%s not found", op)
	}

	return
}

func isRoman(num string) bool {
	if _, err := ROMANS[strings.Split(num, "")[0]]; err {
		return true
	}

	return false
}

func romanToInt(num string) int {

	sum := 0
	n := len(num)

	for i := 0; i < n; i++ {
		if i != n-1 && ROMANS[string(num[i])] < ROMANS[string(num[i+1])] {
			sum += ROMANS[string(num[i+1])] - ROMANS[string(num[i])]
			i++
			continue
		}

		sum += ROMANS[string(num[i])]
	}

	return sum
}

func intToRoman(num int) string {
	var roman string = ""
	var numbers = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	var romans = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	var index = len(romans) - 1

	for num > 0 {
		for numbers[index] <= num {
			roman += romans[index]
			num -= numbers[index]
		}
		index -= 1
	}

	return roman
}

func getNumsAndType(line string, op string) (a, b int, rom bool, err error) {
	nums := strings.Split(line, op)

	if len(nums) > 2 {
		return a, b, rom, fmt.Errorf("many operators")
	}

	firstRomType := isRoman(nums[0])
	secondRomType := isRoman(nums[1])

	if firstRomType != secondRomType {
		return a, b, rom, fmt.Errorf("different format")
	}

	if firstRomType && secondRomType {
		rom = true
		a = romanToInt(nums[0])
		b = romanToInt(nums[1])
	} else {
		a, err = strconv.Atoi(nums[0])
		if err != nil {
			return
		}

		b, err = strconv.Atoi(nums[1])
		if err != nil {
			return
		}
	}

	if a < 1 || a > 10 || b < 0 || b > 10 {
		return a, b, rom, fmt.Errorf("%d or %d less 0 or more 10", a, b)
	}

	return a, b, rom, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Для выхода введите !exit\nВведите пример: ")
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		line = strings.ReplaceAll(line, " ", "")

		if line == "!exit" {
			fmt.Println("exiting..")
			return
		}

		operator, err := findArg(line)
		if err != nil {
			panic(err)
		}

		a, b, isRom, err := getNumsAndType(line, operator)
		if err != nil {
			panic(err)
		}

		result, err := calculation(a, b, operator)
		if err != nil {
			panic(err)
		}

		if isRom {
			if result <= 0 {
				panic("roman numbers can't less 0")
			}

			first := intToRoman(a)
			second := intToRoman(b)
			res := intToRoman(result)

			fmt.Println(first, operator, second, "=", res)
		} else {
			fmt.Println(a, operator, b, "=", result)
		}
	}
}
