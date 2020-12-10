package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	
)

type PasswordField struct {
	min int
	max int
	required string
	password string	
}


func parse ( text []string ) ([]PasswordField) {
   var fields [] PasswordField
   
   for _, t := range text {
	 mi, _ := strconv.Atoi(t[:strings.IndexByte(t, '-')])
	 ma, _ := strconv.Atoi(t[strings.IndexByte(t, '-') + 1: strings.IndexByte(t, ' ')])
	 r := strings.Trim(t[strings.IndexByte(t, ' '): strings.IndexByte(t, ':')], " ")
	 p := strings.Trim(t[strings.IndexByte(t, ':') + 1:], " ") 
    fields = append(fields, PasswordField{
		min: mi,
		max: ma,
		required: r,
		password: p,
	})	
   }

   return fields
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

	fields := parse(text)
	
	var count int
	for _, field := range fields {
		fmt.Printf("'%v'\n", field)
		c := strings.Count(field.password, field.required)

		if c >= field.min && c <= field.max {
			count++
		}
		

	}

	fmt.Println(count)
	
}


