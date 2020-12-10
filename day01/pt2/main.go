package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)
func solve(nums []int, max, count, sum, i int ) int {
	if count == max {
		if sum == 2020 {
			return sum
		}
		return -1
	
	}

	if i > len(nums) {
		return -1
	}

	for j := i +1; j < len(nums); j++ {
		if solve(nums, max, count + 1, sum + nums[i], j) == 2020 {
			return 1	
		}
	}
	return -1
}



func main () {
	fmt.Println("hello world")

	file, err := os.Open("input.txt")

	defer file.Close()
	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	var text [] int

	for scanner.Scan() {
		t := scanner.Text()
		if i, err := strconv.Atoi(t); err == nil {
			text = append(text, i)
		} else {
			log.Fatalf("failed to parse number")
		}
	}	

	i := solve(text, 3, 0, 0, 0)
	
	fmt.Println(i)

	
}