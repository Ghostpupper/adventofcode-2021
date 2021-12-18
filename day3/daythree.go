package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"github.com/Ghostpupper/adventofcode-2021/util"
)

func main() {
	
	partOne(util.ReadInputString("input.txt"))
}

func partOne(textlines []string) {
	string_len := len(textlines[0])
	var bits = make([]int, string_len)
	var gamma_rate = make([]int, string_len)
	var epsilon_rate = make([]int, string_len)
	var gamma_rate_float, epsilon_rate_float float64
	entries := len(textlines)
	for _, line := range textlines {
		int_list := convert_to_int_list(line)
		for position, value := range int_list {
			bits[position] += value
		}
	}
	for position, bit_entry := range bits {
		if bit_entry >= entries/2 {
			gamma_rate_float += math.Pow(2, float64(string_len) - 1 - float64(position))
			gamma_rate[position] = 1
			epsilon_rate[position] = 0
		} else {
			gamma_rate[position] = 0
			epsilon_rate_float += math.Pow(2, float64(string_len) - 1 - float64(position))
			epsilon_rate[position] = 1
		}
	}
	fmt.Print("Res:", int(gamma_rate_float*epsilon_rate_float))

}

func convert_to_int_list(line string) []int {
	var return_list []int
	string_list := strings.Split(line, "")
	for _, digit := range string_list {
		thing, _ := strconv.Atoi(digit)
		return_list = append(return_list, thing)
	}
	return return_list
}
