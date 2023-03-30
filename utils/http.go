package utils

import (
	"Bestdori-Proxy/errors"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"time"
)

func HttpGet(url string, object interface{}) (err error) {
	var Client = http.Client{
		Timeout: time.Second * 10, // 5秒超时
	}

	res, err := Client.Get(url)
	if err != nil {
		if url2Error, ok := err.(*url2.Error); ok && url2Error.Timeout() {
			return errors.RemoteReplyTimeout
		}
		return errors.RemoteReplyErr
	}
	if res.StatusCode != http.StatusOK {
		return errors.RemoteReplyReject
	}
	defer func(Body io.ReadCloser) {
		BErr := Body.Close()
		if BErr != nil {
			err = errors.RemoteReplyReadErr
		}
	}(res.Body)
	if res.StatusCode == http.StatusOK {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return errors.RemoteReplyReadErr
		}
		err = json.Unmarshal(body, object)
		if err != nil {
			return errors.RemoteReplyParseErr
		}
	}
	return nil
}

func HttpPost(url string, payload interface{}, object interface{}) (err error) {
	var Client = http.Client{
		Timeout: time.Second * 10, // 5秒超时
	}

	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return errors.RemoteRequestParseErr
	}

	res, err := Client.Post(url, "application/json", bytes.NewReader(payloadJson))
	if err != nil {
		if url2Error, ok := err.(*url2.Error); ok && url2Error.Timeout() {
			return errors.RemoteReplyTimeout
		}
		return errors.RemoteReplyErr
	}
	if res.StatusCode != http.StatusOK {
		return errors.RemoteReplyReject
	}
	defer func(Body io.ReadCloser) {
		BErr := Body.Close()
		if BErr != nil {
			err = errors.RemoteReplyReadErr
		}
	}(res.Body)
	if res.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return errors.RemoteReplyReadErr
		}
		err = json.Unmarshal(body, object)
		if err != nil {
			return errors.RemoteReplyParseErr
		}
	}
	return nil
}
