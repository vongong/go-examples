package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func fnameNoExt(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
		//if strings.ToLower(filepath.Ext(file.Name())) == "xml" {}
	}

	oldLocation := "/var/www/html/test.txt"
	newLocation := "/var/www/html/src/test.txt"
	err = os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fnameNoExt("base.bat"))
	fmt.Println(fnameNoExt("foo.bar"))
	fmt.Println(time.Now())
	fmt.Println(time.Now().Format("20060102_150405")) //yyymmdd_hh24miss

}
