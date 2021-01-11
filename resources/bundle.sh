#!/bin/bash
fyne bundle -name Arc42LogoPNG -package resources arc42-logo.png  > bundled.go
fyne bundle -name PDFnmbrrlogoPNG -append PDFnmbrr-logo.png  >> bundled.go
# fyne bundle -append PDFnmbrr-splash.png >> bundled.go