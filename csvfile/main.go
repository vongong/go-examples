package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	data := [][]string{
		{"col1", "col2"},
		{"Line1", "Hello Readers of"},
		{"Line2", "golangcode.com"},
	}

	csvdata, err := makeCSV(data)
	if err != nil {
		log.Fatal("Error", err)
	}
	fmt.Println("data:", string(csvdata))

	ohPath := "csvdata"
	ohPre := "OH_"
	dateFormat := "2006_01_02"
	fileName := ohPre + time.Now().Format(dateFormat)

	filePath := filepath.Join(ohPath, getFilename(ohPath, fileName))
	fmt.Println("file path:", filePath)

	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal("Error", err)
	}
	defer file.Close()

	_, err = file.Write(csvdata)
	if err != nil {
		log.Fatal("Error", err)
	}
	fmt.Println("Write Written")
}

func makeCSV(csvData [][]string) ([]byte, error) {
	buf := &bytes.Buffer{}
	w := csv.NewWriter(buf)
	err := w.WriteAll(csvData)
	if err != nil {
		return nil, err
	}
	w.Flush()
	return buf.Bytes(), nil
}

func getFilename(fpath string, basename string) string {
	fileExt := ".csv"
	checkfile := filepath.Join(fpath, basename+fileExt)
	fmt.Println(checkfile)
	_, err := os.Stat(checkfile)
	if os.IsNotExist(err) {
		return basename + fileExt
	}

	var newname string
	for i := 1; i <= 1000000; i++ {
		newname = basename + "-" + strconv.Itoa(i)
		checkfile = filepath.Join(fpath, newname+fileExt)
		_, err = os.Stat(checkfile)
		if os.IsNotExist(err) {
			break
		}
		i++
	}
	return newname + fileExt
}
