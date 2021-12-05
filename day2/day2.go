package main

import (
	"fmt"
	"github.com/Ghostpupper/adventofcode-2021/util"
	"strings"
	"strconv"
)
var position_x, position_y int

func main() {
	partOne(util.ReadInputString("input.txt"))
	fmt.Print("Horizontal position: ", position_x,"\nDepth Position: ", position_y)
	fmt.Print("\nMultiplied Together: ", position_x*position_y)
}

func partOne(textlines []string) {
	position_y, position_x = 0, 0
	for _, line := range textlines{
		lastchar := line[len(line)-1:]
		number, _ := strconv.Atoi(lastchar)
		if strings.Contains(line, "forward"){
			move_forward(number)
		} else if strings.Contains(line, "up"){
			move_up(number)
		} else if strings.Contains(line, "down"){
			move_down(number)
		}
	}
}

func move_forward(value int){
	position_x += value
}

func move_up(value int){
	position_y -= value
	if position_y < 0{
		position_y = 0
	}
}

func move_down(value int){
	position_y += value
}