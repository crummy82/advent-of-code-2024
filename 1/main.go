package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const INPUTFILE string = "input.txt"

// PART 2
func main() {

	readFile, err := os.Open(INPUTFILE)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanWords)
	var col1 []int
	var col2 []int
	var count int
	for fileScanner.Scan() {
		count++
		if count%2 == 1 {
			num, _ := strconv.Atoi(fileScanner.Text())
			col1 = append(col1, num)
		} else {
			num, _ := strconv.Atoi(fileScanner.Text())
			col2 = append(col2, num)
		}
	}
	readFile.Close()

	var totalSim int
	var matches int

	for i := 0; i < len(col1); i++ {
		for j := range col2 {
			if col2[j] == col1[i] {
				matches++
			}
		}
		similarity := col1[i] * matches
		totalSim += similarity
		matches = 0
	}

	fmt.Printf("total diff %d \n", totalSim)
}

// PART 1

// func main() {

// 	readFile, err := os.Open(INPUTFILE)

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fileScanner := bufio.NewScanner(readFile)
// 	fileScanner.Split(bufio.ScanWords)
// 	var col1 []int
// 	var col2 []int
// 	var count int
// 	for fileScanner.Scan() {
// 		count++
// 		if count%2 == 1 {
// 			num, _ := strconv.Atoi(fileScanner.Text())
// 			col1 = append(col1, num)
// 		} else {
// 			num, _ := strconv.Atoi(fileScanner.Text())
// 			col2 = append(col2, num)
// 		}
// 	}

// 	readFile.Close()

// 	slices.Sort(col1)
// 	slices.Sort(col2)

// 	var totalDiff int

// 	for i := 0; i < len(col1); i++ {
// 		diff := math.Abs(float64(col1[i] - col2[i]))
// 		totalDiff += int(diff)
// 	}

// 	fmt.Printf("total diff %d \n", totalDiff)
// }
