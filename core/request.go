package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func dealJsonPayload(data interface{}) io.Reader {
	if data != nil {
		var payload *bytes.Reader
		if c, err := json.Marshal(data); err == nil {
			payload = bytes.NewReader(c)
		}
		return payload
	} else {
		return nil
	}
}

func dealFormPayload(data interface{}) io.Reader {
	if data != nil {
		var payload *strings.Reader

		mapData, _ := json.Marshal(&data)
		m := make(map[string]string)
		json.Unmarshal(mapData, &m)
		form := url.Values{}
		for k, v := range m {
			form.Add(k, v)
		}
		payload = strings.NewReader(form.Encode())
		return payload
	} else {
		return nil
	}
}

func dealGetPayload(r *http.Request, data interface{}) {
	if data != nil {
		mapData, _ := json.Marshal(&data)
		m := make(map[string]string)
		json.Unmarshal(mapData, &m)
		q := r.URL.Query()
		for k, v := range m {
			q.Add(k, v)
		}
		r.URL.RawQuery = q.Encode()
	}
}

//ReqGet get请求
func ReqGet(url string, data interface{}, header interface{}) interface{} { // 创建 http 客户端
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	dealGetPayload(req, data)

	var headers = make(map[string]string)
	headers, ok := header.(map[string]string)
	if ok {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		// 处理错误
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// 处理错误
		fmt.Println(err)
	}
	//fmt.Println(string(body))
	m := Resp{}
	if err := json.Unmarshal(body, &m); err != nil {
		// 处理错误
		return string(body)
	}
	return m
}

//ReqPostForm form格式的post请求
func ReqPostForm(url string, data interface{}, header interface{}) interface{} { // 创建 http 客户端
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, dealFormPayload(data))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	var headers = make(map[string]string)
	headers, ok := header.(map[string]string)
	if ok {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		// 处理错误
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// 处理错误
		fmt.Println(err)
	}
	//fmt.Println(string(body))
	m := Resp{}
	if err := json.Unmarshal(body, &m); err != nil {
		// 处理错误
		return string(body)
	}
	return m
}

//ReqPostJson json格式的post请求
func ReqPostJson(url string, data interface{}, header interface{}) interface{} {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, dealJsonPayload(data))

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Content-Type", "application/json")
	var headers = make(map[string]string)
	headers, ok := header.(map[string]string)
	if ok {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		// 处理错误
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// 处理错误
		fmt.Println(err)
	}
	//fmt.Println(string(body))
	m := Resp{}
	if err := json.Unmarshal(body, &m); err != nil {
		// 处理错误
		return string(body)
	}
	return m
}
