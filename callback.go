/*
 *
 * callback.go
 * callback
 *
 * Created by lintao on 2020/6/1 5:17 下午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package TencentIM

import (
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
)

var mutex sync.Mutex

const (
	StateStateChange           = "State.StateChang"
	SnsFriendAdd               = "Sns.CallbackFriendAdd"
	SnsFriendDelete            = "Sns.CallbackFriendDelete"
	SnsBlackListAdd            = "Sns.CallbackBlackListAdd"
	SnsBlackListDelete         = "Sns.CallbackBlackListDelete"
	C2CBeforeSendMsg           = "C2C.CallbackBeforeSendMsg"
	C2CAfterSendMsg            = "C2C.CallbackAfterSendMsg"
	GroupBeforeCreateGroup     = "Group.CallbackBeforeCreateGroup"
	GroupAfterCreateGroup      = "Group.CallbackAfterCreateGroup"
	GroupBeforeApplyJoinGroup  = "Group.CallbackBeforeApplyJoinGroup"
	GroupBeforeInviteJoinGroup = "Group.CallbackBeforeInviteJoinGroup"
	GroupAfterNewMemberJoin    = "Group.CallbackAfterNewMemberJoin"
	GroupAfterMemberExit       = "Group.CallbackAfterMemberExit"
	GroupBeforeSendMsg         = "Group.CallbackBeforeSendMsg"
	GroupAfterSendMsg          = "Group.CallbackAfterSendMsg"
	GroupAfterGroupFull        = "Group.CallbackAfterGroupFull"
	GroupAfterGroupDestroyed   = "Group.CallbackAfterGroupDestroyed"
	GroupAfterGroupInfoChanged = "Group.CallbackAfterGroupInfoChanged"
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
	e.POST("/imcallback", func(c echo.Context) error {
		var req Req
		var response Response
		if err := c.Bind(&req); err != nil {
			response.ActionStatus = "fail"
			response.ErrorCode = 400
			response.ErrorInfo = err.Error()
			return c.JSON(http.StatusOK, response)
		}

		if appId != req.SdkAppid {
			return nil
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
	return h[req.CallbackCommand](req, c)
}

func AddCommandFuc(command string, commandFunc CommandFunc) {
	mutex.Lock()
	defer mutex.Unlock()
	h[command] = commandFunc
}

var h = map[string]CommandFunc{
	"State.StateChang":                    StateChange,
	"Sns.CallbackFriendAdd":               FriendAdd,
	"Sns.CallbackFriendDelete":            FriendDelete,
	"Sns.CallbackBlackListAdd":            BlackListAdd,
	"Sns.CallbackBlackListDelete":         BlackListDelete,
	"C2C.CallbackBeforeSendMsg":           BeforeSendMsg,
	"C2C.CallbackAfterSendMsg":            AfterSendMsg,
	"Group.CallbackBeforeCreateGroup":     BeforeCreateGroup,
	"Group.CallbackAfterCreateGroup":      AfterCreateGroup,
	"Group.CallbackBeforeApplyJoinGroup":  BeforeApplyJoinGroup,
	"Group.CallbackBeforeInviteJoinGroup": BeforeInviteJoinGroup,
	"Group.CallbackAfterNewMemberJoin":    AfterNewMemberJoin,
	"Group.CallbackAfterMemberExit":       AfterMemberExit,
	"Group.CallbackBeforeSendMsg":         BeforeSendMsg,
	"Group.CallbackAfterSendMsg":          GafterSendMsg,
	"Group.CallbackAfterGroupFull":        AfterGroupFull,
	"Group.CallbackAfterGroupDestroyed":   AfterGroupDestroyed,
	"Group.CallbackAfterGroupInfoChanged": AfterGroupInfoChanged,
}

type CommandFunc func(req Req, c echo.Context) error
