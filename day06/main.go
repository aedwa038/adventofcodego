package main
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"flag"
	
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

func solve1(groups[]string) int {
	cMap := make(map[rune]int) 
	for i := 0; i < len(groups);i++ {
		s := groups[i]
		for _, c := range s {
			sum, ok := cMap[c]
			if ok {
				cMap[c] = sum + 1
			} else {
				cMap[c] = 1
			}
		}
	}
	return len(cMap)
}
func solve2 (groups[]string)(int) {
	cMap := make(map[rune]int)
	e := len(groups) 
	for i := 0; i < len(groups);i++ {
		s := groups[i]
		for _, c := range s {
			sum, ok := cMap[c]
			if ok {
				cMap[c] = sum + 1
			} else {
				cMap[c] = 1
			}
		}
	}
	count := 0
	for _, value := range cMap {
		if( value == e) {
			count++
		}
	}
	fmt.Println(count)
	return count

}

func parse ( text []string ) {
	count := 0
	count2 := 0
	g := make([]string, 0)
	for i := 0; i < len(text);i++ {
		s := text[i]
		if( s == "") {
			fmt.Println(g)
			count += solve1(g)
			count2 += solve2(g)
			g = make([]string, 0)
			continue
		} else {
			g = append(g, s)
		}
		
	}
	fmt.Println("Solution one", count)
	fmt.Println("solution two", count2)
}



func main () {
	flag.Parse()
	text := readFile(*file)
	parse(text)

}