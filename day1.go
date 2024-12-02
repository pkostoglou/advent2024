package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func addOnArraySorted(newNumber int, array []int) []int {
	for index, arrayNumber := range array {
		if arrayNumber < newNumber {
			array = slices.Insert(array, index, newNumber)
			return array
		}
	}
	array = append(array, newNumber)
	return array
}

func dayOne() {
	leftList := []int{}
	rightList := []int{}

	file, err := os.Open("./day1input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbersAsStrings := strings.Split(line, "   ")
		leftNumber, _ := strconv.Atoi(numbersAsStrings[0])
		tempLeftList := addOnArraySorted(leftNumber, leftList)
		rightNumber, _ := strconv.Atoi(numbersAsStrings[1])
		tempRightList := addOnArraySorted(rightNumber, rightList)
		leftList = tempLeftList
		rightList = tempRightList
	}

	// part one
	sum := 0
	for index, rightNumberToCheck := range rightList {
		if rightNumberToCheck > leftList[index] {
			diff := rightNumberToCheck - leftList[index]
			sum = sum + diff
		} else {
			diff := leftList[index] - rightNumberToCheck
			sum = sum + diff
		}
	}
	fmt.Println("Part one answer is: ", sum)

	// part two
	similarity := 0
	for _, rightNumberToCheck := range rightList {
		timesFound := 0
		for _, leftNumberToCheck := range leftList {
			if rightNumberToCheck > leftNumberToCheck {
				break
			}
			if rightNumberToCheck == leftNumberToCheck {
				timesFound = timesFound + 1
			}
		}
		similarity = similarity + rightNumberToCheck*timesFound
	}
	fmt.Println("Part two answer is: ", similarity)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
