package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const max_red = 12
const max_green = 13
const max_blue = 14

func find_color_amount(round_contents string, color string) int {
	red_regex, _ := regexp.Compile("(\\d+) " + color)
	submatch := red_regex.FindStringSubmatch(round_contents)
	if len(submatch) > 0 {
		found_amount, _ := strconv.Atoi(submatch[1])
		return found_amount
	}
	return 0
}

func round_is_valid(round_contents string) bool {
	if find_color_amount(round_contents, "red") > max_red {
		return false
	}
	if find_color_amount(round_contents, "green") > max_green {
		return false
	}
	if find_color_amount(round_contents, "blue") > max_blue {
		return false
	}
	return true
}

func game_is_valid(game_contents string) bool {
	rounds := strings.Split(game_contents, ";")
	for _, round := range rounds {
		if !round_is_valid(round) {
			return false
		}
	}
	return true
}

func get_min_number_of_cubes_in_game(game_contents string) [4]int {
	red_amount := 0
	green_amount := 0
	blue_amount := 0

	rounds := strings.Split(game_contents, ";")
	for _, round := range rounds {
		round_red_amount := find_color_amount(round, "red")
		if round_red_amount > red_amount {
			red_amount = round_red_amount
		}

		round_green_amount := find_color_amount(round, "green")
		if round_green_amount > green_amount {
			green_amount = round_green_amount
		}

		round_blue_amount := find_color_amount(round, "blue")
		if round_blue_amount > blue_amount {
			blue_amount = round_blue_amount
		}
	}

	return [4]int{red_amount, green_amount, blue_amount, int(math.Max(float64(red_amount), 1) * math.Max(float64(green_amount), 1) * math.Max(float64(blue_amount), 1))}
}

func main() {
	file, err := os.Open("day02/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	valid_game_id_sum := 0
	sum_of_power_of_sets := 0

	for scanner.Scan() {
		line := scanner.Text()
		game_id, _ := strconv.Atoi(line[5:strings.Index(line, ":")])
		game_contents := line[strings.Index(line, ":")+1:]
		if game_is_valid(game_contents) {
			valid_game_id_sum += game_id
		}

		game_minimum_cubes := get_min_number_of_cubes_in_game(game_contents)
		sum_of_power_of_sets += game_minimum_cubes[3]
	}

	fmt.Println("Valid game ID sum:", valid_game_id_sum)
	fmt.Println("Sum of the power of the sets:", sum_of_power_of_sets)

	if err = scanner.Err(); err != nil {
		panic(err)
	}
}
