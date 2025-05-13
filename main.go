package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var word string
var myApp fyne.App

func main() {
	myApp = app.New()
	startGame()
	myApp.Run()
}

func startGame() {
	word = WordGenerator()
	SetupGUI()
}
