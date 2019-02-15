// Copyright (c) 2019 Flyu, Inc.
//
// Created by flyu on 2019/01/16.
//
// HTTP 返回结果

package response

import (
	"encoding/json"
	"net/http"
)

const (
	HTTP_SUCCESS = 0
	HTTP_FAILED  = 1
)

/*
SUCCESS:
{
    "code": 0,
    "data": {
        "age": 17,
        "name": "flyu"
    },
    "message": "success"
}
FAILED:
{
    "code": 1,
    "message": "failed"
}
*/
func HTTPJson(w http.ResponseWriter, code int, msg string, data ...interface{}) {
	resp := struct {
		Code    int         `json:"code"`
		Data    interface{} `json:"data,omitempty"`
		Message string      `json:"message"`
	}{}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if code == HTTP_SUCCESS {
		msg = "success"
	}
	resp.Code = code
	resp.Message = msg

	if len(data) == 1 {
		resp.Data = data[0]
	}

	rst, err := json.Marshal(resp)
	if err != nil {
		resp.Code = HTTP_FAILED
		resp.Data = nil
		resp.Message = "parse result failed, err: " + err.Error()
	}
	w.Write(rst)
}

/*
SUCCESS:
{
    "code": 0,
    "data": {
        "age": 17,
        "name": "flyu"
    },
    "message": "success",
    "pagination": {
        "page": 1,
        "page_size": 10,
        "total": 1000
    }
}
FAILED:
{
    "code": 1,
    "message": "failed"
}
*/
func HTTPJsonWithPage(w http.ResponseWriter, code int, msg string, page, pageSize, total int, data ...interface{}) {
	resp := struct {
		Code       int         `json:"code"`
		Data       interface{} `json:"data,omitempty"`
		Message    string      `json:"message"`
		Pagination struct {
			Page     int `json:"page"`
			PageSize int `json:"page_size"`
			Total    int `json:"total"`
		} `json:"pagination"`
	}{}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if code == HTTP_SUCCESS {
		msg = "success"
	}
	resp.Code = code
	resp.Message = msg
	if len(data) == 1 {
		resp.Data = data[0]
	}
	resp.Pagination.Page = page
	resp.Pagination.PageSize = pageSize
	resp.Pagination.Total = total

	rst, err := json.Marshal(resp)
	if err != nil {
		resp.Code = HTTP_FAILED
		resp.Data = nil
		resp.Message = "parse result failed, err: " + err.Error()
	}
	w.Write(rst)
}
