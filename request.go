package qcloudsms

// QRequest is request content to QCloud API.
type QRequest struct {
	Path       string      `json:"-"`
	Type       int         `json:"type"`
	Tel        interface{} `json:"tel"`
	Msg        string      `json:"msg"`
	Extend     string      `json:"extend"`
	Ext        string      `json:"ext"`
	Max        int         `json:"max"`
	BeginTime  int64       `json:"begin_time"`
	EndTime    int64       `json:"end_time"`
	NationCode string      `json:"nationcode"`
	Mobile     string      `json:"mobile"`

	// Generation form client.
	Sig  string `json:"sig"`
	Time int64  `json:"time"`
}

// QSMSReply retrieve SMS reply event.
type QSMSReply struct {
	NationCode string `json:"nationcode"`
	Mobile     string `json:"mobile"`
	Text       string `json:"text"`
	Sign       string `json:"sign"`
	Time       int    `json:"time"`
	Extend     string `json:"extend"`
}

// QSMSStatus retrieve SMS send event.
type QSMSStatus struct {
	UserReceiveTime string `json:"user_receive_time"`
	NationCode      string `json:"nationcode"`
	Mobile          string `json:"mobile"`
	ReportStatus    string `json:"report_status"`
	ErrMsg          string `json:"errmsg"`
	Description     string `json:"description"`
	Sid             string `json:"sid"`
}
