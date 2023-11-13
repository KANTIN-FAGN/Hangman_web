package hangman

import (
	"bufio"
	"fmt"
	"os"
)

func DisplayPendu(start int, end int) {
	file, err := os.Open("hangman.txt")
	if err != nil {
		fmt.Println("Ya une douille sur le fichier : ", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lineCount := 0
	linesToDisplay := []string{}
	for scanner.Scan() {
		lineCount++
		line := scanner.Text()
		if lineCount >= start && lineCount <= end {
			linesToDisplay = append(linesToDisplay, line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Ya une douille sur le fichier : ", err)
		return
	}
	for _, line := range linesToDisplay {
		fmt.Println(line)
	}
}

func DisplayResult(start int, end int) {
	file, err := os.Open("jeu/dico_ascii.txt")
	if err != nil {
		fmt.Println("Ya une douille sur le fichier : ", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lineCount := 0
	linesToDisplay := []string{}
	for scanner.Scan() {
		lineCount++
		line := scanner.Text()
		if lineCount >= start && lineCount <= end {
			linesToDisplay = append(linesToDisplay, line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Ya une douille sur le fichier : ", err)
		return
	}
	for _, line := range linesToDisplay {
		fmt.Println(line)
	}
}
