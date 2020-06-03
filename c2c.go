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

//CallbackCommand	String	回调命令
//From_Account	String	消息发送者 UserID
//To_Account	String	消息接收者 UserID
//MsgSeq	Integer	消息序列号，用于标记该条消息（32位无符号整数）
//MsgRandom	Integer	消息随机数，用于标记该条消息（32位无符号整数）
//MsgTime	Integer	消息的发送时间戳，单位为秒
//单聊消息优先使用 MsgTime 进行排序，同一秒发送的消息则按 MsgSeq 排序，MsgSeq 值越大消息越靠后
//MsgKey	String	该条消息的唯一标识，可根据该标识进行 REST API 撤回单聊消息
//MsgBody	Array	消息体，详情请参见 消息格式描述
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

//发单聊消息之前回调
//App 后台可以通过该回调实时监控用户的单聊消息，包括：
//对发单聊消息进行实时记录（例如记录日志，或者同步到其他系统）。
//拦截用户的单聊发言请求。可拦截所有类型的消息，如文本、图像、自定义消息等。
//修改用户发言内容（例如敏感词过滤，或者增加一些 App 自定义信息）。目前不支持修改语音、图像、文件和视频等富媒体的消息内容，但支持将这些富媒体消息修改成文本、自定义消息等。
//可能触发该回调的场景
//App 用户通过客户端发送单聊消息。
//App 管理员通过 REST API （sendmsg 接口）发送单聊消息
func (s IMServer) CallbackBeforeSendMsg(commandFunc C2CCallBack) {
	AddCommandFuc(C2CBeforeSendMsg, func(req Req, c echo.Context) error {
		var info C2CMsg
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}

//发单聊消息之后回调
//功能说明
//App 后台可以通过该回调实时监控用户的单聊消息，包括：
//对单聊消息进行实时记录（例如记录日志，或者同步到其他系统）。
//对单聊消息进行数据统计（例如人数，消息数等）。
//可能触发该回调的场景
//App 用户通过客户端发送单聊消息。
//App 管理员通过 REST API（sendmsg 接口）发送单聊消息。
//若消息下发失败（例如被脏字过滤），仍会触发该回调。
func (s IMServer) CallbackAfterSendMsg(commandFunc C2CCallBack) {
	AddCommandFuc(C2CAfterSendMsg, func(req Req, c echo.Context) error {
		var info C2CMsg
		if err := c.Bind(&info); err != nil {
			return err
		}
		return commandFunc(req, info)
	})
}
