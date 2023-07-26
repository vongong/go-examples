package model

//MetaDataWC ...
type MetaDataWC struct {
	MemberId     string `json:"memberId"`
	DocumentId   string `json:"documentId"`
	DocumentURL  string `json:"documentUrl"`
	DocumentType string `json:"documentType"`
	FileName     string `json:"filename"`
	ReceiptDate  string `json:"receiptDate"`
	SubmittedBy  string `json:"submittedBy"`
}
