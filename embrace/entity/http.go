package entity

// 请求头部信息
type Header struct {
	XTrace       string `json:"X-Trace"`    // 日志追踪
	AccountID    int64  `json:"Account-ID"` // 用户ID
	AccountName  string `json:"account_name"`
	IP           string `json:"ip"`
	Token        string `json:"Token"`         // Token
	EncryptToken string `json:"Encrypt-Token"` //
	Refer        string `json:"refer"`
}

// 请求主体
type Req struct {
	Source     int    `json:"source"`
	UserID     int64  `json:"user_id"`
	UserName   string `json:"user_name"`
	AccessTime int64  `json:"access_time"`
	Header
}
