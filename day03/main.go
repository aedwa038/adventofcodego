package main
import (
	"bufio"
	"fmt"
	"log"
	"os"
	
)
func parse ( text []string ) ([][]rune) {
	x := len(text)
	y := len(text[0])
    board := make([][] rune, x)
	for i := range  board {
		board[i] = make([]rune, y)
	}
	for i := 0; i < x; i++ {
		t := text[i]
		for j := 0; j < y; j++ {
			board[i][j] = rune(t[j])
		}
	}
	return board
}
func printGrid(grid [][] rune) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%v ", string(grid[i][j]))
		}
		fmt.Println()
	}
}

// CountTrees counts the trees based on slope
func CountTrees(board[][] rune, x, y int) (int) {
   
   width := len(board[0])
   height := len(board)
   i, j := 0, 0
   count := 0
   for  {		
	
	i,j = i + y, j + x
	if j >= width {
		j = j % width
	}
	if i >= height {
		break
	} 
	 if board[i][j] == '#' {
		 count++
	 } 
   }
   return count
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
	board := parse(text)
   printGrid(board)
   fmt.Println(CountTrees(board, 3, 1))
   fmt.Println(CountTrees(board, 1, 1))
   fmt.Println(CountTrees(board, 5, 1))
   fmt.Println(CountTrees(board, 7, 1))
   fmt.Println(CountTrees(board, 1, 2))
   answer := 1
   answer *= CountTrees(board, 3, 1)
   answer *= CountTrees(board, 1, 1)
   answer *= CountTrees(board, 5, 1)
   answer *= CountTrees(board, 7, 1)
   answer *= CountTrees(board, 1, 2)
   fmt.Println(answer)
}