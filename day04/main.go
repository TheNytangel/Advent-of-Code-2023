package main

import (
	"bufio"
	"fmt"
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

func calculate_count_of_winning_numbers(winning_numbers []int, players_numbers []int) int {
	number_of_winning_numbers := 0
	for _, number := range players_numbers {
		if contains(winning_numbers, number) {
			number_of_winning_numbers++
		}
	}
	return number_of_winning_numbers
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

	copies_of_cards := [220]int{}
	for i := range copies_of_cards {
		copies_of_cards[i] = 1
	}
	current_line := 0
	for scanner.Scan() {
		line := scanner.Text()

		winning_numbers := parse_numbers(line[10:40])
		players_numbers := parse_numbers(line[42:])
		count_of_winning_numbers := calculate_count_of_winning_numbers(winning_numbers, players_numbers)

		iter_index := 1
		for iter_index <= count_of_winning_numbers {
			copies_of_cards[current_line+iter_index] += copies_of_cards[current_line]
			iter_index++
		}

		current_line++
	}

	sum_of_scratch_cards := 0
	for i := range copies_of_cards {
		sum_of_scratch_cards += copies_of_cards[i]
	}
	fmt.Println(copies_of_cards)
	fmt.Println("Total cards:", sum_of_scratch_cards)

	if err = scanner.Err(); err != nil {
		panic(err)
	}
}
