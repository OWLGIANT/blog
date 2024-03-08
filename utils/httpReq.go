package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
)

func HTTPRequest(method, url string, preload interface{}) (response []byte, err error) {
	jsonData, err := json.Marshal(preload)
	if err != nil {
		return
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return
	}
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	response, err = io.ReadAll(resp.Body)
	return
}
