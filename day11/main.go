package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var file = flag.String("file", "input.txt", "Input file for processing")

var winTitle string = "Go-SDL2 Render"
var winWidth, winHeight int32 = 1280, 720

const (
	fontPath = "test.ttf"
	fontSize = 32
)

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

func parse(text []string) [][]rune {
	x := len(text)
	board := make([][]rune, x)
	for i := range board {
		board[i] = make([]rune, len(text[i]))
	}
	for i := 0; i < len(text); i++ {
		t := text[i]
		for j := 0; j < len(text[i]); j++ {
			board[i][j] = rune(t[j])
		}
	}

	return board
}

func adjSeats(seats [][]rune, i, j int) int {
	seatsCount := 0

	index := i - 1
	for index >= 0 {
		if seats[index][j] == '#' {
			seatsCount++
			break
		} else if seats[index][j] == 'L' {
			break
		}
		index--
	}
	index = i + 1
	for index < (len(seats)) {
		if seats[index][j] == '#' {
			seatsCount++
			break
		} else if seats[index][j] == 'L' {
			break
		}
		index++
	}
	jindex := j - 1
	for jindex >= 0 {
		if seats[i][jindex] == '#' {
			seatsCount++
			break
		} else if seats[i][jindex] == 'L' {
			break
		}
		jindex--
	}
	jindex = j + 1
	for jindex < (len(seats[i])) {
		if seats[i][jindex] == '#' {
			seatsCount++
			break
		} else if seats[i][jindex] == 'L' {
			break
		}
		jindex++
	}

	index = i - 1
	jindex = j - 1
	for index >= 0 && jindex >= 0 {
		if seats[index][jindex] == '#' {
			seatsCount++
			break
		} else if seats[index][jindex] == 'L' {
			break
		}
		index--
		jindex--
	}
	index = i - 1
	jindex = j + 1
	for index >= 0 && jindex < len(seats[index]) {
		if seats[index][jindex] == '#' {
			seatsCount++
			break
		} else if seats[index][jindex] == 'L' {
			break
		}
		index--
		jindex++
	}

	index = i + 1
	jindex = j - 1
	for index < (len(seats)) && jindex >= 0 {
		if seats[index][jindex] == '#' {
			seatsCount++
			break
		} else if seats[index][jindex] == 'L' {
			break
		}
		index++
		jindex--
	}

	index = i + 1
	jindex = j + 1
	for index < (len(seats)) && jindex < len(seats[index]) {
		if seats[index][jindex] == '#' {
			seatsCount++
			break
		} else if seats[index][jindex] == 'L' {
			break
		}
		index++
		jindex++
	}

	return seatsCount
}

func copy(seats [][]rune) [][]rune {
	x := len(seats)
	board := make([][]rune, x)
	for i := range board {
		board[i] = make([]rune, len(seats[i]))
	}
	for i := 0; i < len(seats); i++ {
		for j := 0; j < len(seats[i]); j++ {
			board[i][j] = seats[i][j]
		}
	}

	return board
}

func countBoard(seats [][]rune) int {
	seatCount := 0
	for i := 0; i < len(seats); i++ {
		for j := 0; j < len(seats[i]); j++ {
			if seats[i][j] == '#' {
				seatCount++
			}
		}
	}

	return seatCount
}

func simulate(seats [][]rune) (bool, [][]rune) {
	b := copy(seats)
	changed := false
	for i := 0; i < len(seats); i++ {
		for j := 0; j < len(seats[i]); j++ {
			if b[i][j] == '.' {
				//fmt.Printf(". ")
				continue
			} else if b[i][j] == '#' {
				seatCount := adjSeats(seats, i, j)
				if seatCount >= 5 {
					b[i][j] = 'L'
					changed = true
				}
				//fmt.Printf("%v ", adjSeats(seats, i, j))
				//	fmt.Printf("# ")
			} else if b[i][j] == 'L' {
				seatCount := adjSeats(seats, i, j)
				if seatCount == 0 {
					b[i][j] = '#'
					changed = true
				}
				//fmt.Printf("%v ", adjSeats(seats, i, j))
				//	fmt.Printf("%c ", seats[i][j])
			}

			//fmt.Printf("%v ", adjSeats(seats, i, j))
			//fmt.Printf("%c ", seats[i][j])
		}
		//	fmt.Println()
	}
	//	fmt.Println()
	return changed, b
}

func initsdl() (*sdl.Window, *sdl.Renderer, *ttf.Font) {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var font *ttf.Font

	window, err := sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		os.Exit(1)
	}

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		os.Exit(1)
	}

	// Load the font for our text
	if font, err = ttf.OpenFont(fontPath, fontSize); err != nil {
		os.Exit(1)
	}

	return window, renderer, font
}

func renderBoard(window *sdl.Window, renderer *sdl.Renderer, font *ttf.Font, seats [][]rune, h, w int32) {
	//var points []sdl.Point
	//points = []sdl.Point{sdl.Point{X: 0, Y: 0}}
	var rect sdl.Rect
	//var font *ttf.Font
	renderer.SetDrawColor(0, 0, 0, 0)
	renderer.Clear()

	/*surface, err := window.GetSurface()
	if err != nil {
		return
	} */
	//var text *sdl.Surface
	for i := 0; i < len(seats); i++ {
		for j := 0; j < len(seats[i]); j++ {
			//points = append(points, sdl.Point{X: int32(i), Y: int32(j)})
			if seats[i][j] == '.' {
				renderer.SetDrawColor(255, 255, 255, 255)
			} else if seats[i][j] == '#' {
				renderer.SetDrawColor(150, 50, 0, 0)
			} else if seats[i][j] == 'L' {
				renderer.SetDrawColor(0, 50, 100, 0)
			}

			/*// Create a red text with the font
			if text, err = font.RenderUTF8Blended("L", sdl.Color{R: 255, G: 0, B: 0, A: 255}); err != nil {
				return
			}
			// Draw the text around the center of the window
			if err := text.Blit(nil, surface, &sdl.Rect{X: -((int32(i) * w) / 2), Y: ((int32(j) * h) / 2), W: w, H: h}); err != nil {
				return
			} */
			rect = sdl.Rect{int32(i) * w, int32(j) * h, w, h}
			//renderer.FillRect(&rect)
			renderer.DrawRect(&rect)

		}
	}

	//defer text.Free()

	//renderer.SetDrawColor(0, 255, 255, 255)
	//renderer.DrawLines(points)
	//window.UpdateSurface()
	renderer.Present()
	sdl.PollEvent()
	sdl.Delay(500)

}

func main() {
	flag.Parse()
	text := readFile(*file)
	b := parse(text)
	//h := int32(len(b))
	//w := int32(len(b[0]))

	if err := ttf.Init(); err != nil {
		os.Exit(1)
	}
	//window, renderer, font := initsdl()
	//defer ttf.Quit()
	//defer font.Close()
	//defer window.Destroy()
	//defer renderer.Destroy()
	changed, board := simulate(b)

	for changed {
		changed, b1 := simulate(board)
		if !changed {
			break
		}
		board = copy(b1)
		//renderBoard(window, renderer, board, winHeight/h, winWidth/w)
		//renderBoard(window, renderer, font, b1, winHeight/h, winWidth/w)
		//fmt.Println(board)
		//fmt.Println(changed)
	}
	//renderBoard(window, renderer, font, board, winHeight/h, winWidth/w)
	sdl.Delay(4000)
	//fmt.Println(board)
	fmt.Println(countBoard(board))

}
