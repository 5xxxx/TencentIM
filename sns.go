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

//CallbackCommand	String	回调命令
//PairList	Array	成功添加的好友对
//From_Account	String	From_Account 的好友表中增加了 To_Account
//To_Account	String	To_Account 被增加到了 From_Account 的好友表中
//Initiator_Account	String	发起加好友请求的用户的 UserID
//ClientCmd	String	触发回调的命令字：
//加好友请求，合理的取值如下：friend_add、FriendAdd
//加好友回应，合理的取值如下：friend_response、FriendResponse
//Admin_Account	String	如果当前请求是后台触发的加好友请求，则该字段被赋值为管理员帐号；否则为空
//ForceFlag	Integer	管理员强制加好友标记：1 表示强制加好友；0 表示常规加好友方式
type CallbackFriend struct {
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

type FriendAddCallBack func(req Req, info CallbackFriend) error

//添加好友之后回调
//App 后台可以通过该回调实时监控用户的新增好友信息。
//可能触发该回调的场景
//App 后台通过 REST API 发起加好友请求，请求添加双向好友，且对方的加好友验证方式是“允许任何人”。
//App 用户通过客户端发起加好友请求，请求添加双向好友，且对方的加好友验证方式是“允许任何人”。
//App 后台通过 REST API 发起加好友请求，请求添加单向好友。
//App 用户通过客户端发起加好友请求，请求添加单向好友。
//App 用户收到加好友请求后，同意添加对方为好友。
//App 后台通过 REST API 强制加好友
//通过调用 导入好友 接口添加好友时，不会触发此回调。
func (s IMServer) CallbackFriendAdd(commandFunc FriendAddCallBack) {
	AddCommandFuc(SnsFriendAdd, func(req Req, c echo.Context) error {
		var info CallbackFriend
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}

//删除好友之后回调
//App 后台可以通过该回调实时监控用户的好友删除信息。
//可能触发该回调的场景
//App 用户通过客户端发起删除好友的请求。
//App 后台通过 REST API 发起删除好友的请求。
func (s IMServer) CallbackFriendDelete(commandFunc FriendAddCallBack) {
	AddCommandFuc(SnsFriendDelete, func(req Req, c echo.Context) error {
		var info CallbackFriend
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}

//添加黑名单之后回调
//App 后台可以通过该回调实时监控黑名单的添加情况。
//可能触发该回调的场景
//App 用户通过客户端发起添加黑名单的请求。
//App 后台通过 REST API 发起添加黑名单的请求。
func (s IMServer) CallbackBlackListAdd(commandFunc FriendAddCallBack) {
	AddCommandFuc(SnsBlackListAdd, func(req Req, c echo.Context) error {
		var info CallbackFriend
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}

//删除黑名单之后回调
//App 后台可以通过该回调实时监控用户黑名单的删除情况。
//可能触发该回调的场景
//App 用户通过客户端发起删除黑名单请求。
//App 后台通过 REST API 发起删除黑名单请求。
func (s IMServer) CallbackBlackListDelete(commandFunc FriendAddCallBack) {
	AddCommandFuc(SnsBlackListDelete, func(req Req, c echo.Context) error {
		var info CallbackFriend
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}
