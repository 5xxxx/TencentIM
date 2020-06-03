/*
 *
 * state.go
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

func StateChange(req Req, c echo.Context) error {
	fmt.Println(req.CallbackCommand)
	return nil
}

type StateInfo struct {
	CallbackCommand string `json:"CallbackCommand"`
	Info            struct {
		Action    string `json:"Action"`
		ToAccount string `json:"To_Account"`
		Reason    string `json:"Reason"`
	} `json:"Info"`
}

type StateCallBack func(req Req, info StateInfo) error

func (s IMServer) CallbackStateChange(commandFunc StateCallBack) {
	AddCommandFuc(StateStateChange, func(req Req, c echo.Context) error {
		var info StateInfo
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}
