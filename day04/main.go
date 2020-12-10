package main
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	"regexp"
	
)

type ID struct {
	data string
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}


var IsAlpha = regexp.MustCompile(`^[0-9a-f]+$`).MatchString
var IsNum = regexp.MustCompile(`^[0-9]+$`).MatchString

var colorSet = map[string]bool{
	"blu": true, 
	"amb":true, 
	"grn": true, 
	"brn": true, 
	"gry": true, 
	"hzl": true, 
	"oth": true,
}

func ( i ID) validate () bool {

	if i.hgt == "" || i.iyr == "" || i.byr == "" || i.eyr == "" || i.hcl == "" || i.pid == "" {
		fmt.Println("Missing fields")
		return false
	}

	if (len(i.hcl) != 7 ) {
		fmt.Println("Bad hcl")
		return false
	}


	 iyr, _ := strconv.Atoi(i.iyr)
	 byr, _ := strconv.Atoi(i.byr)
	 eyr, _ := strconv.Atoi(i.eyr)
	 hgt, _ := strconv.Atoi(i.hgt[:len(i.hgt) - 2])
	 hgtunit := i.hgt[len(i.hgt) -2:]
	 hcl := i.hcl[:1]
	 hclcl := i.hcl[1:]

	 if (iyr < 2010 || iyr > 2020) {
		fmt.Println("bad issue year")
		 return false
	 }

    if (byr < 1920 || byr > 2002) {
		fmt.Println("bad birth year")
		 return false
	 }
	 if (eyr < 2020 || eyr > 2030) {

		fmt.Println("bad expiration datae")
		 return false
	 }
     if hgtunit != "in" && hgtunit != "cm" {
		fmt.Println("bad height units")
		 return false
	 }

	 if hgtunit == "cm" && (hgt < 150 || hgt > 193) {
		fmt.Println("bad heigh in cm")
		 return false
	 }
	 if hgtunit == "in" && (hgt < 59 || hgt > 76) {
		fmt.Println("bad height in in")
		 return false
	 }

	 if hcl != "#" {
		fmt.Println("missing field in hcl")
		 return false
	 } 

	 if len(hclcl) != 6 || IsAlpha(hclcl) != true {
		fmt.Println("bad hair color hex")
		 return false
	 } 
	 if _, ok := colorSet[i.ecl]; !ok {
		fmt.Println("bad eye color")
		 return false
	 } 

	 if len(i.pid) != 9 || IsNum(i.pid) != true {
		fmt.Println("pid is worng")
		 return false
	 } 

return true

}

func new (d string) (ID) {

	fields := strings.Split(d, " ")
	var r ID
	r.data = d
	for _, f := range fields {
		if f != ""{
			v := strings.Split(f, ":")
			
			switch v[0] {
			case "byr":
				r.byr = v[1]	
			case "iyr":
				r.iyr = v[1]
			case "eyr":
				r.eyr = v[1]
			case "hgt":
				r.hgt = v[1]
			case "hcl":
				r.hcl = v[1]
			case "ecl":
				r.ecl = v[1]
			case "pid":
				r.pid = v[1]
			case "cid":
				r.cid = v[1]
			}

		}
	}

	return r
}


func parse ( text []string ) ([]ID) {

	var passport strings.Builder
	var ids [] ID
	for i := 0; i < len(text);i++ {
		//fmt.Println(text[i])
		if text[i] != "" {
			passport.WriteString(text[i] + " ")
		} else {
		
			ids = append(ids, new(passport.String()))
			passport.Reset()
		}
	} 
	return ids
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
	ids := parse(text)

    var count int
	for _, id := range ids {
		fmt.Println(id.hgt)
		b := id.validate()
		fmt.Println(b)
		if b {
			fmt.Println(id.hgt)
			count++
		} 
	}

	fmt.Println(count)
}