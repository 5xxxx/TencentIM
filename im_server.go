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
	"github.com/5xxxx/TencentIM/tools/sign"
	"github.com/go-resty/resty/v2"
	"math/rand"
	"net/http"
	"strconv"
)

type IMServer struct {
	AppId      int
	Identifier string
	SecretKey  string
	Expire     int
	Sig        string
}

func NewIMServer(appId, expire int, identifier, secretKey string, opts ...ServerOption) (IMServer, error) {
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
	for _, opt := range opts {
		if err := opt.SetOption(&server); err != nil {
			return IMServer{}, err
		}
	}

	return server, nil
}

func (s IMServer) ListenCallback() {
	Run(strconv.Itoa(s.AppId))
}

func (s IMServer) request(url string, requestJson []byte) ([]byte, error) {
	body := bytes.NewBuffer(requestJson)
	client := resty.New()
	var respCheck struct {
		ErrorInfo string `json:"ErrorInfo"`
		ErrorCode int    `json:"ErrorCode"`
	}
	response, err := client.R().
		SetHeader("Content-Type", "application/json; charset=utf-8").
		SetResult(&respCheck).
		SetBody(body).Post(url)
	if err != nil {
		return nil, err
	}

	if response.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("http erro status code is %d", response.StatusCode())
	}

	if respCheck.ErrorCode != 0 {
		return nil, fmt.Errorf("操作失败，错误码 %d ,错误信息 %s \n 详情请查询 "+
			"https://cloud.tencent.com/document/product/269/1671", respCheck.ErrorCode, respCheck.ErrorInfo)
	}

	return response.Body(), nil

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
	return fmt.Sprintf("%s/%s%s?sdkappid=%d&identifier=%s&usersig=%s&random=%d&contenttype=json",
		BASE_URL, VERSION, path, s.AppId, s.Identifier, s.Sig, rand.Intn(4294967294))
}
