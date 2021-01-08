# PDF Numberer (pdfnmbrr)

Adding page numbers to existing PDFs, and some more...

## Statistics

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0) 
[![Go Report Card](https://goreportcard.com/badge/github.com/gernotstarke/pdfnmbrr)](https://goreportcard.com/report/github.com/gernotstarke/pdfnmbrr)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=gernotstarke_pdfnmbrr&metric=alert_status)](https://sonarcloud.io/dashboard?id=gernotstarke_pdfnmbrr)
[![Maintainability](https://api.codeclimate.com/v1/badges/c481ef8142826f71ff65/maintainability)](https://codeclimate.com/github/gernotstarke/pdfnmbrr/maintainability)
[![golangci-lint](https://github.com/gernotstarke/pdfnmbrr/workflows/golangci-lint/badge.svg)](https://github.com/gernotstarke/pdfnmbrr)

##### The following won't currently work as some tests require filesystem access
![Go](https://github.com/gernotstarke/pdfnmbrr/workflows/Go/badge.svg) 
[![Coverage Status](https://coveralls.io/repos/github/gernotstarke/pdfnmbrr/badge.svg?branch=main)](https://coveralls.io/github/gernotstarke/pdfnmbrr?branch=main) 


## Intro, Requirements and Some History

### The Problem and History
We create PDF files from various sources (ppt, odt, doc, keynote, html...) and want to create nice printable handouts from these sources.

With these various document sources, it's neither possible to create consistent pagenumbers, nor to add custom header information.

![the problem](./images/AToPDF_the_Problem.png)

Instead of flushing down many many bucks into Adob*s hungry mouth I decided (way back in 2013) to create my own solution.
Back then I used the https://griffon-framework.org[Griffon Framework] for the first version, plus the awesome https://github.com/itext/itextpdf[itext-pdf] library for handling pdfs.
All was based upon the (then current) Java version 6.0.

But times are changing: Java 6 isn't universally available ony longer, Griffon changed direction and hasn't become as mainstream as I hoped for.
The original app refuses to start any longer...


### Preconditions

1. The *content* of the files has been converted to pdf (so we're not concerned about the original data or file formats nor the tools to create/manipulate the original content).
2. The dimensions (pagesize) of the pdf files identical (we assume A4-portrait in most cases).
3. The existing pdf files are *not* encrypted.

4. We don't want to make any assumptions about user infrastructure, so we don't want the user to install things like Java...

### Use Cases

Ooops - the following use case diagram looks like a big app - which our little pdf numberer is not :-)

![use cases](./images/AToPdf_use_cases.jpg)

...

## Thanx and Credits

* Thanx to Horst Rutter, creator and maintainer of the [pdfcpu](https://pdfcpu.io) PDF library
* Thanx to Andrew Williams, creator and core maintainer of the [fyne.io](https://fyne.io) cross-platform ui library
   