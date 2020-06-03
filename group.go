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

//CallbackCommand	String	回调命令
//GroupId	String	操作的群 ID
//Operator_Account	String	发起创建群组请求的操作者 UserID
//Owner_Account	String	请求创建的群的群主 UserID
//Type	String	请求创建的 群组形态介绍，例如 Private，Public 和 ChatRoom
//Name	String	请求创建的群组的名称
//MemberList	Array	请求创建的群组的初始化成员列表
//UserDefinedDataList	Array	用户建群时的自定义字段，这个字段默认是没有的，需要开通，详见 自定义字段
type CallbackGroup struct {
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

type CallGroup func(req Req, info CallbackGroup) error

//创建群组之前回调
//功能说明
//App 后台可以通过该回调实时监控用户创建群组的请求，包括后台可以拒绝用户创建群组的请求。
//可能触发该回调的场景
//App 用户通过客户端创建群组
//App 管理员通过 REST API 创建群组
func (s IMServer) CallbackBeforeCreateGroup(commandFunc CallGroup) {
	AddCommandFuc(GroupBeforeCreateGroup, func(req Req, c echo.Context) error {
		var info CallbackGroup
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}

//创建群组之后回调
//功能说明
//App 后台可以通过该回调实时监控用户创建群组的信息，包括： 通知 App 后台有群组创建成功，App 后台可以据此进行数据同步等操作。
//可能触发该回调的场景
//App 用户通过客户端创建群组成功
//App 管理员通过 REST API 创建群组成功
func (s IMServer) CallbackAfterCreateGroup(commandFunc CallGroup) {
	AddCommandFuc(GroupAfterCreateGroup, func(req Req, c echo.Context) error {
		var info CallbackGroup
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}

//申请入群之前回调
//功能说明
//App 后台可以通过该回调实时监控用户申请加群的请求，包括：App 后台可以拦截用户申请加群的操作。
//可能触发该回调的场景
//App 用户通过客户端发起加群申请
func (s IMServer) CallbackBeforeApplyJoinGroup(commandFunc CallGroup) {
	AddCommandFuc(GroupBeforeApplyJoinGroup, func(req Req, c echo.Context) error {
		var info CallbackGroup
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}

//拉人入群之前回调
//功能说明
//App 后台可以通过该回调实时监控群成员拉其他用户入群的请求，包括：App 后台可以拦截群成员直接将其他用户拉入群的请求
//可能触发该回调的场景
//App 用户通过客户端发起将其他用户拉入群的请求。
//App 管理员通过 REST API 添加用户到群组。
func (s IMServer) CallbackBeforeInviteJoinGroup(commandFunc CallGroup) {
	AddCommandFuc(GroupBeforeInviteJoinGroup, func(req Req, c echo.Context) error {
		var info CallbackGroup
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}

//新成员入群之后回调
//功能说明
//App 后台可以通过该回调实时监控群成员加入的消息，包括：通知 App 后台有成员入群，App 可以据此进行必要的数据同步。
//可能触发该回调的场景
//App 用户通过客户端主动申请加群并得到通过。
//App 用户通过客户端拉其他人入群成功。
//App 管理员通过 REST API 添加用户到群组。
func (s IMServer) CallbackAfterNewMemberJoin(commandFunc CallGroup) {
	AddCommandFuc(GroupAfterNewMemberJoin, func(req Req, c echo.Context) error {
		var info CallbackGroup
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}

//群成员离开之后回调
//功能说明
//App 后台可以通过该回调实时监控用户的退群动态，包括：对用户退群进行实时记录（例如记录日志，或者同步到其他系统）。
//可能触发该回调的场景
//App 用户通过客户端退群。
//App 用户通过客户端踢人。
//App 管理员通过 REST API 删除群成员
func (s IMServer) CallbackAfterMemberExit(commandFunc CallGroup) {
	AddCommandFuc(GroupAfterMemberExit, func(req Req, c echo.Context) error {
		var info CallbackGroup
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}

//群内发言之前回调
//功能说明
//App 后台可以通过该回调实时监控用户的群发消息，包括：
//对群消息进行实时记录（例如记录日志，或者同步到其他系统）。
//拦截用户在群内发言的请求。
//修改用户发言内容（例如敏感词过滤，或者增加一些 App 自定义信息）
//可能触发该回调的场景
//App 用户通过客户端发送群消息。
//App 管理员通过 REST API 发送群组消息。
func (s IMServer) CallbackGroupBeforeSendMsg(commandFunc CallGroup) {
	AddCommandFuc(GroupBeforeSendMsg, func(req Req, c echo.Context) error {
		var info CallbackGroup
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}

//群内发言之后回调
//功能说明
//App 后台可以通过该回调实时监控用户的群发消息，包括：通知 App 后台有群组消息发送成功，App 可以据此进行必要的数据同步。
//可能触发该回调的场景
//App 用户通过客户端发送群消息。
//App 管理员通过 REST API 发送群组消息。
func (s IMServer) CallbackGroupAfterSendMsg(commandFunc CallGroup) {
	AddCommandFuc(GroupAfterSendMsg, func(req Req, c echo.Context) error {
		var info CallbackGroup
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}

//群组满员之后回调
//功能说明
//App 后台可以通过该回调实时监控群组满员的动态，包括：删除一部分不活跃的群成员，以确保用户能够加入该群。
//可能触发该回调的场景
//App 用户通过客户端申请加群。
//App 用户通过客户端邀请加群。
//App 管理员通过 REST API 增加群组成员。
func (s IMServer) CallbackGroupAfterGroupFull(commandFunc CallGroup) {
	AddCommandFuc(GroupAfterGroupFull, func(req Req, c echo.Context) error {
		var info CallbackGroup
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}

//群组解散之后回调
//功能说明
//App 后台可以通过该回调实时监控群组的解散动态，包括：对群组的解散实时记录（例如记录日志，或者同步到其他系统）。
//可能触发该回调的场景
//App 用户通过客户端解散群组。
//App 管理员通过 REST API 解散群组。
func (s IMServer) CallbackAfterGroupDestroyed(commandFunc CallGroup) {
	AddCommandFuc(GroupAfterGroupDestroyed, func(req Req, c echo.Context) error {
		var info CallbackGroup
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}

//群组资料修改之后回调
//功能说明
//App 后台可以通过该回调实时监控群组资料（群名称、群简介、群公告、群头像及群维度自定义字段）的变更，包括对修改群组资料的实时记录（例如记录日志，或者同步到其他系统）。
//可能触发该回调的场景
//哪些内容会触发回调
//群组资料包括 群基础资料 和 群组维度自定义字段。
//目前，群基础资料中的群组名称、群组简介、群组公告和群组头像 URL 被修改后，可能触发该回调。其他群基础资料被修改后暂不会触发。
//哪些方式会触发回调
//App 用户通过客户端修改群组资料。
//App 管理员通过 REST API 修改群组资料
func (s IMServer) CallbackAfterGroupInfoChanged(commandFunc CallGroup) {
	AddCommandFuc(GroupAfterGroupInfoChanged, func(req Req, c echo.Context) error {
		var info CallbackGroup
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}
