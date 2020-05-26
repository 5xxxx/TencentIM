/*
 *
 * ban_speak.go
 * server
 *
 * Created by lintao on 2020/4/17 2:12 下午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 * 全局禁言管理。
 *
 */

package TencentIM

import "encoding/json"

//https://cloud.tencent.com/document/product/269/4230
//设置帐号的群组消息全局禁言。
//C2CmsgNospeakingTime 和 GroupmsgNospeakingTime 是选填字段，但不能两个都不填
//account	String	必填	设置禁言配置的帐号
//C2CmsgNospeakingTime	Integer	选填	单聊消息禁言时间，单位为秒，非负整数，
//最大值为4294967295（十六进制 0xFFFFFFFF）。等于0代表取消帐号禁言；
//等于最大值4294967295（十六进制 0xFFFFFFFF）代表帐号被设置永久禁言；其它代表该帐号禁言时间
//GroupmsgNospeakingTime	Integer	选填	群组消息禁言时间，单位为秒，非负整数，最大值为4294967295（十六进制 0xFFFFFFFF）。
//等于0代表取消帐号禁言；最大值4294967295（十六进制 0xFFFFFFFF）代表帐号被设置永久禁言；其它代表该帐号禁言时间
func (s IMServer) SetNoSpeaking(account string, C2CmsgNospeakingTime, groupmsgNospeakingTime int64) error {
	var req struct {
		Account                string `json:"Set_Account"`
		C2CmsgNospeakingTime   int64  `json:"C2CmsgNospeakingTime"`
		GroupmsgNospeakingTime int64  `json:"GroupmsgNospeakingTime"`
	}
	req.Account = account
	req.C2CmsgNospeakingTime = C2CmsgNospeakingTime
	req.GroupmsgNospeakingTime = groupmsgNospeakingTime
	if _, err := s.requestWithPath(SETNOSPEAKING, req); err != nil {
		return err
	}

	return nil
}

//https://cloud.tencent.com/document/product/269/4229
//查询帐号的群组消息全局禁言。
//C2CmsgNospeakingTime	Number	单聊消息禁言时长，单位为秒，非负整数。等于 0 代表没有被设置禁言；
//等于最大值4294967295（十六进制 0xFFFFFFFF）代表被设置永久禁言；其它代表该帐号禁言时长，如果等于3600表示该帐号被禁言一小时
//GroupmsgNospeakingTime	Number	群组消息禁言时长，单位为秒，非负整数。等于0代表没有被设置禁言；
//等于最大值4294967295（十六进制 0xFFFFFFFF）代表被设置永久禁言；其它代表该帐号禁言时长，如果等于3600表示该帐号被禁言一小时
func (s IMServer) GetNoSpeaking(account string) (C2CmsgNospeakingTime int64, groupmsgNospeakingTime int64, err error) {
	v, err := s.requestWithPath(GETNOSPEAKING, []byte(`{"Get_Account":"`+account+`"}`))
	if err != nil {
		return 0, 0, err
	}

	var resp struct {
		ErrorCode              int    `json:"ErrorCode"`
		ErrorInfo              string `json:"ErrorInfo"`
		C2CmsgNospeakingTime   int64  `json:"C2CmsgNospeakingTime"`
		GroupmsgNospeakingTime int64  `json:"GroupmsgNospeakingTime"`
	}

	if err = json.Unmarshal(v, &resp); err != nil {
		return 0, 0, err
	}

	return resp.C2CmsgNospeakingTime, resp.GroupmsgNospeakingTime, nil
}
