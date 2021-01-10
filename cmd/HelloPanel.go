package main

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"pdfcpuSamples/domain"
)

func main() {

	a := app.New()
	a.Settings().SetTheme(theme.LightTheme())
	w := a.NewWindow("Hello Card or Panel (named group of widgets")

	targetDirField := widget.NewLabel("/Users/gernotstarke/_target")
	targetDirButton := widget.NewButton("Target", func() {
		targetDirField.SetText( fmt.Sprintf("%v", domain.TargetDir))
	})

	srcDirField := widget.NewLabel("another directory")
	srcDirButton := widget.NewButton( "Source", func(){})


	targetDirContainer := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		targetDirButton,
		targetDirField,
	)
	srcDirContainer := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		srcDirButton,
		srcDirField,
	)

	dirContainer := fyne.NewContainerWithLayout( layout.NewVBoxLayout(),
		srcDirContainer,
		targetDirContainer)

	container := widget.NewCard("", "Directories", dirContainer)

	w.SetContent(container)
	//w.Resize(fyne.NewSize(600, 200))
	w.ShowAndRun()

}
