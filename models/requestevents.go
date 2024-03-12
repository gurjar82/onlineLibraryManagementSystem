package models

type Requestevents struct {
	Reqid int `json:"reqid"`
	Bookid string `json:"bookid"`
	Readerid int `json:"reader"`
	Requestdate string `json:"requestdate"`
	Approvaldate string `json:"approvaldate"`
	Approverid int `json:"approverid"`
	Requesttype string `json:"requesttype"`
}