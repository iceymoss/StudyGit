package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func httpPost() {

	config := map[string]interface{}{}
	config["Id"] = 123
	fmt.Println(config)

	//JSON序列化
	configData, _ := json.Marshal(config)
	_ = bytes.NewBuffer([]byte(configData))

	resp, err := http.Post("http://localhost:8080/v1/auth/login",
		"application/json",
		strings.NewReader(string(configData)))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))

}

func main() {
	httpPost()

}
