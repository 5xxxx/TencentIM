/*
 *
 * app_info.go
 * server
 *
 * Created by lintao on 2020/4/17 2:14 下午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 * 运营管理
 */

package TencentIM

import "encoding/json"

// https://cloud.tencent.com/document/product/269/4193
// 拉取运营数据
// App 管理员可以通过该接口拉取最近30天的运营数据，可拉取的字段见下文可拉取的运营字段
// Field 选填	该字段用来指定需要拉取的运营数据，不填默认拉取所有字段。详细可参阅下文可拉取的运营字段
func (s IMServer) GetAppInfo(field ...AppInfoField) ([]Info, error) {

	var req struct {
		RequestField []AppInfoField `json:"RequestField"`
	}
	req.RequestField = make([]AppInfoField, 0)
	for _, f := range field {
		req.RequestField = append(req.RequestField, f)
	}

	v, err := s.requestWithPath(GETAPPINFO, nil)
	if err != nil {
		return nil, err
	}

	var appInfo AppInfo
	if err = json.Unmarshal(v, &appInfo); err != nil {
		return nil, err
	}

	return appInfo.Result, nil
}

// 下载消息记录
// https://cloud.tencent.com/document/product/269/1650
//App 管理员可以通过该接口获取 App 中某天某小时的所有单发或群组消息记录的下载地址。
//消息记录以日志文件形式保存并使用 GZip 压缩，通过该接口获取到下载地址后，请自行下载并处理；消息记录文件每小时产生一次，
//例如0点（00:00~00:59）的数据在01:00后开始处理，一般1小时内处理完毕（消息较多则处理时间较长）；
//文件有效期3天，无论是否下载过，都会在3天后删除；获取到的下载地址存在有效期，请在过期前进行下载，若地址失效，请通过该接口重新获取
//ChatType	String	必填	消息类型，C2C 表示单发消息 Group 表示群组消息
//MsgTime	String	必填	需要下载的消息记录的时间段，2015120121表示获取2015年12月1日21:00 - 21:59的消息的下载地址。该字段需精确到小时。每次请求只能获取某天某小时的所有单发或群组消息记录
func (s IMServer) GetHistory(chatType, msgTime string) ([]Msg, error) {
	var req struct {
		ChatType string `json:"ChatType"`
		MsgTime  string `json:"MsgTime"`
	}
	req.ChatType = chatType
	req.MsgTime = msgTime
	v, err := s.requestWithPath(GET_HISTORY, req)
	if err != nil {
		return nil, err
	}
	var resp AppHistory
	if err = json.Unmarshal(v, &resp); err != nil {
		return nil, err
	}

	return resp.File, err
}
