package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

func findMatch(lines []string, pattern string) int {
	r := regexp.MustCompile(pattern)
	result := 0
	for i := range lines {
		v := r.FindAllString(lines[i], -1)
		result += len(v)
	}
	return result
}

func forward(lines []string) int {
	return findMatch(lines, "XMAS")
}

func backwards(lines []string) int {
	return findMatch(lines, "SAMX")
}

func part01(filePath string) int {
	// normal
	data := readFile(filePath)
	
	dataCols := []string{}
	// columns
	for y := range len(data) {
		newLine := []byte{}
		for x := range data[y] {
			newLine = append(newLine, data[x][y])
		}
		dataCols = append(dataCols, string(newLine))
	}
	
	// diagonal
	diag := []string{}
	for j := range data {
		newLine1 := []byte{}
		newLine2 := []byte{}
		newLine3 := []byte{}
		newLine4 := []byte{}
		for i := range data {
			newLine1 = append(newLine1, data[i][j])
			newLine2 = append(newLine2, data[j][i])
			newLine3 = append(newLine3, data[i][len(data) - 1 - j])
			newLine4 = append(newLine4, data[j][len(data) - 1 - i])
			j++
			if j >= len(data) {
				break
			}
		}
		diag = append(diag, string(newLine1))
		diag = append(diag, string(newLine3))
		// the two diagonals from corner to corner must each exist only once
		if len(newLine2) != len(data) {
			diag = append(diag, string(newLine2))
		}
		if len(newLine4) != len(data) {
			diag = append(diag, string(newLine4))
		}
	}
	
	data = append(data, dataCols...)
	data = append(data, diag...)
	result := forward(data)
	result += backwards(data)
	return result
}

func checkPart02(part []string) bool {
	// check single letters
	if string(part[1][1]) != "A" {
		return false
	}
	if string(part[0][0]) != "M" && string(part[0][0]) != "S" {
		return false
	}
	if string(part[0][2]) != "M" && string(part[0][2]) != "S" {
		return false
	}
	if string(part[2][0]) != "M" && string(part[2][0]) != "S" {
		return false
	}
	if string(part[2][2]) != "M" && string(part[2][2]) != "S" {
		return false
	}
	// check for "SAM" and "MAS"
	if string(part[0][0]) == "M" && string(part[2][2]) != "S" {
		return false
	}
	if string(part[0][0]) == "S" && string(part[2][2]) != "M" {
		return false
	}
	if string(part[0][2]) == "M" && string(part[2][0]) != "S" {
		return false
	}
	if string(part[0][2]) == "S" && string(part[2][0]) != "M" {
		return false
	}
	
	return true
}

func part02(filePath string) int {
	data := readFile(filePath)
	result := 0
	for y := 0; y < len(data) - 2; y++ {
		for x := 0; x < len(data) - 2; x++ {
			part := []string{}
			part = append(part, string(data[y][x:(x+3)]))
			part = append(part, string(data[y + 1][x:(x+3)]))
			part = append(part, string(data[y + 2][x:(x+3)]))
			if checkPart02(part) {
				result++
			}
		}
	}
	return result
}

func main() {
	fmt.Println(part01("input.txt"))
	fmt.Println(part02("input.txt"))
}
