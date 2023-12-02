package main

import (
	"bufio"
	"fmt"
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

func main() {
	file, err := os.Open("day02/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	valid_game_id_sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		game_id, _ := strconv.Atoi(line[5:strings.Index(line, ":")])
		game_contents := line[strings.Index(line, ":")+1:]
		if game_is_valid(game_contents) {
			valid_game_id_sum += game_id
		}
	}

	fmt.Println("Valid game ID sum:", valid_game_id_sum)

	if err = scanner.Err(); err != nil {
		panic(err)
	}
}
