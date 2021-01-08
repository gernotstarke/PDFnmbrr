package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(theme.LightTheme())
	w := a.NewWindow("Hello Images via direct")

	appLogo1 := canvas.NewImageFromFile("resources/arc42-logo" +
		".png")
	appLogo1.Resize(fyne.NewSize(200, 100))
	appLogo1.FillMode = canvas.ImageFillOriginal

	appLogo3 := canvas.NewImageFromFile("resources/PDFnumbrr-logo.png")
	appLogo3.FillMode = canvas.ImageFillOriginal
	appLogo3.Resize( fyne.NewSize( 200, 210))

	container := fyne.NewContainerWithLayout( layout.NewHBoxLayout(), logoHeader2())

//	container := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
//		appLogo1,
//		layout.NewSpacer(),
//		appLogo3,
//	)

	w.SetContent(container)
	w.Resize(fyne.NewSize(600, 200))
	w.ShowAndRun()

}

func logoHeader2() fyne.CanvasObject {
	appLogo1 := canvas.NewImageFromFile("resources/arc42-logo.png")
	appLogo1.FillMode = canvas.ImageFillOriginal
	appLogo1.Resize(fyne.NewSize(80, 40))

	appLogo3 := canvas.NewImageFromFile("resources/PDFnumbrr-logo.png")
	appLogo3.FillMode = canvas.ImageFillOriginal
	appLogo3.Resize(fyne.NewSize(100, 120))

	container := fyne.NewContainerWithLayout(
		layout.NewGridLayout(4),
		//fyne.NewContainerWithLayout(layout.NewBorderLayout( appLogo3, nil, nil, nil)) ,
		fyne.NewContainerWithLayout(layout.NewVBoxLayout(), appLogo1),
		layout.NewSpacer(),
		layout.NewSpacer(),
		appLogo3,
	)
	return container
}
