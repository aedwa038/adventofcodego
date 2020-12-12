package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

type instruction struct {
	oper string
	arg  int
}

func parse(text []string) []instruction {

	instructions := make([]instruction, 0)
	for _, line := range text {
		if line == "" {
			continue
		}
		s := strings.Split(line, " ")
		oper := s[0]
		arg, _ := strconv.Atoi(s[1])
		instructions = append(instructions, instruction{
			oper: oper,
			arg:  arg,
		})

	}

	return instructions
}

func run(instructions []instruction) int {
	cache := make(map[int]bool)
	acc := 0
	pc := 0
	for {
		ins := instructions[pc]
		switch ins.oper {
		case "nop":

			if ok, _ := cache[pc]; !ok {
				cache[pc] = true
			} else {
				return acc
			}

			pc++
		case "jmp":
			if ok, _ := cache[pc]; !ok {
				cache[pc] = true
			} else {
				return acc
			}
			pc += ins.arg
		case "acc":
			if ok, _ := cache[pc]; !ok {
				cache[pc] = true
			} else {
				return acc
			}
			acc += ins.arg
			pc++
		default:
			fmt.Println("HIT")
			return acc
		}
	}

}

func debug(i []instruction, pc int, cache map[int]bool, acc int) bool {
	if pc >= len(i) {
		fmt.Println(acc)
		return true
	}
	if ok, _ := cache[pc]; ok {
		return false
	}
	ins := i[pc]
	if ins.oper == "acc" {
		if win := debug(i, pc+1, cache, acc+ins.arg); !win {
			cache[pc] = win
			return false
		}

	} else if ins.oper == "nop" {
		if win := debug(i, pc+1, cache, acc); !win {
			cache[pc] = win
			if win = debug(i, pc+ins.arg, cache, acc); !win {
				cache[pc] = win
				return false
			}
		}
	} else if ins.oper == "jmp" {
		cache[pc] = true
		if win := debug(i, pc+ins.arg, cache, acc); !win {
			if win = debug(i, pc+1, cache, acc); !win {
				cache[pc] = win
				return false
			}
		}
	}

	cache[pc] = true

	return true

}
func main() {
	flag.Parse()
	text := readFile(*file)
	instructions := parse(text)
	acc := run(instructions)
	cache := make(map[int]bool)
	fmt.Println(acc)
	fmt.Println(debug(instructions, 0, cache, 0))

}
