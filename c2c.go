/*
 *
 * c2c.go
 * callback
 *
 * Created by lintao on 2020/6/1 5:16 下午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package TencentIM

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func BeforeSendMsg(req Req, c echo.Context) error {
	fmt.Println(req.CallbackCommand)
	return nil
}

func AfterSendMsg(req Req, c echo.Context) error {
	fmt.Println(req.CallbackCommand)
	return nil
}

type C2CMsg struct {
	CallbackCommand string `json:"CallbackCommand"`
	FromAccount     string `json:"From_Account"`
	ToAccount       string `json:"To_Account"`
	MsgSeq          int    `json:"MsgSeq"`
	MsgRandom       int    `json:"MsgRandom"`
	MsgTime         int    `json:"MsgTime"`
	MsgKey          string `json:"MsgKey"`
	MsgBody         []struct {
		MsgType    string `json:"MsgType"`
		MsgContent struct {
			Text string `json:"Text"`
		} `json:"MsgContent"`
	} `json:"MsgBody"`
}

type C2CCallBack func(req Req, info C2CMsg) error

func (s IMServer) CallBackBeforeSendMsg(commandFunc C2CCallBack) {
	AddCommandFuc(C2CBeforeSendMsg, func(req Req, c echo.Context) error {
		var info C2CMsg
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}

func (s IMServer) CallBackAfterSendMsg(commandFunc C2CCallBack) {
	AddCommandFuc(C2CAfterSendMsg, func(req Req, c echo.Context) error {
		var info C2CMsg
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}
