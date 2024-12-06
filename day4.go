package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var searching_map = [4]byte{'X', 'M', 'A', 'S'}

func horizontal(posX int, posY int, array []string, reverse bool) bool {
	line := array[posY]
	currentPos := posX
	currentSearchingIndex := 0
	for posX <= len(line)-1 {
		if currentPos >= len(line) && !reverse {
			break
		}
		if currentPos < 0 && reverse {
			break
		}
		if currentSearchingIndex >= 4 {
			break
		}
		if line[currentPos] != searching_map[currentSearchingIndex] {
			break
		}
		if reverse {
			currentPos--
		} else {
			currentPos++
		}
		currentSearchingIndex++
	}
	return currentSearchingIndex == 4
}

func vertical(posX int, posY int, array []string, reverse bool) bool {
	if reverse {
		if posY-3 < 0 {
			return false
		}
		line := []byte{array[posY][posX], array[posY-1][posX], array[posY-2][posX], array[posY-3][posX]}
		lineS := string(line)
		if lineS == "XMAS" {
			return true
		}

	} else {
		if posY+3 >= len(array) {
			return false
		}
		line := []byte{array[posY][posX], array[posY+1][posX], array[posY+2][posX], array[posY+3][posX]}
		lineS := string(line)
		if lineS == "XMAS" {
			return true
		}
	}

	return false
}

func diagonalToBottom(posX int, posY int, array []string, reverse bool) bool {
	if reverse {
		if posY-3 < 0 || posX-3 < 0 {
			return false
		}
		line := []byte{array[posY][posX], array[posY-1][posX-1], array[posY-2][posX-2], array[posY-3][posX-3]}
		lineS := string(line)
		if lineS == "XMAS" {
			return true
		}
	} else {
		if posY+3 >= len(array) || posX+3 >= len(array[posX]) {
			return false
		}
		line := []byte{array[posY][posX], array[posY+1][posX+1], array[posY+2][posX+2], array[posY+3][posX+3]}
		lineS := string(line)
		if lineS == "XMAS" {
			return true
		}
	}
	return false
}

func diagonalToTop(posX int, posY int, array []string, reverse bool) bool {
	if reverse {
		if posY+3 >= len(array) || posX-3 < 0 {
			return false
		}
		line := []byte{array[posY][posX], array[posY+1][posX-1], array[posY+2][posX-2], array[posY+3][posX-3]}
		lineS := string(line)
		if lineS == "XMAS" {
			return true
		}
	} else {
		if posY-3 < 0 || posX+3 >= len(array[posX]) {
			return false
		}
		line := []byte{array[posY][posX], array[posY-1][posX+1], array[posY-2][posX+2], array[posY-3][posX+3]}
		lineS := string(line)
		if lineS == "XMAS" {
			return true
		}
	}
	return false
}

func diagonalToBottomMas(posX int, posY int, array []string, reverse bool) bool {
	if reverse {
		if posY-2 < 0 || posX-2 < 0 {
			return false
		}
		line := []byte{array[posY][posX], array[posY-1][posX-1], array[posY-2][posX-2]}
		lineS := string(line)
		if lineS == "MAS" {
			return true
		}
	} else {
		if posY+2 >= len(array) || posX+2 >= len(array[posX]) {
			return false
		}
		line := []byte{array[posY][posX], array[posY+1][posX+1], array[posY+2][posX+2]}
		lineS := string(line)
		if lineS == "MAS" {
			return true
		}
	}
	return false
}

func diagonalToTopMas(posX int, posY int, array []string, reverse bool) bool {
	if reverse {
		if posY+2 >= len(array) || posX-2 < 0 {
			return false
		}
		line := []byte{array[posY][posX], array[posY+1][posX-1], array[posY+2][posX-2]}
		lineS := string(line)
		if lineS == "MAS" {
			return true
		}
	} else {
		if posY-2 < 0 || posX+2 >= len(array[posX]) {
			return false
		}
		line := []byte{array[posY][posX], array[posY-1][posX+1], array[posY-2][posX+2]}
		lineS := string(line)
		if lineS == "MAS" {
			return true
		}
	}
	return false
}

func isXShapeMas(posX int, posY int, array []string) bool {
	masOne := []byte{array[posY-1][posX-1], array[posY][posX], array[posY+1][posX+1]}
	masOneS := string(masOne)
	masOneRevert := []byte{array[posY+1][posX+1], array[posY][posX], array[posY-1][posX-1]}
	masOneRevertS := string(masOneRevert)

	if masOneS != "MAS" && masOneRevertS != "MAS" {
		return false
	}

	masTwo := []byte{array[posY-1][posX+1], array[posY][posX], array[posY+1][posX-1]}
	masTwoS := string(masTwo)
	masTwoRevert := []byte{array[posY+1][posX-1], array[posY][posX], array[posY-1][posX+1]}
	masTwoRevertS := string(masTwoRevert)

	if masTwoS != "MAS" && masTwoRevertS != "MAS" {
		return false
	}

	return true
}

func dayFour() {
	array := []string{}
	count := 0
	countXShapeMas := 0

	file, err := os.Open("./day4input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		array = append(array, line)
	}
	for indexY, line := range array {
		for indexX := range line {
			if horizontal(indexX, indexY, array, false) {
				count++
			}

			if horizontal(indexX, indexY, array, true) {
				count++
			}

			if vertical(indexX, indexY, array, false) {
				count++
			}

			if vertical(indexX, indexY, array, true) {
				count++
			}

			if diagonalToBottom(indexX, indexY, array, false) {
				count++
			}

			if diagonalToBottom(indexX, indexY, array, true) {
				count++
			}

			if diagonalToTop(indexX, indexY, array, false) {
				count++
			}

			if diagonalToTop(indexX, indexY, array, true) {
				count++
			}

			if diagonalToBottomMas(indexX, indexY, array, false) {
				if isXShapeMas(indexX+1, indexY+1, array) {
					countXShapeMas++
				}
			}

			if diagonalToBottomMas(indexX, indexY, array, true) {
				if isXShapeMas(indexX-1, indexY-1, array) {
					countXShapeMas++
				}
			}

			if diagonalToTopMas(indexX, indexY, array, false) {
				if isXShapeMas(indexX+1, indexY-1, array) {
					countXShapeMas++
				}
			}

			if diagonalToTopMas(indexX, indexY, array, true) {
				if isXShapeMas(indexX-1, indexY+1, array) {
					countXShapeMas++
				}
			}

		}
	}

	fmt.Println("First answer is: ", count)
	fmt.Println("Second answer is: ", countXShapeMas/2)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
