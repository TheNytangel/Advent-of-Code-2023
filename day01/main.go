package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var numbers = [...]string{
	"0",
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func GetFirstDigit(line string) int32 {
	earliest_idx_in_str := len(line) + 1
	number_idx_found := 0
	for i := 0; i < len(numbers); i++ {
		str_idx := strings.Index(line, numbers[i])
		if str_idx != -1 && str_idx < earliest_idx_in_str {
			earliest_idx_in_str = str_idx
			number_idx_found = i
		}
	}
	return int32(number_idx_found % 10)
}

func GetLastDigit(line string) int32 {
	latest_idx_in_str := -1
	number_idx_found := 0
	for i := 0; i < len(numbers); i++ {
		str_idx := strings.LastIndex(line, numbers[i])
		if str_idx != -1 && str_idx > latest_idx_in_str {
			latest_idx_in_str = str_idx
			number_idx_found = i
		}
	}
	return int32(number_idx_found % 10)
}

func main() {
	file, err := os.Open("day01/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	calibration_sum := int32(0)
	for scanner.Scan() {
		line := scanner.Text()

		calibration_value := GetFirstDigit(line)*10 + GetLastDigit(line)
		calibration_sum += calibration_value
	}

	fmt.Printf("Sum of calibration values: %d", calibration_sum)

	if err = scanner.Err(); err != nil {
		panic(err)
	}
}
