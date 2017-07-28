package qcloudsms

type QStatus struct {
	Result     int    `json:"result"`
	ErrMsg     string `json:"errmsg"`
	Mobile     string `json:"mobile"`
	NationCode string `json:"nationcode"`
	Sid        string `json:"sid"`
	Fee        int    `json:"fee"`
}

// QResponse to retrieve QCloud API Response.
type QResponse struct {
	Result int    `json:"result"`
	ErrMsg string `json:"errmsg"`
	Sid    string `json:"sid"`
	CallID string `json:"callid"`
	Fee    int    `json:"fee"`
	Ext    string `json:"ext"`

	Detail []QStatus   `json:"detail"`
	Data   interface{} `json:"data"`
}

func (qr *QResponse) Error() int {
	return qr.Result
}

func (qr *QResponse) Desc() string {
	return qr.ErrMsg
}

func (qr *QResponse) Status() (status []QSMSStatus) {
	if qr.Data == nil {
		return
	}
	status, ok := qr.Data.([]QSMSStatus)
	if !ok {
		return
	}
	return
}

func (qr *QResponse) Reply() (status []QSMSReply) {
	if qr.Data == nil {
		return
	}
	status, ok := qr.Data.([]QSMSReply)
	if !ok {
		return
	}
	return
}
