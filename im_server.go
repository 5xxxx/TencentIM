/*
 *
 * im_server.go
 * server
 *
 * Created by lintao on 2020/4/16 2:20 下午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package TencentIM

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/NSObjects/TencentIM/tools/sign"
)

type IMServer struct {
	AppId      int
	Identifier string
	SecretKey  string
	Expire     int
	Sig        string
}

func NewIMServer(appId, expire int, identifier, secretKey string) (IMServer, error) {
	server := IMServer{
		AppId:      appId,
		Identifier: identifier,
		SecretKey:  secretKey,
		Expire:     expire,
	}
	var err error
	if server.Sig, err = server.userSig(); err != nil {
		return IMServer{}, err
	}

	return server, nil
}

func (s IMServer) request(url string, requestJson []byte) ([]byte, error) {
	body := bytes.NewBuffer(requestJson)
	// Create client
	fmt.Println(string(requestJson))
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("POST", url, body)

	// Headers
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http erro status code is %d", resp.StatusCode)
	}
	var respCheck struct {
		ErrorInfo string `json:"ErrorInfo"`
		ErrorCode int    `json:"ErrorCode"`
	}
	if err = json.Unmarshal(respBody, &respCheck); err != nil {
		return nil, err
	}

	if respCheck.ErrorCode != 0 {
		return nil, fmt.Errorf("操作失败，错误码 %d ,错误信息 %s \n 详情请查询 "+
			"https://cloud.tencent.com/document/product/269/1671", respCheck.ErrorCode, respCheck.ErrorInfo)
	}

	fmt.Println("response Body : ", string(respBody))
	return respBody, nil

}

func (s IMServer) requestWithPath(path IMPath, v interface{}) (jsons []byte, err error) {
	var b []byte
	switch a := v.(type) {
	case []byte:
		b = a
	default:
		b, err = json.Marshal(&v)
		if err != nil {
			return nil, err
		}
	}

	return s.request(s.combineURL(path), b)
}

func (s IMServer) userSig() (string, error) {
	userSig, err := sign.GenSig(s.AppId, s.SecretKey, s.Identifier, s.Expire)
	if err != nil {
		return "", err
	}
	return userSig, nil
}

func (s IMServer) combineURL(path IMPath) string {
	rand.Seed(time.Now().Unix())

	return fmt.Sprintf("%s/%s%s?sdkappid=%d&identifier=%s&usersig=%s&random=%d&contenttype=json",
		BASE_URL, VERSION, path, s.AppId, s.Identifier, s.Sig, rand.Intn(4294967294))
}
