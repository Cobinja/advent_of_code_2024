package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

func part01(filePath string) int {
	data := readFile(filePath)
	left := []int{}
	right := []int{}
	for i := range data {
		line := data[i]
		split := strings.Split(line, "   ")
		fmt.Println(split)
		l, e := strconv.Atoi(split[0])
		check(e)
		r, e := strconv.Atoi(split[1])
		check(e)
		left = append(left, l)
		right = append(right, r)
	}
	slices.Sort(left)
	slices.Sort(right)
	totalDistance := 0
	
	for i := range left {
		dist := Abs(left[i] - right[i])
		totalDistance += dist
	}
	
	return totalDistance
}

func part02(filePath string) int {
	data := readFile(filePath)
	left := []int{}
	right := []int{}
	for i := range data {
		line := data[i]
		split := strings.Split(line, "   ")
		l, e := strconv.Atoi(split[0])
		check(e)
		r, e := strconv.Atoi(split[1])
		check(e)
		left = append(left, l)
		right = append(right, r)
	}
	slices.Sort(left)
	slices.Sort(right)
	
	simScore := 0
	
	for i := range left {
		l := left[i]
		if l == 0 {
			break
		}
		
		start := slices.Index(right, l)
		if start == -1 {
			continue
		}
		rightSlice := right[start:]
		for i := range rightSlice {
			if rightSlice[i] != l {
				simScore += l * i
				break
			}
		}
	}
	fmt.Println("simScore: ", simScore)
	return simScore
}

func main() {
	// fmt.Println(part01("input.txt"))
	fmt.Println(part02("input.txt"))
}
