package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the input file
	file, err := os.Open("day2.input")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	table := make([][]int, 1000)
	var safe = 0
	// Process each line
	var j = 0
	for scanner.Scan() {
		line := scanner.Text()
		// Check for errors during scanning
		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading file: %v\n", err)
		}
		// Split the line into parts based on commas
		parts := strings.Split(line, " ")
		// Convert each part to an integer
		var row []int

		for _, part := range parts {
			num, err := strconv.Atoi(strings.TrimSpace(part))
			if err != nil {
				fmt.Printf("Error converting %q to integer: %v\n", part, err)
				continue
			}
			row = append(row, num)
		}

		table[j] = row
		j = j + 1
	}
	println("len:", len(table[999]))

	j = 0
	lastnumber := 0
	ascending := true
	i := 0
	for i < 1000 {
		ascending = true
		println(i)
		j = 0
		for j < len(table[i]) {
			number := table[i][j]
			println(i, number)

			if j == 0 {
				lastnumber = 0
			}
			if j == 1 {
				lastnumber = table[i][j-1]

				if number < lastnumber {
					ascending = false
					if lastnumber-number > 3 {
						j = j + 1
						break
					}
				} else if number-lastnumber > 3 {
					j = j + 1
					break
				} else if number == lastnumber {
					j = j + 1
					break
				}
			}
			if j >= 2 {
				lastnumber = table[i][j-1]

				if number < lastnumber {
					if ascending == true {
						j = j + 1
						break
					}
					if lastnumber-number > 3 {
						j = j + 1
						break
					}
				}

				if number > lastnumber {
					if ascending == false {
						j = j + 1
						break
					}
					if number-lastnumber > 3 {
						j = j + 1
						break
					}
				}

				if number == lastnumber {
					j = j + 1
					break
				}

				if j == len(table[i])-1 {
					fmt.Println(" i,j,number,lastnumber", i, j, lastnumber, number, ascending)
				}
				if j == len(table[i])-1 {
					safe = safe + 1
					fmt.Println("safe++")
				}
			}
			j = j + 1
		}
		i = i + 1
	}
	println("safe: ", safe)
}
