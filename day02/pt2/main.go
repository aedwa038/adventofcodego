package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	
)




func parse ( text []string ) {
   
   for _, t := range text {
	   fmt.Printn(t)
   }
}

func main () {
	fmt.Println("hello world")

	file, err := os.Open("input.txt")

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

	parse(text)
	


	
}


