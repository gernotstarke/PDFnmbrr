package samplepdfs

import (
	"pdfcpuSamples/information"
	"testing"
)

// test if the example files exist
func TestSamplePDFs(t *testing.T) {
	valid, err := information.ValidatePDFFile(OnePageFile)

	if err != nil {
		t.Errorf("ValidatePDFFile: errof with file %v: %v", OnePageFile, err)
	}

	if valid == false {
		t.Errorf("ValidatePDFFile: %v not valid", OnePageFile)
	}

	// check if disguised markdown file exists
	disguisedExists, _ := information.FileExists(DisguisedMarkdownFile)

	if !disguisedExists {
		t.Errorf("File %v should exist but does not.", DisguisedMarkdownFile)
	}
}
