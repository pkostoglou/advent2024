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

func isLineSafe(numbers []int, withTolerance bool) bool {
	increasing := numbers[0] < numbers[1]
	for index, number := range numbers[1:] {
		if numbers[index] < number == increasing {
			diff := numbers[index] - number
			if (diff >= 1 && diff <= 3) || (-diff >= 1 && -diff <= 3) {
				continue
			} else {
				if withTolerance {
					original := slices.Clone(numbers)
					original2 := slices.Clone(numbers)
					one := isLineSafe(slices.Delete(numbers, index, index+1), false)
					two := isLineSafe(slices.Delete(original, index+1, index+2), false)
					three := isLineSafe(slices.Delete(original2, 0, 1), false)
					return one || two || three
				}
				return false

			}
		} else {
			if withTolerance {
				original := slices.Clone(numbers)
				original2 := slices.Clone(numbers)
				one := isLineSafe(slices.Delete(numbers, index, index+1), false)
				two := isLineSafe(slices.Delete(original, index+1, index+2), false)
				three := isLineSafe(slices.Delete(original2, 0, 1), false)
				return one || two || three
			}
			return false
		}
	}
	return true
}

func dayTwo() {
	countSafe := 0
	countSafeWithTolerance := 0

	file, err := os.Open("./day2input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbersAsStrings := strings.Split(line, " ")
		numbers := []int{}
		for _, numberAsString := range numbersAsStrings {
			number, _ := strconv.Atoi(numberAsString)
			numbers = append(numbers, number)
		}
		if isLineSafe(numbers, false) {
			countSafe++
		}
		if isLineSafe(numbers, true) {
			countSafeWithTolerance++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("first answer is :", countSafe)
	fmt.Println("second answer is :", countSafeWithTolerance)
}
