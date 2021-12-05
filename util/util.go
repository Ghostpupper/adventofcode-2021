package util

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func readInput(file_name string) []int {
	file, err := os.Open(file_name)
 
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