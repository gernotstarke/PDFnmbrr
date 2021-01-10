package ui

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"net/url"
	"pdfcpuSamples/domain"
)


// Appl exposes the fyne application - mainly to enable the quit-function to stop the app.
var Appl fyne.App


// CreateMainUI creates and shows the main graphical user interface.
// It creates by delegating to "Panel" functions which will create their respective panel.
func CreateMainUI() {

	Appl = app.New()
	Appl.Settings().SetTheme(theme.LightTheme())
	w := Appl.NewWindow(domain.AppName )

	container := fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		logoHeaderPanel(),
		directoriesPanel(),
		configurationPanel(),
		widget.NewSeparator(),
		okCancelPanel(),
		widget.NewSeparator(),
		statusLine("no source directory selected"))

	w.SetContent(container)
	//w.Resize(fyne.NewSize(600, 400))
	w.SetFixedSize(true)
	w.ShowAndRun()
	w.CenterOnScreen()
}


func parseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	return link
}

func logoHeaderPanel() *fyne.Container {
	arc42Logo := canvas.NewImageFromFile("resources/arc42-logo.png")
	arc42Logo.FillMode = canvas.ImageFillContain
	arc42Logo.SetMinSize(fyne.NewSize(80, 40))
	arc42Logo.Resize(fyne.NewSize(80, 40))

	arc42Link := widget.NewHyperlinkWithStyle("arc42.org",
		parseURL("https://arc42.org"), fyne.TextAlignLeading, fyne.TextStyle{Bold: false})

	appLogo := canvas.NewImageFromFile("resources/PDFnumbrr-logo.png")
	appLogo.FillMode = canvas.ImageFillContain
	appLogo.SetMinSize(fyne.NewSize(200, 120))
	appLogo.Resize(fyne.NewSize(200, 120))

	container := fyne.NewContainerWithLayout(
		layout.NewHBoxLayout(),
		fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
			arc42Logo,
			layout.NewSpacer(),
			arc42Link),
		layout.NewSpacer(),
		appLogo,
	)
	return container
}

func directoriesPanel() fyne.CanvasObject {

	dirContainer := fyne.NewContainerWithLayout( layout.NewVBoxLayout(),
		srcDirSelectorGroup(),
		targetDirSelectorGroup())
	dirPanel :=  widget.NewCard("", "Directories", dirContainer)

	return dirPanel
}

func srcDirSelectorGroup() *fyne.Container {

	//fmt.Println(dir)
	srcDirField := widget.NewLabel("/Users/gernotstarke")

	srcDirButton := widget.NewButton("Source", func() {
		srcDirField.SetText(domain.SourceDir)
	})

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
	targetDirField := widget.NewLabel("/Users/gernotstarke/_target")

	targetDirButton := widget.NewButton("Target", func() {
		targetDirField.SetText( fmt.Sprintf("%v", domain.TargetDir))
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
//
// ==== Configuration Panel ========
//
func configurationPanel() fyne.CanvasObject {
	configContainer := fyne.NewContainerWithLayout( layout.NewVBoxLayout(),
		evenifyConfigGroup(),
		headerConfigGroup(),
		pageConfigGroup(),
		)

	return widget.NewCard("","Configuration", configContainer)
}

func evenifyConfigGroup() *fyne.Container {
	evenifyCheckbox := widget.NewCheck("Evenify?", func(bool) {})
	evenifyText := widget.NewEntry()
	evenifyText.SetText("Diese Seite bleibt absichtlich frei")

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

	headingLabel := widget.NewLabel("Header text: ")

	headingEntry := widget.NewEntry()
	headingEntry.SetPlaceHolder ("This text will be placed in the header of each page")

	return fyne.NewContainerWithLayout( layout.NewHBoxLayout(),
		headingLabel,
		headingEntry)
}


func pageConfigGroup() *fyne.Container {
	pagePrefixEntry := widget.NewEntry()
	pagePrefixEntry.SetText( "Seite")

	pageNrPositionSelect := widget.NewSelect([]string{"outside", "inside", "center 3"},
	     func(s string) { fmt.Println("selected", s) })

	return fyne.NewContainerWithLayout( layout.NewHBoxLayout(),
		widget.NewLabel( "Page prefix:"),
		pagePrefixEntry,
		layout.NewSpacer(),
		pageNrPositionSelect)
}

func statusLine(msg string ) *fyne.Container {

	statusMsg := canvas.NewText(msg, DarkRedColor)
	statusMsg.TextSize = 9
	statusMsg.Alignment = fyne.TextAlignTrailing

	versionLabel := canvas.NewText("v." + domain.VersionStr, NavyColor)
	versionLabel.TextSize = 10
	versionLabel.Alignment = fyne.TextAlignCenter

	return fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		versionLabel,
		widget.NewSeparator(),
		layout.NewSpacer(),
		statusMsg,
	)
}

func okCancelPanel() *fyne.Container {

	OKButton := widget.NewButton("Process PDFs", func() {})
	OKButton.Disable()

	CancelButton := widget.NewButton("Cancel", quitApp)

	return fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		layout.NewSpacer(),
		CancelButton,
		OKButton)
}


func quitApp() {
	Appl.Quit()
}

