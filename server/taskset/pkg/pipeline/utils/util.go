package utils

import (
	"bytes"
	"encoding/gob"
	"math/rand"
	"time"
	"io/ioutil"
	"net/http"
	"strings"
	"net/url"
)

const (
	_RANDOM_SEEK string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func DeepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

func GetRandomString(l int) string {
	bytes := []byte(_RANDOM_SEEK)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func DoRequest(method, url, body string, headers map[string]string) ([]byte, int, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, 0, err
	}
	if "" != body {
		req.Body = ioutil.NopCloser(strings.NewReader(body))
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 30 * time.Second}
	var result []byte
	var rsp *http.Response

	rsp, err = client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer rsp.Body.Close()
	result, err = ioutil.ReadAll(rsp.Body)
	return result, rsp.StatusCode, err
}

func PostForm(addr string, form map[string]string) ([]byte, int, error) {

	value := url.Values{}

	if form != nil {
		for k, v := range form {
			value.Add(k, v)
		}
	}

	rsp, err := http.PostForm(addr, value)

	if err != nil {
		return nil, 0, err
	}

	defer rsp.Body.Close()

	result, err := ioutil.ReadAll(rsp.Body)
	
	if err != nil {
		return nil, 0, err
	}

	return result, rsp.StatusCode, err

}