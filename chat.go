/*
 *
 * ptp_chat.go
 * server
 *
 * Created by lintao on 2020/4/17 9:29 上午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 * 单聊消息
 */

package TencentIM

import "encoding/json"

//单发单聊消息
//管理员向帐号发消息，接收方看到消息发送者是管理员。
//管理员指定某一帐号向其他帐号发消息，接收方看到发送者不是管理员，而是管理员指定的帐号。
//该接口不会检查发送者和接收者的好友关系（包括黑名单），同时不会检查接收者是否被禁言
//MsgTime	Integer	消息时间戳，UNIX 时间戳
//MsgKey	String	消息唯一标识，用于撤回。长度不超过50个字符
func (s IMServer) SendMsg(b SingleChatMsg) (msgTime int64, msgKey string, err error) {
	//for _, c := range b.MsgBody {
	//	c.CheckType()
	//}
	if b.FromAccount == "" {
		b.FromAccount = s.Identifier
	}

	var resp struct {
		ActionStatus string `json:"ActionStatus"`
		ErrorInfo    string `json:"ErrorInfo"`
		ErrorCode    int    `json:"ErrorCode"`
		MsgTime      int64  `json:"MsgTime"`
		MsgKey       string `json:"MsgKey"`
	}

	var v []byte
	if v, err = s.requestWithPath(SENDMSG, b); err != nil {
		return
	}
	if err = json.Unmarshal(v, &resp); err != nil {
		return
	}

	return resp.MsgTime, resp.MsgKey, nil
}

//批量发单聊消息
//支持一次对最多500个用户进行单发消息。
//与单发消息相比，该接口更适用于营销类消息、系统通知 tips 等时效性较强的消息。
//管理员指定某一帐号向目标帐号批量发消息，接收方看到发送者不是管理员，而是管理员指定的帐号。
//该接口不触发回调请求。
//该接口不会检查发送者和接收者的好友关系（包括黑名单），同时不会检查接收者是否被禁言
func (s IMServer) BATCHChatMsg(b BatchChatMsg) (msgKey string, errList []ErrorList, err error) {

	if b.FromAccount == "" {
		b.FromAccount = s.Identifier
	}
	var resp struct {
		ActionStatus string      `json:"ActionStatus"`
		ErrorInfo    string      `json:"ErrorInfo"`
		MsgKey       string      `json:"MsgKey"`
		ErrorList    []ErrorList `json:"ErrorList"`
	}

	var v []byte
	if v, err = s.requestWithPath(BATCHSENDMSG, b); err != nil {
		return
	}
	if err = json.Unmarshal(v, &resp); err != nil {
		return
	}

	return resp.MsgKey, resp.ErrorList, nil
}

//导入单聊消息
//导入历史单聊消息到即时通信 IM。
//平滑过渡期间，将原有即时通信实时单聊消息导入到即时通信 IM。
//该接口不会触发回调。
//该接口会根据 From_Account ， To_Account ， MsgRandom ， MsgTimeStamp 字段的值对导入的消息进行去重。仅当这四个字段的值都对应相同时，才判定消息是重复的，消息是否重复与消息内容本身无关。
//重复导入的消息不会覆盖之前已导入的消息（即消息内容以首次导入的为准）。
func (s IMServer) ImportMsg(b SingleChatMsg) (err error) {

	if b.FromAccount == "" {
		b.FromAccount = s.Identifier
	}

	if _, err = s.requestWithPath(IMPORTMSG, b); err != nil {
		return
	}

	return nil
}

//查询单聊消息
//管理员按照时间范围查询某单聊会话的消息记录。
//查询的单聊会话由请求中的 From_Account 和 to_Account 指定。查询结果包含会话双方互相发送的消息，具体每条消息的发送方和接收方由每条消息里的 From_Account 和 to_Account 指定。
//请求中的 From_Account 和 to_Account 字段值互换，查询结果不变。
//查询结果包含被撤回的消息，由消息里的 MsgFlagBits 字段标识。
//若想通过 REST API 撤回单聊消息 接口撤回某条消息，可先用本接口查询出该消息的 MsgKey，然后再调用撤回接口进行撤回。
//可查询的消息记录的时间范围取决于漫游消息存储时长，默认是7天。支持在控制台修改消息漫游时长，延长消息漫游时长是增值服务。具体请参考 漫游消息存储。
//首次查询时，建议 MaxCnt 填20，应答里的 Complete 字段表示是否已拉完该时间段内的消息。
func (s IMServer) GetRoamMsg(m RoamMsgReq) (err error) {

	var v []byte
	if v, err = s.requestWithPath(ADMIN_GETROAMMSG, m); err != nil {
		return
	}

	var resp RoamMsgResp
	if err = json.Unmarshal(v, &resp); err != nil {
		return
	}

	return
}

//撤回单聊消息
//管理员撤回单聊消息。
//该接口可以撤回所有单聊消息，包括客户端发出的单聊消息，由 REST API 单发 和 批量发 接口发出的单聊消息。
//若需要撤回由客户端发出的单聊消息，您可以开通 发单聊消息之前回调 或 发单聊消息之后回调 ，通过该回调接口记录每条单聊消息的 MsgKey ，然后填在本接口的 MsgKey 字段进行撤回。您也可以通过 查询单聊消息 查询出待撤回的单聊消息的 MsgKey 后，填在本接口的 MsgKey 字段进行撤回。
//若需要撤回由 REST API 单发 和 批量发 接口发出的单聊消息，需要记录这些接口回包里的 MsgKey 字段以进行撤回。
//调用该接口撤回消息后，该条消息的离线、漫游存储，以及消息发送方和接收方的客户端的本地缓存都会被撤回。
//该接口可撤回的单聊消息没有时间限制，即可以撤回任何时间的单聊消息。
func (s IMServer) MsgWithdraw(m MsgWithdraw) error {
	if _, err := s.requestWithPath(ADMIN_MSGWITHDRAW, m); err != nil {
		return err
	}

	return nil
}
