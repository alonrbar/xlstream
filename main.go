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
	// Note: To create a ReaderAt interface for reading from GCS see: 
	// https://github.com/carbocation/genomisc/compare/8c9ef122d6ab47285dd12288bd9adbad1e735edd...4611e6df9ce70ee9967d19f580ce147ab96106c4
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
	//
	// Excel worksheet XML structure is as follows:
	//
	// <worksheet ...>
	//   ...
	//   <sheetData>
        //     <row r="1">
        //       <c r="A1" s="1" t="s">
        //         <v>0</v>
        //         ...
        //       </c>
        //     </row>
        //     ...
        //   </sheetData>
	//   ...
	//  </worksheet>
	//
	// Read more at http://officeopenxml.com/SScontentOverview.php
	//
	
	//
	// Pseudo code:
	//
	// find row start element
	//
	// for
	// 	find next element
	// 	if start column
	// 		read value
	// 		continue
	//	if end row
	//		return	
	//  	if end of file
	//  		return
	//
}
