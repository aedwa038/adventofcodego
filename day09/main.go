package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var file = flag.String("file", "input.txt", "Input file for processing")
var size = flag.Int("p", 25, "preamble size")

func readFile(filename string) []string {

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatalf("failed to open")
	}
	scanner := bufio.NewScanner(file)
	var text []string
	for scanner.Scan() {
		t := scanner.Text()
		text = append(text, t)
	}

	return text
}

func parse(text []string) []int {
	nums := make([]int, 0)
	for i := 0; i < len(text); i++ {
		n, _ := strconv.Atoi(text[i])
		nums = append(nums, n)
	}

	return nums
}

func search(n int, nums []int) bool {

	cache := make(map[int]bool)
	for _, num := range nums {
		target := n - num
		if _, ok := cache[target]; ok {
			return true
		}
		cache[num] = true
	}

	return false
}
func solve(nums []int, p int) (int, []int) {

	for i := p; i < len(nums); i++ {
		n := nums[i]
		l := nums[i-p : i]
		if !search(n, l) {
			return i, nums[:i]
		}
	}
	return -1, []int{}
}

func solve2(nums []int, sum int) []int {
	start := 0
	end := 1
	rsum := nums[start]

	for start < len(nums) {

		if sum == rsum {
			return nums[start:end]
		}

		if rsum < sum {
			if len(nums) >= end {
				rsum += nums[end]
				end++
			}
		}

		for rsum > sum {
			rsum -= nums[start]
			start++
		}
	}

	return []int{}
}

func main() {
	flag.Parse()
	text := readFile(*file)
	nums := parse(text)
	i, n := solve(nums, *size)
	fmt.Println("Part1: ", nums[i])
	fmt.Println(n)
	l := solve2(n, nums[i])
	fmt.Println(l)
	sort.Ints(l)
	fmt.Println(l)
	fmt.Println("Part2: ", l[0]+l[len(l)-1])

}
