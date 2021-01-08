package main

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"log"
	"os"
	"pdfcpuSamples/ui"
)

var sourceDir string
var targetDir string

func logoHeader() *fyne.Container {
	arc42Logo := canvas.NewImageFromFile("resources/arc42-logo.png")
	arc42Logo.FillMode = canvas.ImageFillContain
	arc42Logo.SetMinSize( fyne.NewSize( 80, 40))
	arc42Logo.Resize(fyne.NewSize(80, 40))


	appLogo := canvas.NewImageFromFile("resources/PDFnumbrr-logo.png")
	appLogo.FillMode = canvas.ImageFillContain
	appLogo.SetMinSize( fyne.NewSize( 200,120))
	appLogo.Resize(fyne.NewSize(200, 120))

	container := fyne.NewContainerWithLayout(
		layout.NewHBoxLayout(),
		fyne.NewContainerWithLayout(layout.NewVBoxLayout(), arc42Logo),
		layout.NewSpacer(),
		appLogo,
	)
	return container
}

func srcDirSelectorUI() *fyne.Container {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(dir)
	srcDirField := widget.NewLabel("/Users/gernotstarke")

	srcDirButton := widget.NewButton("Source", func() {
		sourceDir = dir
		srcDirField.SetText(sourceDir)
	})

	srcDirLabel := canvas.NewText("0 PDFs found", ui.DarkBlueColor)
	srcDirLabel.TextSize = 9

	return fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		srcDirButton,
		srcDirField,
		layout.NewSpacer(),
		srcDirLabel,
	)
}

func targetDirSelectorUI() *fyne.Container {
	targetDirField := widget.NewLabel("/Users/gernotstarke/_target")

	targetDirButton := widget.NewButton("Target", func() {
		targetDir = "Users/gernotstarke/_target"
		targetDirField.SetText(fmt.Sprintf("%v", targetDir))
	})

	targetValid := widget.NewCheck("valid:", func(bool) {})
	targetValid.Disable()

	return fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		targetDirButton,
		//layout.NewSpacer(),
		targetDirField,
		layout.NewSpacer(),
		targetValid,
	)
}


// a really bad implementation, as it returns a new container with every call...
// need to be able to just modify the message

func statusLine(msg string) *fyne.Container {

	label := canvas.NewText(msg, ui.PinkColor)
	label.TextSize = 10
	label.Alignment = fyne.TextAlignTrailing

	return fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		label)
}

func main() {

	a := app.New()
	a.Settings().SetTheme(theme.LightTheme())
	w := a.NewWindow("Simple Sample Layout ")

	line := canvas.NewLine(ui.DarkBlueColor)
	line.Resize(fyne.NewSize(595, 1))
	line.StrokeWidth = 1


	line2 := canvas.NewLine(ui.DarkGrayColor)
	line2.Resize( fyne.NewSize(595,1))
	line2.StrokeWidth = 1

	container := fyne.NewContainerWithLayout( layout.NewVBoxLayout(),
		logoHeader(),
		widget.NewSeparator(),
		srcDirSelectorUI(),
		widget.NewSeparator(),
		targetDirSelectorUI(),
		widget.NewSeparator(),
		statusLine( "no source directory selected"))


	w.SetContent(container)
	w.Resize(fyne.NewSize(600, 400))
	w.SetFixedSize(true)
	w.ShowAndRun()
	w.CenterOnScreen()
}
