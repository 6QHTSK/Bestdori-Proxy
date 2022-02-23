package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func HttpGet(url string, object interface{}) (errorCode int, err error) {
	var Client = http.Client{
		//Timeout: time.Second * 10, // 5秒超时
	}

	res, err := Client.Get(url)
	if err != nil {
		return http.StatusBadGateway, err
	}
	if res.StatusCode != http.StatusOK {
		return http.StatusNotFound, fmt.Errorf("未找到资源")
	}
	defer func(Body io.ReadCloser) {
		BErr := Body.Close()
		if BErr != nil {
			err = BErr
		}
	}(res.Body)
	if res.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		err = json.Unmarshal(body, object)
	}
	return http.StatusOK, nil
}

func HttpPost(url string, payload interface{}, object interface{}) (errorCode int, err error) {
	var Client = http.Client{
		//Timeout: time.Second * 10, // 5秒超时
	}

	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	res, err := Client.Post(url, "application/json", bytes.NewReader(payloadJson))
	if err != nil {
		return http.StatusBadGateway, err
	}
	if res.StatusCode != http.StatusOK {
		return http.StatusNotFound, fmt.Errorf("未找到资源")
	}
	defer func(Body io.ReadCloser) {
		BErr := Body.Close()
		if BErr != nil {
			err = BErr
		}
	}(res.Body)
	if res.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		err = json.Unmarshal(body, object)
	}
	return http.StatusOK, nil
}
