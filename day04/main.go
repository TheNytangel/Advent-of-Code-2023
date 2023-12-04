package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func contains(list []int, number int) bool {
	for _, b := range list {
		if b == number {
			return true
		}
	}
	return false
}

func calculate_score(winning_numbers []int, players_numbers []int) int {
	number_of_winning_numbers := 0
	for _, number := range players_numbers {
		if contains(winning_numbers, number) {
			number_of_winning_numbers += 1
		}
	}
	if number_of_winning_numbers == 0 {
		return 0
	}
	return int(math.Pow(2, float64(number_of_winning_numbers-1)))
}

func parse_numbers(section string) []int {
	var allNumbers []int
	split := strings.Split(section, " ")
	for _, s := range split {
		number, err := strconv.Atoi(strings.Trim(s, " "))
		if err == nil {
			allNumbers = append(allNumbers, number)
		}
	}
	return allNumbers
}

func main() {
	file, err := os.Open("day04/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total_score := 0
	for scanner.Scan() {
		line := scanner.Text()

		winning_numbers := parse_numbers(line[10:40])
		players_numbers := parse_numbers(line[42:])
		total_score += calculate_score(winning_numbers, players_numbers)
	}

	fmt.Println("Total score:", total_score)

	if err = scanner.Err(); err != nil {
		panic(err)
	}
}
