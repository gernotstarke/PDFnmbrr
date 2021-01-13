package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"pdfcpuSamples/resources"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(theme.LightTheme())
	w := a.NewWindow("Hello Images via resource")

	appLogo1 := canvas.NewImageFromResource(resources.Arc42LogoPNG )
	appLogo1.Resize(fyne.NewSize(200, 100))
	appLogo1.FillMode = canvas.ImageFillOriginal

	//appLogo3 := canvas.NewImageFromResource(resourcePDFnmbrrLogoPng )
	//appLogo3.FillMode = canvas.ImageFillOriginal
	//appLogo3.Resize( fyne.NewSize( 200, 210))


	container := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		appLogo1,
//		layout.NewSpacer(),
//		appLogo3,
	)

	w.SetContent(container)
	w.Resize(fyne.NewSize(600, 200))
	w.ShowAndRun()

}

