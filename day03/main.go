package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(fileName string) []string{
	readFile, err := os.Open(fileName)
	check(err)
	
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	readFile.Close()
	return fileLines
}

func readFileAsOneString(fileName string) string {
	readFile, err := os.Open(fileName)
	check(err)
	
	fileScanner := bufio.NewScanner(readFile)
	// fileScanner.Split(bufio.ScanLines)
	
	var fileLines string
	for fileScanner.Scan() {
		fileLines += fileScanner.Text()
	}
	readFile.Close()
	return fileLines
}

func part01(filePath string) int {
	data := readFile(filePath)
	result := 0
	for i := range data {
		line := data[i]
		r := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
		m := r.FindAllString(line, -1)
		for j := range m {
			mul := m[j][4:(len(m[j]) - 1)]
			split := strings.Split(mul, ",")
			left, _ := strconv.Atoi(split[0])
			right, _ := strconv.Atoi(split[1])
			result +=  left * right
		}
		
	}
	return result
}

func part02(filePath string) int {
	data := readFileAsOneString(filePath)
	
	do := regexp.MustCompile(`do\(\)`)
	dont := regexp.MustCompile(`don\'t\(\)`)
	r := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	
	result := 0
	for {
		dontIdx := dont.FindStringIndex(data)
		var m []string
		if dontIdx == nil {
			m = r.FindAllString(data, -1)
		} else if dontIdx[0] >= 0 {
			m = r.FindAllString(data[:dontIdx[0]], -1)
		}
		
		for j := range m {
			mul := m[j][4:(len(m[j]) - 1)]
			split := strings.Split(mul, ",")
			left, _ := strconv.Atoi(split[0])
			right, _ := strconv.Atoi(split[1])
			result +=  left * right
		}
		
		if dontIdx == nil {
			break
		}
		data = data[dontIdx[1]:]
		doIdx := do.FindStringIndex(data)
		data = data[doIdx[1]:]
	}
	
	return result
}

func main() {
	// fmt.Println(part01("input.txt"))
	fmt.Println(part02("input.txt"))
}
