package qcloudsms

// QRequest to perpare request to QCloud API.
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
