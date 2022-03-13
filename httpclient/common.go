package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Post(url string, jsonStr string) (respJson string) {

	byteStr := []byte(jsonStr)
	url = "http://" + url
	fmt.Println("333" + jsonStr)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(byteStr))
	fmt.Println(req)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	fmt.Println("1", req)
	if err != nil {
		fmt.Println("err!!!", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	respJson = string(body)
	fmt.Println(respJson)
	return
}
func JSONToMap(str string) map[string]interface{} {

	var tempMap = make(map[string]interface{})

	err := json.Unmarshal([]byte(str), &tempMap)

	if err != nil {
		panic(err)
	}

	return tempMap
}
