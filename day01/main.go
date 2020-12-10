package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")

	defer file.Close()
	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	var text []int

	for scanner.Scan() {
		t := scanner.Text()
		if i, err := strconv.Atoi(t); err == nil {
			text = append(text, i)
		} else {
			log.Fatalf("failed to parse number")
		}
	}

	m := make(map[int]int)
	for _, num := range text {

		t := 2020 - num
		_, ok := m[t]
		if ok == true {
			fmt.Printf("Numbers found %v %v \n", num, t)
			fmt.Printf("Result: %v\n", num*t)
		}
		m[num] = t
	}

}
