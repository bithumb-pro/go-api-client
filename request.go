package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"bytes"
)

var contentType = "application/json"

func get(url string, r interface{}) error {
	resp := doGet(url)
	nil := doParse(resp, r)
	return nil
}

func post(url string, params interface{}, r interface{}) error {
	jsonBytes, err := json.Marshal(params)
	if err != nil {
		return err
	}
	resp := doPost(url, jsonBytes)
	nil := doParse(resp, r)
	return nil
}

func doParse(resp []byte, in interface{}) error {
	err := json.Unmarshal(resp, in)
	if err != nil {
		return err
	}
	return nil
}

func doGet(url string) []byte {
	resp, err := http.Get(url)
	return handleResp(resp, err)
}

func doPost(url string, data []byte) []byte {
	body := bytes.NewReader(data)
	resp, err := http.Post(url, contentType, body)
	return handleResp(resp, err)
}

func handleResp(resp *http.Response, err error, ) []byte {
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return r
}
