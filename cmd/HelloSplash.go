package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/driver/desktop"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"pdfcpuSamples/resources"
	"time"
)

func main() {
	a := app.New()

	createAndDisplaySplash()

	w := a.NewWindow("Hello Splash")

	time.Sleep( time.Millisecond * 3000)
	container := fyne.NewContainerWithLayout(
		layout.NewHBoxLayout(),
		widget.NewLabel("one label"),
		widget.NewSeparator(),
		widget.NewLabel("another label"))

	w.SetContent(container)
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(600, 200))
	//w.ShowAndRun()
}

func createAndDisplaySplash()  {
	if drv, ok := fyne.CurrentApp().Driver().(desktop.Driver); ok {
		splashWindow := drv.CreateSplashWindow()

//		splash := canvas.NewImageFromResource(resources.PDFnmbrrSplash)
		//splash.Resize(fyne.NewSize(600, 543))
		//splash.FillMode = canvas.ImageFillContain


		splashWindow.SetContent(
			fyne.NewContainerWithLayout(layout.NewGridWrapLayout(fyne.NewSize(500, 440)),
				canvas.NewImageFromResource( resources.PDFnmbrrSplash)) )

		splashWindow.Show()

		go func() {
			time.Sleep(time.Second * 3)
			splashWindow.Close()
		}()
	}
}
