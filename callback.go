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
	StateStateChange           = "State.StateChange"
	SnsFriendAdd               = "Sns.CallbackFriend"
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
	SdkAppid        string `query:"SdkAppid"`
	CallbackCommand string `query:"CallbackCommand"`
	Contenttype     string `query:"contenttype"`
	ClientIP        string `query:"ClientIP"`
	OptPlatform     string `query:"OptPlatform"`
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
		req.SdkAppid = c.QueryParam("SdkAppid")
		req.CallbackCommand = c.QueryParam("CallbackCommand")
		req.OptPlatform = c.QueryParam("OptPlatform")
		req.ClientIP = c.QueryParam("ClientIP")
		req.Contenttype = c.QueryParam("contenttype")

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
	e.Logger.Fatal(e.Start(":80"))
}

func command(req Req, c echo.Context) error {
	if cmd, ok := h[req.CallbackCommand]; ok {
		return cmd(req, c)
	}

	return nil
}

func AddCommandFuc(command string, commandFunc CommandFunc) {
	mutex.Lock()
	defer mutex.Unlock()
	h[command] = commandFunc
}

var h = map[string]CommandFunc{
	StateStateChange:           StateChange,
	SnsFriendAdd:               FriendAdd,
	SnsFriendDelete:            FriendDelete,
	SnsBlackListAdd:            BlackListAdd,
	SnsBlackListDelete:         BlackListDelete,
	C2CBeforeSendMsg:           BeforeSendMsg,
	C2CAfterSendMsg:            AfterSendMsg,
	GroupBeforeCreateGroup:     BeforeCreateGroup,
	GroupAfterCreateGroup:      AfterCreateGroup,
	GroupBeforeApplyJoinGroup:  BeforeApplyJoinGroup,
	GroupBeforeInviteJoinGroup: BeforeInviteJoinGroup,
	GroupAfterNewMemberJoin:    AfterNewMemberJoin,
	GroupAfterMemberExit:       AfterMemberExit,
	GroupBeforeSendMsg:         BeforeSendMsg,
	GroupAfterSendMsg:          GafterSendMsg,
	GroupAfterGroupFull:        AfterGroupFull,
	GroupAfterGroupDestroyed:   AfterGroupDestroyed,
	GroupAfterGroupInfoChanged: AfterGroupInfoChanged,
}

type CommandFunc func(req Req, c echo.Context) error
