package qcloudsms

// QResponse to retrieve QCloud API Response.
type QResponse struct {
	Result int    `json:"result"`
	ErrMsg string `json:"errmsg"`
	Ext    string `json:"ext"`
	Sid    string `json:"sid"`
	Fee    int    `json:"fee"`
}

func (qr *QResponse) Error() int {
	return qr.Result
}

func (qr *QResponse) Desc() string {
	return qr.ErrMsg
}
