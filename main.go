package main

import (
	"archive/zip"
	"bufio"
	"encoding/xml"
	"fmt"
	"log"
	"strings"
)

func main() {
	fmt.Print("Hello")

	sheet, err := openFirstSheet("test.xlsx")
	if err != nil {
		log.Fatalf("openFirstSheet failed: %v", err)
	}

	readRows(sheet)

	fmt.Print("Bye")
}

func openFirstSheet(src string) (*zip.File, error) {
	zipReader, err := zip.OpenReader(src)
	if err != nil {
		return nil, err
	}
	defer zipReader.Close()

	for _, f := range zipReader.File {
		if strings.HasSuffix(f.Name, "/sheet1.xml") {
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

	_ = xml.NewDecoder(reader)
	_ = bufio.NewReader(reader)

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
