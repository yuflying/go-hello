package extension

import (
	"bytes"
	"embrace/constant"
	"embrace/entity"
	"encoding/json"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func GetRealIP(r *http.Request) string {
	// todo: return x-forwarded-for for cdn
	if ip, _, e := net.SplitHostPort(r.RemoteAddr); e == nil {
		return ip
	}
	return r.RemoteAddr
}

func GetHeader(r *http.Request) entity.Header {
	return entity.Header{
		AccountID:    1,
		AccountName:  "flyu",
		IP:           GetRealIP(r),
		XTrace:       "", // 日志追踪
		Token:        "", // Token
		EncryptToken: "", //
		Refer:        "",
	}
}

func GetReq(r *http.Request) entity.Req {
	req := entity.Req{
		Source:     0,
		UserID:     0,
		UserName:   "",
		AccessTime: time.Now().Unix(),
		Header:     GetHeader(r),
	}
	req.UserID = req.Header.AccountID
	req.UserName = req.Header.AccountName
	return req
}

func ResponseJson(w http.ResponseWriter, code int, msg string, data ...interface{}) {
	var resp entity.Response
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if code == constant.CODE_SUCCESS {
		msg = "success"
	}
	resp.Code = code
	resp.Message = msg

	if len(data) == 1 {
		resp.Data = data[0]
	} else if len(data) == 2 {
		resp.Data = data[0]
		if v, ok := data[1].(entity.Pagination); ok {
			resp.Pagination = &v
		}
	}

	rst, _ := json.Marshal(resp)
	w.Write(rst)
}

func Get(url string, value interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, &value)
}

func Post(url string, header map[string]string, params interface{}, value interface{}) error {

	data, err := json.Marshal(params)
	resp, err := post(url, header, bytes.NewReader(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}
	if err = json.Unmarshal(body, value); err != nil {
		return err
	}
	return err
}

func post(url string, header map[string]string, body io.Reader) (resp *http.Response, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	return client.Do(req)
}
