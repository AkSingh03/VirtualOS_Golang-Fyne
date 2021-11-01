package main

import (

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var myApp fyne.App = app.New()

var myWindow fyne.Window = myApp.NewWindow("Ak OS")

var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var btn4 fyne.Widget

var img fyne.CanvasObject;
var DeskBtn fyne.Widget

var panelContent *fyne.Container

//Functios for Buttons

func main(){

	myApp.Settings().SetTheme(theme.DarkTheme())
	img = canvas.NewImageFromFile("103.jpeg")

	btn1 = widget.NewButtonWithIcon("Ak's Weather App" , theme.InfoIcon() , func() {

		showWeatherApp(myWindow)

	})

	btn2 = widget.NewButtonWithIcon("Ak's Calc" , theme.ContentAddIcon() , func() {

		showCalc(myWindow)

	})

	btn3 = widget.NewButtonWithIcon("Ak's Gallery" , theme.ColorPaletteIcon() , func() {

		showGallery(myWindow)

	})

	btn4 = widget.NewButtonWithIcon("Ak's TextEditor" , theme.ContentRedoIcon(), func() {

		showTextEditor(myWindow)

	})

	DeskBtn = widget.NewButtonWithIcon("Ak's Desktop" , theme.HomeIcon(), func() {

		myWindow.SetContent(container.NewBorder(panelContent ,nil,nil,nil, img))

	})

	panelContent = container.NewVBox((container.NewGridWithColumns(4, btn1 , btn2 , btn3 , btn4)))

	myWindow.Resize(fyne.NewSize(1280,720))
	myWindow.CenterOnScreen()

	myWindow.SetContent(
		container.NewBorder(panelContent, nil, nil, nil, img),
	)
	
	myWindow.ShowAndRun()
}