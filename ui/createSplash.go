package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/driver/desktop"
	"pdfcpuSamples/resources"
	"time"
)

func CreateAndDisplaySplash()  {

	if drv, ok := fyne.CurrentApp().Driver().(desktop.Driver); ok {
		splashWindow := drv.CreateSplashWindow()

		splash := canvas.NewImageFromResource(resources.PDFnmbrrSplash)
		splash.SetMinSize( fyne.NewSize(600, 540))


		splashWindow.SetContent(splash)
		splashWindow.Show()

		go func() {
			time.Sleep(time.Second * 3)
			splashWindow.Close()
		}()
	}
}
