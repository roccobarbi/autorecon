package network

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type JsonGetRequest[T any] struct {
	BaseUrl  string
	Query    map[string]string
	response []T
}

func (r *JsonGetRequest[T]) Request() []T {
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
			fmt.Println("response ReadAll failed.")
			log.Fatal(err)
		}
		err = json.Unmarshal(bodyBytes, &r.response)
		if err != nil {
			fmt.Println("json.Unmarshal failed.")
			log.Fatal(err)
		}
	}
	return r.response
}

func (r *JsonGetRequest[T]) SetBaseUrl(baseUrl string) {
	r.BaseUrl = baseUrl
}

func (r *JsonGetRequest[T]) SetQueryKeyValue(key string, value string) {
	if r.Query == nil {
		r.Query = make(map[string]string)
	}
	r.Query[key] = value
}
