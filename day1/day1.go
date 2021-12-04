package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)
var buffer [4]int
var counter int

func readInput() []int {
	file, err := os.Open("input.txt")
 
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
 
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var intlines []int
 
	for scanner.Scan() {
		txtline := scanner.Text()
		intline, _ := strconv.Atoi(txtline)
		intlines = append(intlines, intline)
	}
 
	file.Close()
 
	return intlines 
}

func left_shift(new_number int){
	buffer[0] = buffer[1]
	buffer[1] = buffer[2]
	buffer[2] = buffer[3]
	buffer[3] = new_number
}

func increase_comparison(){
	window1 := buffer[0] + buffer[1] + buffer[2]
	window2 := buffer[1] + buffer[2] + buffer[3]
	fmt.Print("window comparison: ", window1, " vs ", window2)
	if window2 > window1{
		fmt.Print(" increasing \n")
		counter++
	} else if window1 == window2{
		fmt.Print(" no change \n")
	} else {
		fmt.Print(" decreasing \n")
	}
}

func main() {
	counter = 0
	inputArray := readInput()
	for i := 0; i < len(buffer); i++ {
		buffer[i] = inputArray[i]
	}
	increase_comparison()
	for _, this_line := range inputArray[4:] {
		left_shift(this_line)
		increase_comparison()
	}
	fmt.Print("Counter is: ", strconv.Itoa(counter))
}