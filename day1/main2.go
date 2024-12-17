package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Open the input file
	file, err := os.Open("day1.input")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	var one []int
	var two []int
	var sum = 0
	// Process each line
	for scanner.Scan() {
		line := scanner.Text()
		// Split the line into parts based on commas
		parts := strings.Split(line, "   ")

		// Convert each part to an integer
		for i, part := range parts {
			num, err := strconv.Atoi(strings.TrimSpace(part))
			if err != nil {
				fmt.Printf("Error converting %q to integer: %v\n", part, err)
				continue
			}
			if i == 0 {
				one = append(one, num)
			} else {
				two = append(two, num)

			}
		}

	}
	// Print the tuple

	sort.Ints(one)
	sort.Ints(two)
	fmt.Println("one")
	fmt.Println(one)
	fmt.Println("two")
	fmt.Println(one)
	var counter = 0
	for _, i := range one {
		counter = 0
		for _, j := range two {
			if i == j {
				counter = counter + 1
			}
		}
		fmt.Println("counter: ", counter)
		sum = sum + (counter * i)
		fmt.Println("Sum: ", sum)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

}
