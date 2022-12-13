package req

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/jeffcail/go-tron/utils"
)

var client *http.Client

func init() {
	def := http.DefaultTransport
	defPot, ok := def.(*http.Transport)
	if !ok {
		panic("init transport 出错")
	}
	defPot.MaxIdleConns = 100
	defPot.MaxIdleConnsPerHost = 100
	defPot.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client = &http.Client{
		Timeout:   time.Second * time.Duration(20),
		Transport: defPot,
	}
}

// Get
func Get(url string, header map[string]string, params map[string]interface{}) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if header != nil {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}

	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			v, _ := utils.ToString(val)
			q.Add(key, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	if r.StatusCode != 200 {
		return nil, err
	}

	bb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return bb, err
}

// Post
func Post(url string, header map[string]string, params map[string]interface{}) (rr []byte, e error) {
	defer func() {
		if rr == nil && e == nil {
			fmt.Println("")
		}
	}()
	dd, _ := json.Marshal(params)
	re := bytes.NewReader(dd)
	req, err := http.NewRequest("POST", url, re)
	if err != nil {
		return nil, err
	}
	if header != nil {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}
	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	if r.StatusCode != 200 {
		return nil, errors.New(r.Status)
	}

	bb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return bb, err
}
