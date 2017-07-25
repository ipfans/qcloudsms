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

func (qr *QResponse) Status() (status []DeliveryStatus) {
	if qr.Data == nil {
		return
	}
	status, ok := qr.Data.([]DeliveryStatus)
	if !ok {
		return
	}
	return
}

type DeliveryStatus struct {
	UserReceiveTime string `json:"user_receive_time"`
	NationCode      string `json:"nationcode"`
	Mobile          string `json:"mobile"`
	ReportStatus    string `json:"report_status"`
	ErrMsg          string `json:"errmsg"`
	Description     string `json:"description"`
	Sid             string `json:"sid"`
}
