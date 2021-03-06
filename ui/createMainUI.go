package ui

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"net/url"
	"pdfcpuSamples/domain"
	"pdfcpuSamples/resources"
)

// Appl exposes the fyne application - mainly to enable the quit-function to stop the app.
var Appl fyne.App
var Window fyne.Window

// CreateMainUI creates and shows the main graphical user interface.
// It creates by delegating to "Panel" functions which will create their respective panel.
func CreateMainUI() {

	Appl = app.New()

	// CreateAndDisplaySplash()

	Appl.Settings().SetTheme(theme.LightTheme())
	Window = Appl.NewWindow(domain.AppName)

	container := fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		logoHeaderPanel(),
		directoriesPanel(),
		configurationPanel(),
		widget.NewSeparator(),
		okCancelPanel(),
		widget.NewSeparator(),
		statusLine("no source directory selected"))

	Window.SetContent(container)
	Window.Resize(fyne.NewSize(600, 400))
	Window.SetFixedSize(true)
	Window.CenterOnScreen()
	Window.ShowAndRun()
}

func parseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	return link
}

func logoHeaderPanel() *fyne.Container {
	arc42Logo := canvas.NewImageFromResource(resources.Arc42LogoPNG)
	arc42Logo.FillMode = canvas.ImageFillContain
	arc42Logo.SetMinSize(fyne.NewSize(80, 40))
	arc42Logo.Resize(fyne.NewSize(80, 40))

	/*
		arc42Link := widget.NewHyperlinkWithStyle("arc42.org",
			parseURL("https://arc42.org"), fyne.TextAlignLeading, fyne.TextStyle{Bold: false})
	*/
	appLogo := canvas.NewImageFromResource(resources.PDFnmbrrlogoPNG)
	appLogo.FillMode = canvas.ImageFillContain
	appLogo.SetMinSize(fyne.NewSize(200, 120))
	appLogo.Resize(fyne.NewSize(200, 120))

	container := fyne.NewContainerWithLayout(
		layout.NewHBoxLayout(),
		fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
			arc42Logo,
			layout.NewSpacer()),
		layout.NewSpacer(),
		appLogo,
	)
	return container
}

func directoriesPanel() fyne.CanvasObject {

	dirContainer := fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		srcDirSelectorGroup(),
		targetDirSelectorGroup())
	dirPanel := widget.NewCard("", "Directories", dirContainer)

	return dirPanel
}

func srcDirSelectorGroup() *fyne.Container {
	srcDirField := widget.NewEntry()
	srcDirField.SetText(domain.SourceDirName())



	srcDirButton := widget.NewButton("Source", func() {
		dialog.ShowFolderOpen(func(list fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, Window)
				return
			}
			if list == nil {
				return
			}

			_, err = list.List()
			if err != nil {
				dialog.ShowError(err, Window)
				return
			}
			srcDirField.SetText( list.Name())
			fmt.Printf("Folder %s :\n%s", list.Name(),  list.String())
			// dialog.ShowInformation("Folder Open", out, Window)
		}, Window)
	})
	srcDirButton.SetIcon( theme.FolderOpenIcon() )

	srcDirLabel := canvas.NewText("nothing selected", NavyColor)
	srcDirLabel.TextSize = 9

	return fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		srcDirButton,
		srcDirField,
		layout.NewSpacer(),
		srcDirLabel,
	)
}

func targetDirSelectorGroup() *fyne.Container {
	targetDirField := widget.NewEntry()
	targetDirField.SetText(domain.TargetDirName())


	targetDirButton := widget.NewButton("Target", func() {
		dialog.ShowFolderOpen(func(list fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, Window)
				return
			}
			if list == nil {
				return
			}

			_, err = list.List()
			if err != nil {
				dialog.ShowError(err, Window)
				return
			}
			targetDirField.SetText( list.Name())
			fmt.Printf("Folder %s :\n%s", list.Name(),  list.String())
			// dialog.ShowInformation("Folder Open", out, Window)
		}, Window)
	})
	targetDirButton.SetIcon( theme.FolderOpenIcon() )
	targetDirLabel := canvas.NewText("", NavyColor)
	targetDirLabel.TextSize = 9

	return fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		targetDirButton,
		targetDirField,
		layout.NewSpacer(),
		targetDirLabel,
	)
}

//
// ==== Configuration Panel ========
//
func configurationPanel() fyne.CanvasObject {
	configContainer := fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		evenifyConfigGroup(),
		headerConfigGroup(),
		pageConfigGroup(),
	)

	return widget.NewCard("", "Configuration", configContainer)
}

func evenifyConfigGroup() *fyne.Container {
	evenifyText := widget.NewEntry()
	evenifyText.Disable()
	evenifyText.SetText(domain.BlankPageText())

	evenifyCheckbox := widget.NewCheck("Evenify?", func(value bool) {
		if value == false {
			evenifyText.Disable()
		} else {
			evenifyText.Enable()
		}
	})

	concatenateCheckbox := widget.NewCheck("Concatenate?", func(bool) {})
	concatenateCheckbox.Disable()

	return fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		evenifyCheckbox,
		evenifyText,
		layout.NewSpacer(),
		concatenateCheckbox,
	)
}

func headerConfigGroup() *fyne.Container {

	headingEntry := widget.NewEntry()
	headingEntry.SetPlaceHolder(domain.HeaderText())

	headerCheckbox := widget.NewCheck("Header?", func(value bool) {
		if value == false {
			headingEntry.Disable()
		} else {
			headingEntry.Enable()
		}
	})

	return fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		headerCheckbox,
		headingEntry)
}

func pageConfigGroup() *fyne.Container {
	pagePrefixEntry := widget.NewEntry()
	pagePrefixEntry.SetText(domain.PageNumberPrefix())

	pageNrPositionSelect := widget.NewSelectEntry([]string{"outside", "inside", "center"})
	pageNrPositionSelect.SetText("outside")

	return fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		widget.NewLabel("Page prefix:"),
		pagePrefixEntry,
		layout.NewSpacer(),
		widget.NewLabel("Page nr position:"),
		pageNrPositionSelect)
}

func statusLine(msg string) *fyne.Container {

	statusMsg := canvas.NewText(msg, DarkRedColor)
	statusMsg.TextSize = 9
	statusMsg.Alignment = fyne.TextAlignTrailing

	versionLabel := canvas.NewText("v."+domain.VersionStr, NavyColor)
	versionLabel.TextSize = 10
	versionLabel.Alignment = fyne.TextAlignCenter

	return fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		versionLabel,
		widget.NewSeparator(),
		layout.NewSpacer(),
		statusMsg,
	)
}

func okCancelPanel() fyne.CanvasObject {

	OKButton := widget.NewButton("Process PDFs", func() {})
	OKButton.Disable()

	CancelButton := widget.NewButton("Cancel", quitApp)

	buttons := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		layout.NewSpacer(),
		CancelButton,
		OKButton)

	okCancelPanel := widget.NewCard("", "Processing", buttons)

	return okCancelPanel
}

func quitApp() {
	Appl.Quit()
}
