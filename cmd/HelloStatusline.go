package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"pdfcpuSamples/ui"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(theme.LightTheme())
	w := a.NewWindow("Hello Status Line")



	container := fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		layout.NewSpacer(),
		widget.NewSeparator(),
		statusLine2("no status available"),

	)

	w.SetContent(container)
	w.Resize(fyne.NewSize(600, 200))
	w.ShowAndRun()

}

func statusLine2(msg string) *fyne.Container {

	statusMsg := canvas.NewText(msg, ui.DarkRedColor)
	statusMsg.TextSize = 10
	statusMsg.Alignment = fyne.TextAlignTrailing

	byArc42 := canvas.NewText("created by arc42.org", ui.NavyColor)
	byArc42.TextSize = 10
	byArc42.Alignment = fyne.TextAlignCenter

	return fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		byArc42,
		widget.NewSeparator(),
		layout.NewSpacer(),
		statusMsg,
	)
}



