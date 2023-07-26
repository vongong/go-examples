package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	url := "https://tng-file-transfer.ckp-dev.centene.com/fileTransfer?folder=OHIO"
	method := "POST"
	filetype := "Test.txt"

	//test 1
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, err := os.Open(filetype)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(part, file)
	writer.Close()

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("Content-Type", writer.FormDataContentType())
	req.Header.Set("Content-Type", writer.FormDataContentType())
	// req.Header.Set("Content-Type", "multipart/form-data;")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	//defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ret value:", string(body))
	fmt.Println("ret statusCode:", res.StatusCode)

	// test2
	payload = &bytes.Buffer{}
	writer = multipart.NewWriter(payload)
	filename := "test2.txt"
	data := []byte("test from thing")
	part, err = writer.CreateFormFile("file", filename)
	if err != nil {
		log.Fatal(err)
	}
	_, err = part.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	err = writer.Close()
	if err != nil {
		log.Fatal(err)
	}

	client = &http.Client{}
	req, err = http.NewRequest(method, url, payload)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ret value:", string(body))
	fmt.Println("ret statusCode:", res.StatusCode)

}
