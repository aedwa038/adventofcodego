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

func parse ( text []string ) {

	for i := 0; i < len(text);i++ {
		fmt.Println(text[i])	
	} 
}



func main () {
	flag.Parse()
	text := readFile(*file)
	parse(text)

}