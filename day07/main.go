package main
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"flag"
	"strings"
	"container/list"
	"strconv"
	
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


type graph struct {
	nodes map[string][]string
	w map[string][]int
}

func newGraph(n [] string, c []string) (graph)  {
	var nodes map[string][]string
	var w map[string][]int
	nodes = make(map[string][]string)
	w = make(map[string][]int)
	for _ , key := range n {
		nodes[key] = make([]string, 0)
		w[key] = make([]int, 0)
	}

	for i, line  := range c {
		s := strings.Split(line, ",")
		for _, bag := range s {
			bag = strings.Trim(bag, " .")
			if bag != "no other bags" {
				quantity, _ := strconv.Atoi(strings.Trim(bag[:1], " "))
				bag = strings.Trim(bag[1:], " ")
				if bag[len(bag) - 1] == 's' {
					bag = bag[:len(bag) -1]
				}
				nodes[n[i]] = append(nodes[n[i]], bag)
				w[n[i]] = append(w[n[i]], quantity)
			}
		}
	}

  return graph{nodes: nodes, w: w}	
}

func (g graph) reverse() (graph) {
   var nodes map[string][]string
    nodes = make(map[string][]string)
	for key := range g.nodes {
		nodes[key] = make([]string, 0)
	}

	for key, value := range g.nodes {
		for _, v := range value {
			nodes[v] = append(nodes[v], key)
		}
	}
	return graph{nodes: nodes}
}

func (g graph) println () {
	for key, value := range g.nodes {
		fmt.Println(key, "= [")
		for _, v := range value {
			fmt.Println("\t", v, ",")
		}
		fmt.Println("]")
	}
}

func (g graph) printW() {
	for key, value := range g.nodes {
		fmt.Println(key, "= [")
		for i, c := range value {
			fmt.Println("\t", c, "(", g.w[key][i], ")")	
		}
		fmt.Println("]\n")
	}
}

func (g graph) bfsCount (key string) (int) {
	return len(search(key, g)) - 1
}

func (g graph) bfs(key string) (map[string]bool) {
	return search(key, g)
}
func search (key string, g graph) (map[string]bool) {
	q := list.New()
	visited := map[string]bool{}
	q.PushBack(key)
	for q.Len() > 0 {
		e  := q.Front()
		q.Remove(e)
		n := string(e.Value.(string))
		if(visited[n]) {
			continue
		}
		for _, node := range g.nodes[n] {
			q.PushBack(node) 
		}
		
		visited[n] = true		
	}

	return visited
}


func parse ( text []string )([]string, []string) {
	nodes := make([]string, 0)
	children := make([]string, 0)
	for _, line := range text {
		if line == ""  {
			continue
		}
		s := strings.Split(line, "contain")
		nodes = append(nodes, s[0][0:len(s[0]) - 2])
		children = append(children, s[1])	
	
	}
return nodes, children	 
}

func (g graph) dfsCount (key string) (int) {
	return traverse(key, g) - 1 
}

func traverse (key string, g graph) (int) {
	total := 0
	for i, node := range g.nodes[key] {
		total += traverse(node, g) * g.w[key][i] 
	}
	return total + 1 
}

func main () {
	flag.Parse()
	text := readFile(*file)
	nodes, c := parse(text)
    g := newGraph(nodes, c)
	r := g.reverse()
	visited := r.bfs("shiny gold bag")
	fmt.Println(len(visited) - 1)
	fmt.Println(g.dfsCount("shiny gold bag"))
	
	
}