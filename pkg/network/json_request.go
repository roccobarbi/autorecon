package network

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type JsonGetRequest struct {
	BaseUrl  string
	Query    map[string]string
	response []byte
}

func (r *JsonGetRequest) Request() []byte {
	if r.response != nil {
		return r.response // cached from previous use
	}
	var sb strings.Builder
	sb.WriteString(r.BaseUrl)
	index := 0
	for key, value := range r.Query {
		if index == 0 {
			sb.WriteString("?")
		} else {
			sb.WriteString("&")
		}
		sb.WriteString(key)
		sb.WriteString("=")
		sb.WriteString(value)
		index += 1
	}
	resp, err := http.Get(sb.String())
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)
	if resp.StatusCode == 200 {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		r.response = bodyBytes
	}
	return r.response
}

func (r *JsonGetRequest) SetBaseUrl(baseUrl string) {
	r.BaseUrl = baseUrl
}

func (r *JsonGetRequest) SetQueryKeyValue(key string, value string) {
	if r.Query == nil {
		r.Query = make(map[string]string)
	}
	r.Query[key] = value
}
