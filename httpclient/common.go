package httpclient

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func Post(jsonStr string)(respJson string) {
	url := "127.0.0.1"
	byteStr := []byte(jsonStr)
	req, _ := http.NewRequest("post",url,bytes.NewBuffer(byteStr))
	req.Header.Set("Content-Type","application/json")
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	respJson = string(body)
	return
}