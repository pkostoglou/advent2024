package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func dayThree() {
	re1 := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	re2 := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)|don\'t\(\)|do\(\)`)
	sum := 0
	sum2 := 0

	file, err := os.Open("./day3input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	hasDont := false
	for scanner.Scan() {
		line := scanner.Text()
		matches := re1.FindAllString(line, -1)
		for _, mulAsString := range matches {
			extractedNumbers := mulAsString[4 : len(mulAsString)-1]
			numbersAsStrings := strings.Split(extractedNumbers, ",")
			number1, _ := strconv.Atoi(numbersAsStrings[0])
			number2, _ := strconv.Atoi(numbersAsStrings[1])
			sum += number1 * number2
		}
		matches2 := re2.FindAllString(line, -1)
		for _, mulOrCommandAsString := range matches2 {
			if mulOrCommandAsString == "don't()" {
				hasDont = true
				continue
			}
			if mulOrCommandAsString == "do()" {
				hasDont = false
				continue
			}
			if hasDont {
				continue
			}
			extractedNumbers := mulOrCommandAsString[4 : len(mulOrCommandAsString)-1]
			numbersAsStrings := strings.Split(extractedNumbers, ",")
			number1, _ := strconv.Atoi(numbersAsStrings[0])
			number2, _ := strconv.Atoi(numbersAsStrings[1])
			sum2 += number1 * number2
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("first answer is: ", sum)
	fmt.Println("second answer is: ", sum2)
}
