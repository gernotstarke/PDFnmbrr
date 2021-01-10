package main

import (
	"fyne.io/fyne"
	"pdfcpuSamples/domain"
	"pdfcpuSamples/ui"
)

var appl fyne.App



func main() {

	domain.SetupInitialState()
	ui.CreateMainUI()

}
