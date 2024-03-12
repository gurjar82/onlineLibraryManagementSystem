    package models

	type  Issueregistery struct {
		Issueid int `json:"issueid"`
		ISBN string `json:"isbn"`
		Readerid int `json:"readerid"`
		Issueapproverid string `json:"issueapproverid"`
		Issuestatus string `json:"issuestatus"`
		Issuedate string `json:"issuedate"`
		Expectedreturndate string `json:"expectedreturndate"`
		Returndate string `json:"returndate"`
		Returnapproverid string `json:"returnapproverid"` 
	}