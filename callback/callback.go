/*
 *
 * callback.go
 * callback
 *
 * Created by lintao on 2020/6/1 5:17 下午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package callback

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Req struct {
	SdkAppid        string `json:"SdkAppid" form:"SdkAppid" query:"SdkAppid"`
	CallbackCommand string `json:"CallbackCommand" form:"CallbackCommand" query:"CallbackCommand"`
	Contenttype     string `json:"contenttype" form:"contenttype" query:"contenttype"`
	ClientIP        string `json:"ClientIP" form:"ClientIP" query:"ClientIP"`
	OptPlatform     string `json:"OptPlatform" form:"OptPlatform" query:"OptPlatform"`
}

type Response struct {
	ActionStatus string `json:"ActionStatus"`
	ErrorInfo    string `json:"ErrorInfo"`
	ErrorCode    int    `json:"ErrorCode"`
}

func Run(appId string) {
	e := echo.New()
	e.POST("/", func(c echo.Context) error {
		var req Req
		var response Response
		if err := c.Bind(&req); err != nil {
			response.ActionStatus = "fail"
			response.ErrorCode = 400
			response.ErrorInfo = err.Error()
			return c.JSON(http.StatusOK, response)
		}

		if err := command(req, c); err != nil {
			response.ActionStatus = "fail"
			response.ErrorCode = 400
			response.ErrorInfo = err.Error()
			return c.JSON(http.StatusOK, response)
		}

		response.ActionStatus = "OK"
		response.ErrorCode = 0
		return c.JSON(http.StatusOK, response)
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func command(req Req, c echo.Context) error {

	return nil
}
