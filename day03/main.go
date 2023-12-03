package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func get_complete_number(line string, starting_index int) int {
	completeNumber := string(line[starting_index])
	iterIndex := starting_index - 1
	for iterIndex >= 0 {
		if unicode.IsDigit(rune(line[iterIndex])) {
			completeNumber = string(line[iterIndex]) + completeNumber
			iterIndex--
		} else {
			break
		}
	}
	iterIndex = starting_index + 1
	for iterIndex < len(line) {
		if unicode.IsDigit(rune(line[iterIndex])) {
			completeNumber += string(line[iterIndex])
			iterIndex++
		} else {
			break
		}
	}

	converted, _ := strconv.Atoi(completeNumber)
	return converted
}

func check_row_for_number(line string, number_range []int) []int {
	beginIndex := Max(0, number_range[0]-1)
	endIndex := Min(len(line), number_range[1]+1)

	stringToCheck := line[beginIndex:endIndex]
	exp, _ := regexp.Compile("\\d+")

	allString := exp.FindAllStringIndex(stringToCheck, -1)
	var allNumbers []int
	for _, regexFoundNumberRange := range allString {
		// adjust range by up to 3
		frontAdjustment := Max(0, beginIndex-3)
		endAdjustment := Min(len(line), endIndex+3)

		allNumbers = append(allNumbers, get_complete_number(line[frontAdjustment:endAdjustment], regexFoundNumberRange[0]+(beginIndex-frontAdjustment)))
	}

	return allNumbers
}

func check_all_numbers_in_line(lines []string, row_index int, max_index int) int {
	exp, _ := regexp.Compile("\\*")

	line := lines[row_index]
	numberRanges := exp.FindAllStringIndex(line, -1)

	sumFromLine := 0
	for _, numberRange := range numberRanges {
		var gearNumbers []int
		gearNumbers = append(gearNumbers, check_row_for_number(line, numberRange)[:]...)
		if row_index > 0 {
			gearNumbers = append(gearNumbers, check_row_for_number(lines[row_index-1], numberRange)[:]...)
		}
		if row_index < max_index {
			gearNumbers = append(gearNumbers, check_row_for_number(lines[row_index+1], numberRange)[:]...)
		}
		if len(gearNumbers) == 2 {
			fmt.Println("Found a gear", gearNumbers)
			gearRatio := gearNumbers[0] * gearNumbers[1]
			sumFromLine += gearRatio
		}
	}
	return sumFromLine
}

func main() {
	file, err := os.Open("day03/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var fileAsList []string

	for scanner.Scan() {
		line := scanner.Text()
		fileAsList = append(fileAsList, line)
	}

	sumOfFile := 0
	for i := range fileAsList {
		sumOfFile += check_all_numbers_in_line(fileAsList, i, len(fileAsList)-1)
	}
	fmt.Println("Sum of valid numbers in file:", sumOfFile)

	if err = scanner.Err(); err != nil {
		panic(err)
	}
}
