/*
 *
 * data_manage.go
 * server
 *
 * Created by lintao on 2020/4/17 11:36 上午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 * 用户资料管理
 */

package TencentIM

import "encoding/json"

// https://cloud.tencent.com/document/product/269/1639
// 拉取资料
// 支持拉取好友和非好友的资料字段。
// 支持拉取 标配资料字段 和 自定义资料字段。
// 建议每次拉取的用户数不超过100，避免因回包数据量太大导致回包失败。
// 请确保请求中的所有帐号都已导入即时通信 IM，如果请求中含有未导入即时通信 IM 的帐号，即时通信 IM 后台将会提示错误。
//
func (s IMServer) GetPortrait(p PortraitReq) (resp PortraitResp, err error) {
	var v []byte

	if v, err = s.requestWithPath(PORTRAIT_GET, p); err != nil {
		return
	}

	if err = json.Unmarshal(v, &resp); err != nil {
		return
	}

	return
}

//设置资料
//支持 标配资料字段 和 自定义资料字段 的设置。
func (s IMServer) SetPortrait(p PortraitSet) error {
	if _, err := s.requestWithPath(PORTRAIT_SET, p); err != nil {
		return err
	}
	return nil
}
