/*
 *
 * group.go
 * callback
 *
 * Created by lintao on 2020/6/1 5:17 下午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package TencentIM

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func BeforeCreateGroup(req Req, c echo.Context) error {
	fmt.Println(req.CallbackCommand)
	return nil
}

func AfterCreateGroup(req Req, c echo.Context) error {
	fmt.Println(req.CallbackCommand)
	return nil
}

func BeforeApplyJoinGroup(req Req, c echo.Context) error {
	fmt.Println(req.CallbackCommand)
	return nil
}

func BeforeInviteJoinGroup(req Req, c echo.Context) error {
	fmt.Println(req.CallbackCommand)
	return nil
}

func AfterNewMemberJoin(req Req, c echo.Context) error {
	fmt.Println(req.CallbackCommand)
	return nil
}

func AfterMemberExit(req Req, c echo.Context) error {
	fmt.Println(req.CallbackCommand)
	return nil
}

func GafterSendMsg(req Req, c echo.Context) error {
	fmt.Println(req.CallbackCommand)
	return nil
}

func AfterGroupFull(req Req, c echo.Context) error {
	fmt.Println(req.CallbackCommand)
	return nil
}

func AfterGroupDestroyed(req Req, c echo.Context) error {
	fmt.Println(req.CallbackCommand)
	return nil
}

func AfterGroupInfoChanged(req Req, c echo.Context) error {
	fmt.Println(req.CallbackCommand)
	return nil
}

type CallBackGroup struct {
	CallbackCommand  string `json:"CallbackCommand"`
	GroupID          string `json:"GroupId"`
	OperatorAccount  string `json:"Operator_Account"`
	OwnerAccount     string `json:"Owner_Account"`
	Type             string `json:"Type"`
	Name             string `json:"Name"`
	RequestorAccount string `json:"Requestor_Account" `
	MemberList       []struct {
		MemberAccount string `json:"Member_Account"`
	} `json:"MemberList"`
	UserDefinedDataList []struct {
		Key   string `json:"Key"`
		Value string `json:"Value"`
	} `json:"UserDefinedDataList"`
	DestinationMembers []struct {
		MemberAccount string `json:"Member_Account"`
	} `json:"DestinationMembers"`
	JoinType      string `json:"JoinType"`
	NewMemberList []struct {
		MemberAccount string `json:"Member_Account"`
	} `json:"NewMemberList"`
	ExitType       string `json:"ExitType"`
	ExitMemberList []struct {
		MemberAccount string `json:"Member_Account"`
	} `json:"ExitMemberList"`
	FromAccount string `json:"From_Account"`
	Random      int    `json:"Random"`
	MsgBody     []struct {
		MsgType    string `json:"MsgType"`
		MsgContent struct {
			Text string `json:"Text"`
		} `json:"MsgContent"`
	} `json:"MsgBody"`
	MsgSeq       int    `json:"MsgSeq"`
	MsgTime      int    `json:"MsgTime"`
	Introduction string `json:"Introduction"`
	Notification string `json:"Notification"`
	FaceURL      string `json:"FaceUrl"`
}

type CallGroup func(req Req, info CallBackGroup) error

func (s IMServer) CallBackBeforeCreateGroup(commandFunc CallGroup) {
	AddCommandFuc(GroupBeforeCreateGroup, func(req Req, c echo.Context) error {
		var info CallBackGroup
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}

func (s IMServer) CallBackAfterCreateGroup(commandFunc CallGroup) {
	AddCommandFuc(GroupAfterCreateGroup, func(req Req, c echo.Context) error {
		var info CallBackGroup
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}

func (s IMServer) CallBackBeforeApplyJoinGroup(commandFunc CallGroup) {
	AddCommandFuc(GroupBeforeApplyJoinGroup, func(req Req, c echo.Context) error {
		var info CallBackGroup
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}

func (s IMServer) CallBackBeforeInviteJoinGroup(commandFunc CallGroup) {
	AddCommandFuc(GroupBeforeInviteJoinGroup, func(req Req, c echo.Context) error {
		var info CallBackGroup
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}

func (s IMServer) CallBackAfterNewMemberJoin(commandFunc CallGroup) {
	AddCommandFuc(GroupAfterNewMemberJoin, func(req Req, c echo.Context) error {
		var info CallBackGroup
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}

func (s IMServer) CallBackAfterMemberExit(commandFunc CallGroup) {
	AddCommandFuc(GroupAfterMemberExit, func(req Req, c echo.Context) error {
		var info CallBackGroup
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}

func (s IMServer) CallBackGroupBeforeSendMsg(commandFunc CallGroup) {
	AddCommandFuc(GroupBeforeSendMsg, func(req Req, c echo.Context) error {
		var info CallBackGroup
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}

func (s IMServer) CallBackGroupAfterSendMsg(commandFunc CallGroup) {
	AddCommandFuc(GroupAfterSendMsg, func(req Req, c echo.Context) error {
		var info CallBackGroup
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}

func (s IMServer) CallBackGroupAfterGroupFull(commandFunc CallGroup) {
	AddCommandFuc(GroupAfterGroupFull, func(req Req, c echo.Context) error {
		var info CallBackGroup
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}

func (s IMServer) CallBackAfterGroupDestroyed(commandFunc CallGroup) {
	AddCommandFuc(GroupAfterGroupDestroyed, func(req Req, c echo.Context) error {
		var info CallBackGroup
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}

func (s IMServer) CallBackAfterGroupInfoChanged(commandFunc CallGroup) {
	AddCommandFuc(GroupAfterGroupInfoChanged, func(req Req, c echo.Context) error {
		var info CallBackGroup
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}
