package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://yyds.thousandquant.com/api/v1/alert/getList"
	method := "POST"

	payload := strings.NewReader(`{}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjEzMjUzODQxNDRAcXEuY29tIiwiZXhwIjoxNzA4NTgzNzk5fQ.1HrFkFZpmikl-EJs1jvH_SRvZU4AGkfDdbl89oUxNGM")
	req.Header.Add("AppName", "sand.thousandquant.com")
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
