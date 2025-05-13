package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"slices"
	"strings"
	"time"
)

var w fyne.Window

var entries [30]*canvas.Text
var border [30]*canvas.Rectangle
var input *widget.Entry
var grid *fyne.Container

var green = color.NRGBA{R: 76, G: 175, B: 80, A: 255}
var yellow = color.NRGBA{R: 255, G: 235, B: 59, A: 255}
var gray = color.NRGBA{R: 158, G: 158, B: 158, A: 255}
var red = color.NRGBA{R: 213, G: 0, B: 0, A: 255}

var start = 0
var end = 5

func SetupGUI() {
	w = myApp.NewWindow("Gordle")

	input = widget.NewEntry()
	input.SetPlaceHolder("your guess...")
	input.OnChanged = func(text string) {
		if len(text) >= 5 {
			var temp []string
			for i := 0; i < 5; i++ {
				temp = append(temp, string(text[i]))
			}
			input.Text = strings.Join(temp, "")
		}
	}
	input.OnSubmitted = func(guess string) {
		guess = strings.ToLower(guess)
		correctWord := dictionaryCheck(guess)
		if correctWord {
			input.SetText("")
			gameLogic(guess)
		}
	}

	grid = container.NewGridWithColumns(5)
	var content *fyne.Container
	for i := range entries {
		i := i
		border[i] = canvas.NewRectangle(color.NRGBA{R: 120, G: 120, B: 120, A: 255})
		entries[i] = canvas.NewText("", theme.Color(theme.ColorNameForeground))
		entries[i].TextSize = 50
		entries[i].TextStyle = fyne.TextStyle{Bold: true}
		content = container.New(layout.NewPaddedLayout(), border[i], container.NewCenter(entries[i]))
		grid.Add(content)
	}

	border := canvas.NewRectangle(color.NRGBA{R: 120, G: 120, B: 120, A: 255})
	border.StrokeWidth = 2
	border.StrokeColor = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	border.FillColor = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x19}

	w.SetContent(container.NewPadded(container.NewBorder(nil, input, nil, nil, grid)))
	w.Canvas().Focus(input)
	w.Resize(fyne.NewSize(450, 700))
	w.CenterOnScreen()
	w.SetFixedSize(true)
	myApp.Settings().SetTheme(&myTheme{})
	w.Show()
}

func entryUpdate(guess string) {
	temp := make([]string, 5)
	for i := start; i < end; i++ {
		temp[i-start] = string(word[i-start])
		entries[i].Text = strings.ToUpper(string(guess[i-start]))
		entries[i].Refresh()
		if guess[i-start] == word[i-start] {
			border[i].FillColor = green
			border[i].Refresh()
			temp[i-start] = ""
		}
	}

	for i := start; i < end; i++ {
		if guess[i-start] != word[i-start] {
			if slices.Contains(temp, string(guess[i-start])) {
				border[i].FillColor = yellow
				border[i].Refresh()
				for j := range temp {
					if temp[j] == string(guess[i-start]) {
						temp[j] = ""
						break
					}
				}
			} else {
				border[i].FillColor = gray
				border[i].Refresh()
			}
		}
	}

	start += 5
	end += 5
}

func gameOver() {
	time.Sleep(2 * time.Second)
	for i := range border {
		if i >= 10 && i <= 14 {
			message := "WORD:"
			border[i].FillColor = green
			entries[i].Text = string(message[i-10])
		} else if i >= 15 && i <= 19 {
			border[i].FillColor = green
			entries[i].Text = string(word[i-15])
		} else {
			border[i].FillColor = red
			entries[i].Text = ""

		}
		fyne.Do(func() {
			border[i].Refresh()
			entries[i].Refresh()
		})
	}
}

func gameWin() {
	time.Sleep(2 * time.Second)
	for i := range border {
		border[i].FillColor = green
	}
}
