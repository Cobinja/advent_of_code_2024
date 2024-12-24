package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func Abs[T constraints.Integer](x T) T {
	if x < 0 {
			return -x
	}
	return x
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

func check_line_safety(line []int) bool {
	isSafe := true
		lineDirection := ""
		
		for j := range len(line) - 1 {
			if !isSafe {
				break
			}
			a := line[j]
			b := line[j + 1]
			
			if a == b {
				isSafe = false
				break
			}
			
			diff := b - a
			if Abs(diff) > 3 {
				isSafe = false
				break
			}
			direction := ""
			if diff > 0 {
				direction = "i"
			} else {
				direction = "d"
			}
			
			if direction != lineDirection {
				if lineDirection == "" {
					lineDirection = direction
				} else {
					isSafe = false
					break
				}
			}
		}
		return isSafe
}

func removeIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func part01(filePath string) int {
	data := readFile(filePath)
	safe := 0
	for i := range data {
		lineData := strings.Split(data[i], " ")
		var line []int = make([]int, 0)
		for j := range lineData {
			v, e := strconv.Atoi(lineData[j])
			check(e)
			line = append(line, v)
		}
		
		isSafe := true
		lineDirection := ""
		
		for j := range len(line) - 1 {
			if !isSafe {
				break
			}
			a := line[j]
			b := line[j + 1]
			
			if a == b {
				isSafe = false
				break
			}
			
			diff := b - a
			if Abs(diff) > 3 {
				isSafe = false
				break
			}
			direction := ""
			if diff > 0 {
				direction = "i"
			} else {
				direction = "d"
			}
			
			if direction != lineDirection {
				if lineDirection == "" {
					lineDirection = direction
				} else {
					isSafe = false
					break
				}
			}
		}
		if isSafe {
			safe ++
		}
	}
	fmt.Println(safe)
	return safe
}

func part02(filePath string) int {
	data := readFile(filePath)
	safe := 0
	for i := range data {
		lineData := strings.Split(data[i], " ")
		var line []int = make([]int, 0)
		for j := range lineData {
			v, e := strconv.Atoi(lineData[j])
			check(e)
			line = append(line, v)
		}
		isSafe := check_line_safety(line)
		
		
		for i := range line {
			newLine := removeIndex(line, i)
			isSafe = check_line_safety(newLine)
			if isSafe {
				break
			}
		}
	
		
		if isSafe {
			safe++
		}
	}
	return safe
}

func main() {
	// fmt.Println(part01("input.txt"))
	fmt.Println(part02("input.txt"))
}
