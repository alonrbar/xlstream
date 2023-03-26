package main

import (
	"archive/zip"
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Print("Hello")

	// Open file stream.
	filePath := "text.xlsx"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("os.Open(%q) failed: %v", filePath, err)
	}
	defer f.Close()

	// Get file size.
	fi, err := f.Stat()
	if err != nil {
		log.Fatalf("f.Stat() failed: %v", err)
	}

	// Open specific Excel sheet stream.
	sheetName := "sheet1"
	sheet, err := openSheet(f, sheetName, fi.Size())
	if err != nil {
		log.Fatalf("openSheet(%q) failed: %v", sheetName, err)
	}

	// Read rows
	readRows(sheet)

	fmt.Print("Bye")
}

func openSheet(r io.ReaderAt, sheetName string, zipSize int64) (*zip.File, error) {
	zr, err := zip.NewReader(r, zipSize)
	if err != nil {
		return nil, err
	}

	for _, f := range zr.File {
		if strings.HasSuffix(f.Name, fmt.Sprintf("/%v.xml", sheetName)) {
			return f, nil
		}
	}

	return nil, fmt.Errorf("first sheet not found")
}

func readRows(sheet *zip.File) {
	reader, err := sheet.Open()
	if err != nil {
		log.Fatalf("Failed to open sheet: %v", err)
	}
	defer reader.Close()

	// Two possible APIs for stream-reading an XML.
	_ = xml.NewDecoder(reader)
	_ = bufio.NewReader(reader)

	// TODO: Implement readRow using one of the above APIs.
	for {
		readRow()
	}
}

func readRow() {
	// find row start element

	// for
	// 	find next element
	// 	if start column
	// 		read value
	// 		continue
	//	if end row
	//		return
	//  if end of file
	//  	return
}
