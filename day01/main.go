package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func GetFirstDigit(line string) int32 {
	for _, character := range line {
		if unicode.IsDigit(character) {
			return character - '0'
		}
	}
	return 0
}

func GetLastDigit(line string) int32 {
	for i := len(line) - 1; i >= 0; i-- {
		rune_value := rune(line[i])
		if unicode.IsDigit(rune_value) {
			return rune_value - '0'
		}
	}
	return 0
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
