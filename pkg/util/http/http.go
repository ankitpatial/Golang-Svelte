/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package http

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"roofix/pkg/util/log"
	"time"
)

type HttpReqExt func(req *http.Request)

func client() *http.Client {
	return &http.Client{
		Timeout: time.Minute * 15, // keeping in mind that max time for lambda execution
	}
}

// Get call an url with content-type: application/json
func Get(ep string, ext ...HttpReqExt) (data []byte, err error) {
	log.Info("GET %s", ep)
	req, err := http.NewRequest("GET", ep, nil)
	if err != nil {
		return nil, err
	}

	for _, e := range ext {
		e(req)
	}

	return send(req)
}

func Post(ep string, data interface{}, ext ...HttpReqExt) ([]byte, error) {
	log.Info("POST %s", ep)
	var req *http.Request
	var err error
	if data != nil {
		var body []byte
		switch data.(type) {
		case url.Values:
			body = []byte(data.(url.Values).Encode())
		default:
			body, _ = json.Marshal(data)
		}

		//fmt.Println("** POST BODY **")
		//fmt.Printf("%s\n", string(body))

		req, err = http.NewRequest("POST", ep, bytes.NewReader(body))
	} else {
		req, err = http.NewRequest("POST", ep, nil)
	}

	if err != nil {
		return nil, err
	}

	for _, e := range ext {
		e(req)
	}

	return send(req)
}

func send(req *http.Request) ([]byte, error) {
	req.Close = true
	res, err := client().Do(req)
	if err != nil {
		return nil, err
	}
	defer func(b io.ReadCloser) {
		_ = b.Close()
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= http.StatusMultipleChoices {
		msg := fmt.Sprintf("%s, %s", res.Status, string(body))
		return body, errors.New(msg)
	}

	return body, nil
}

// WithAuthBearer sets request header, Authorization: Bearer token...
func WithAuthBearer(token string) HttpReqExt {
	return func(req *http.Request) {
		req.Header.Add("Authorization", "Bearer "+token)
	}
}

// WithAuthBasic sets request header, Authorization: Basic base64(username:password)
func WithAuthBasic(username, password string) HttpReqExt {
	return func(req *http.Request) {
		auth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		req.Header.Add("Authorization", "Basic "+auth)
	}
}

// WithContentTypeForm sets request header, Content-Current: application/x-www-form-urlencoded
func WithContentTypeForm() HttpReqExt {
	return func(req *http.Request) {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
}

// WithContentTypeJson sets request header, Content-Current: application/json
func WithContentTypeJson() HttpReqExt {
	return func(req *http.Request) {
		req.Header.Set("Content-Type", "application/json")
	}
}

func WithHeader(key, value string) HttpReqExt {
	return func(req *http.Request) {
		req.Header.Set(key, value)
	}
}
