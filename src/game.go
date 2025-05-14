package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"math/rand"
	"time"
)

var turns = 1

func WordGenerator() string {
	seed := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(seed)
	randomNum := rng.Intn(len(words))

	word := words[randomNum]
	return word
}

func dictionaryCheck(guess string) bool {
	for _, value := range words {
		if guess == value {
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
