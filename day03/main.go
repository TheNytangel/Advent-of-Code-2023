package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

func check_row_for_character(line string, number_range []int) bool {
	beginIndex := Max(0, number_range[0]-1)
	endIndex := Min(len(line), number_range[1]+1)

	stringToCheck := line[beginIndex:endIndex]
	match, _ := regexp.MatchString("[^.\\d]", stringToCheck)
	return match
}

func check_all_numbers_in_line(lines []string, row_index int, max_index int) int {
	exp, _ := regexp.Compile("\\d+")

	line := lines[row_index]
	numberRanges := exp.FindAllStringIndex(line, -1)

	sumFromLine := 0
	for _, numberRange := range numberRanges {
		fmt.Println(row_index, line[numberRange[0]:numberRange[1]])
		if check_row_for_character(line, numberRange) || (row_index > 0 && check_row_for_character(lines[row_index-1], numberRange) || (row_index < max_index && check_row_for_character(lines[row_index+1], numberRange))) {
			parsedNumber, _ := strconv.Atoi(line[numberRange[0]:numberRange[1]])
			sumFromLine += parsedNumber
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
