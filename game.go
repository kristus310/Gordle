package main

import (
	"bufio"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"math/rand"
	"os"
	"time"
)

var turns = 1

func WordGenerator() string {
	seed := rand.NewSource(time.Now().UnixNano())
	randomNumber := rand.New(seed)
	lines := countLines()
	targetLine := randomNumber.Intn(lines)

	scanner, file := readFile()
	defer file.Close()

	currentLine := 1
	text := ""
	for scanner.Scan() {
		if currentLine == targetLine {
			text = scanner.Text()
			break
		}
		currentLine++
	}
	if text == "" {
		dialog.ShowInformation("Error", "While choosing random word, something failed!", w)
		os.Exit(1)
	}
	return text
}

func readFile() (*bufio.Scanner, os.File) {
	file, err := os.Open("wordlist.txt")
	errorCheck(err)
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		dialog.ShowInformation("Error", "Opening the word wordlist failed!", w)
		os.Exit(1)
	}
	return scanner, *file
}

func countLines() int {
	scanner, file := readFile()
	defer file.Close()

	lines := 0
	for scanner.Scan() {
		lines++
	}
	return lines
}

func errorCheck(err error) {
	if err != nil {
		dialog.ShowInformation("Error", "Opening the word wordlist failed!", w)
		return
	}
}

func dictionaryCheck(guess string) bool {
	scanner, file := readFile()
	defer file.Close()
	for scanner.Scan() {
		if scanner.Text() == guess {
			return true
		}
	}
	return false
}

func checkStatus(guess string) {
	var win, lose bool
	if guess == word {
		win = true
	} else if turns == 6 {
		lose = true
	} else {
		return
	}

	input.Disable()
	go func() {
		if win {
			gameWin()
		} else if lose {
			gameOver()
		}
	}()
	fyne.Do(func() {
		playAgain := widget.NewButton("Play again!", func() {
			start = 0
			end = 5
			turns = 1
			w.Close()
			startGame()
		})
		time.Sleep(2 * time.Second)
		w.SetContent(container.NewPadded(container.NewBorder(nil, playAgain, nil, nil, grid)))
	})
}

func gameLogic(guess string) {
	entryUpdate(guess)
	checkStatus(guess)
	turns++
}
