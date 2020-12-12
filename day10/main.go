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

func solve(nums []int, index int) int {
	fmt.Println("solve", nums, len(nums))
	return solve2(nums, index)
}

func solve2(nums []int, index int) int {
	//fmt.Println(index)
	if index >= len(nums)-1 {
		return 1
	}
	count := 0
	for i := index + 1; i < len(nums); i++ {
		diff := nums[i] - nums[index]
		//fmt.Println(diff)
		if diff <= 3 {
			count += solve2(nums, i)
		} else {
			break
		}
	}

	return count

}
func solve3(nums []int, index int, cache map[int]int) int {
	//fmt.Println(index)
	if _, ok := cache[index]; ok {
		return cache[index]
	}

	if index >= len(nums)-1 {
		cache[index] = 1
		return cache[index]
	}

	count := 0
	for i := index + 1; i < len(nums) && 3 >= (nums[i]-nums[index]); i++ {

		count += solve3(nums, i, cache)

	}
	cache[index] = count
	return cache[index]

}
func main() {
	flag.Parse()
	text := readFile(*file)
	nums := parse(text)

	sort.Ints(nums)
	jolt1Count := 0
	jolt3Count := 0
	diff := nums[0] - 0
	if diff == 1 {
		jolt1Count++
	} else if diff == 3 {
		jolt3Count++
	}

	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff == 1 {
			jolt1Count++
		}
		if diff >= 3 {
			jolt3Count++
		}
	}
	jolt3Count++

	fmt.Println(jolt3Count, jolt1Count)
	fmt.Println(jolt1Count * jolt3Count)

	fmt.Println(nums)
	nums = append(nums, nums[len(nums)-1]+3)
	nums = append(nums, 0)
	sort.Ints(nums)
	//fmt.Println(solve(nums, 0))
	fmt.Println(solve3(nums, 0, map[int]int{}))

}
