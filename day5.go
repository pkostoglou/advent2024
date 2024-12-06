package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type scoreType struct {
	score int
	value int
}

func dayFive() {

	readingUpdates := false
	pagesMap := make(map[int][]int)
	count := 0
	count2 := 0

	file, err := os.Open("./day5input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			readingUpdates = true
			continue
		}
		if readingUpdates {
			numbersAsStrings := strings.Split(line, ",")
			flag := false
			middleNumber := 0
			for index, numberAsString := range numbersAsStrings[:len(numbersAsStrings)-1] {
				number, _ := strconv.Atoi(numberAsString)

				if index == len(numbersAsStrings)/2 {
					middleNumber = number
				}
				for _, restOfNumbersAString := range numbersAsStrings[index+1:] {
					numberToCheck, _ := strconv.Atoi(restOfNumbersAString)
					arrayToCheck := pagesMap[numberToCheck]
					if slices.Contains(arrayToCheck, number) {
						flag = true
					}
				}
			}
			if !flag {
				count = count + middleNumber
			} else {
				scoreArray := []scoreType{}
				for index, numberToGetScoreFor := range numbersAsStrings {
					numberForScore, _ := strconv.Atoi(numberToGetScoreFor)
					scoreArray = append(scoreArray, scoreType{score: 0, value: numberForScore})
					arrayForCheck := pagesMap[numberForScore]
					for _, numberToCompareS := range numbersAsStrings {
						numberToCompare, _ := strconv.Atoi(numberToCompareS)
						if slices.Contains(arrayForCheck, numberToCompare) {
							scoreArray[index].score++
						}
					}
				}
				sort.Slice(scoreArray, func(i, j int) bool {
					return scoreArray[i].score < scoreArray[j].score
				})
				count2 = count2 + scoreArray[len(scoreArray)/2].value
			}
		} else {
			numbersAsStrings := strings.Split(line, "|")
			leftNumber, _ := strconv.Atoi(numbersAsStrings[0])
			rightNumber, _ := strconv.Atoi(numbersAsStrings[1])
			_, ok := pagesMap[leftNumber]
			if !ok {
				pagesMap[leftNumber] = []int{rightNumber}
			} else {
				pagesMap[leftNumber] = append(pagesMap[leftNumber], rightNumber)
			}
		}

	}

	fmt.Println("First answer is: ", count)
	fmt.Println("Second answer is: ", count2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
