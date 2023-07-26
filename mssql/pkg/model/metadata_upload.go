package model

// A file metadata upload request to the Document API
type MetaUploadRequest struct {
	Id             MetaUploadId `json:"id"`
	DocumentFormat string       `json:"documentFormat"`
	DocumentURL    string       `json:"documentUrl"`
	DocumentType   string       `json:"documentType"`
	FileName       string       `json:"filename"`
	FunctionalArea string       `json:"functionalArea"`
	ObsoleteFlag   string       `json:"obsoleteFlag"`
	ReceiptDate    string       `json:"receiptDate"`
	SubmittedBy    string       `json:"submittedBy"`
}

// A file metadata upload response from the Document API
type MetaUploadResponse struct {
	Id       MetaUploadId `json:"id"`
	Response string       `json:"response"`
}

//MetaUploadId
type MetaUploadId struct {
	ExternalMemberId   string `json:"externalMemberId"`
	ExternalDocumentId string `json:"externalDocumentId"`
	// RequestDateTime is found in the response from the Document API only. It is not required for a request
	RequestDateTime string `json:"requestdatetime,omitempty"`
}
