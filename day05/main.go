package main
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"flag"
	"sort"
	
)

var file = flag.String("file", "input.txt", "Input file for processing")

func readFile(filename string) ([]string) {
	
    file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatalf("failed to open")
	}
	scanner := bufio.NewScanner(file)
	var text [] string
	for scanner.Scan() {
		t := scanner.Text()
		text = append(text, t)
	}
	
	return text
}

func parse ( text []string ) {

	for i := 0; i < len(text);i++ {
		fmt.Println(text[i])	
	} 
}

func binSearch (line string) (int, int)  {

	lo := 0
	hi := 127

	sl := 0
	sh := 7
	for _, c := range line {
		mid := lo + (hi - lo) / 2
		if c == 'F' {
			hi = mid - 1
		} else if c == 'B' {
			lo = mid + 1
		} 
		mid = sl + (sh - sl) / 2
		if c == 'L' {
			sh = mid - 1
		} else if c == 'R' {
			sl = mid + 1
		} 
		
	}

	return lo, sl
}

func main () {
	flag.Parse()
	text := readFile(*file)
	//parse(text)
	highest := 0
	var list []int
	for _, line := range text {
		r,c:= binSearch(line)
		seatID := (r * 8 ) + c
		list = append(list, seatID)
		if seatID > highest {
			highest = seatID
		}
		fmt.Println(line, c,r, seatID)

	}
	
	fmt.Println(highest)
	sort.Ints(list)
	fmt.Println(list)
	for i := 1; i < len(list) - 1; i++ {
		d := list[i] - list[i - 1]
		dl := list[i + 1] - list[i]
		if d > 1 || dl > 1  {
			fmt.Println(text[i], text[i - 1], text[i + 1], d, dl, list[i], list[i - 1], list[i + 1] )
		}
		
		
	} 
}