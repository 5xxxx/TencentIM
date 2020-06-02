/*
 *
 * sns.go
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

func FriendAdd(req Req, c echo.Context) error {
	fmt.Println(req.CallbackCommand)
	return nil
}

func FriendDelete(req Req, c echo.Context) error {
	fmt.Println(req.CallbackCommand)
	return nil
}

func BlackListAdd(req Req, c echo.Context) error {
	fmt.Println(req.CallbackCommand)
	return nil
}

func BlackListDelete(req Req, c echo.Context) error {
	fmt.Println(req.CallbackCommand)
	return nil
}

type CallbackFriendAdd struct {
	CallbackCommand string `json:"CallbackCommand"`
	PairList        []struct {
		FromAccount      string `json:"From_Account"`
		ToAccount        string `json:"To_Account"`
		InitiatorAccount string `json:"Initiator_Account"`
	} `json:"PairList"`
	ClientCmd    string `json:"ClientCmd"`
	AdminAccount string `json:"Admin_Account"`
	ForceFlag    int    `json:"ForceFlag"`
}

type FriendAddCallBack func(req Req, info CallbackFriendAdd) error

func (s IMServer) CallBackFriendAdd(commandFunc FriendAddCallBack) {
	AddCommandFuc(SnsFriendAdd, func(req Req, c echo.Context) error {
		var info CallbackFriendAdd
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}

func (s IMServer) CallBackFriendDelete(commandFunc FriendAddCallBack) {
	AddCommandFuc(SnsFriendDelete, func(req Req, c echo.Context) error {
		var info CallbackFriendAdd
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}

func (s IMServer) CallBackBlackListAdd(commandFunc FriendAddCallBack) {
	AddCommandFuc(SnsBlackListAdd, func(req Req, c echo.Context) error {
		var info CallbackFriendAdd
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}

func (s IMServer) CallBackBlackListDelete(commandFunc FriendAddCallBack) {
	AddCommandFuc(SnsBlackListDelete, func(req Req, c echo.Context) error {
		var info CallbackFriendAdd
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}
