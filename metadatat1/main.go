package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// A file metadata upload response from the Document API
type MetaUploadResponse struct {
	Id       MetaUploadId `json:"id" xml:"id"`
	Response string       `json:"response" xml:"response"`
}

//MetaUploadId
type MetaUploadId struct {
	ExternalMemberId   string `json:"externalMemberId" xml:"externalMemberId"`
	ExternalDocumentId string `json:"externalDocumentId" xml:"externalDocumentId"`
	// RequestDateTime is found in the response from the Document API only. It is not required for a request
	//RequestDateTime string `json:"requestdatetime,omitempty"`
}

func main() {
	inJson := []byte(`[
{
	"id": {
		"requestdatetime": "2016-10-11 08:44:46.706",
		"externalMemberId": "BW00010046201",
		"externalDocumentId": "EAA0"
		},
	"response": "Success"
	},
	{
	"id": {
		"requestdatetime": "2016-10-11 08:44:46.706",
		"externalMemberId": "BW00012466601",
		"externalDocumentId": "EAA1"
		},
	"response": "Success"
	},
	{
	"id": {
		"requestdatetime": "2016-10-11 08:44:46.706",
		"externalMemberId": "BW00012466601",
		"externalDocumentId": "EAA2"
		},
	"response": "Success"
	},
	{
	"id": {
		"requestdatetime": "2016-10-11 08:44:46.706",
		"externalMemberId": "BW00012466601",
		"externalDocumentId": "EAA3"
		},
	"response": "Success"
	},
	{
	"id": {
		"requestdatetime": "2016-10-11 08:44:46.706",
		"externalMemberId": "BW00012466601",
		"externalDocumentId": "EAA4"
		},
	"response": "Success"
	},
	{
	"id": {
		"requestdatetime": "2016-10-11 08:44:46.706",
		"externalMemberId": "BW00012466601",
		"externalDocumentId": "EAA5"
		},
	"response": "Success"
	},
	{
	"id": {
		"requestdatetime": "2016-10-11 08:44:46.706",
		"externalMemberId": "BW00012466601",
		"externalDocumentId": "EAA6"
		},
	"response": "Success"
	},
	{
	"id": {
		"requestdatetime": "2016-10-11 08:44:46.706",
		"externalMemberId": "BW00012466601",
		"externalDocumentId": "EAA7"
		},
	"response": "Success"
	},
	{
	"id": {
		"requestdatetime": "2016-10-11 08:44:46.706",
		"externalMemberId": "BW00012466601",
		"externalDocumentId": "EAA8"
		},
	"response": "Success"
	}
]`)

	inJson2 := []byte(`[
{
	"id": {
		"requestdatetime": "2016-10-11 08:44:46.706",
		"externalMemberId": "BW00010046201",
		"externalDocumentId": "EAA0"
		},
	"response": "Success"
	}]`)

	//Test 1
	fmt.Println("Test1:")
	resp := make([]MetaUploadResponse, 0)
	err := json.Unmarshal(inJson, &resp)
	if err != nil {
		fmt.Println("error:", err)
	}
	//fmt.Printf("Unmarshall Json: %+v\n", resp)

	//output, err := xml.MarshalIndent(&resp, "  ", "    ")
	output, err := xml.Marshal(&resp)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("Marshall Xml: %+v\n", string(output))

	//Test 2 Merge response
	fmt.Println("Test2:")
	respComb := make([]MetaUploadResponse, 0)

	for i := 0; i < 5; i++ {
		resp = make([]MetaUploadResponse, 0)
		err = json.Unmarshal(inJson2, &resp)
		if err != nil {
			fmt.Println("error:", err)
		}
		respComb = append(respComb, resp...)
	}

	fmt.Printf("combine Json: %+v\n", respComb)

	//Test 3
	resp = make([]MetaUploadResponse, 0)
	err = json.Unmarshal(inJson, &resp)
	if err != nil {
		fmt.Println("error:", err)
	}

	//count := 0
	n := 2
	respMini := make([]MetaUploadResponse, 0)
	for _, v := range resp {
		respMini = append(respMini, v)
		if len(respMini) >= n {
			outMini, err := xml.Marshal(&respMini)
			if err != nil {
				fmt.Printf("error: %v\n", err)
			}
			fmt.Printf("send: %+v\n\n", string(outMini))
			respMini = make([]MetaUploadResponse, 0)
		}
	}
	if len(respMini) > 0 {
		outMini, err := xml.Marshal(&respMini)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		fmt.Printf("send: %+v\n\n", string(outMini))
	}

}
