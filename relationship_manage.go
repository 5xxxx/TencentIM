/*
 *
 * relation_manage.go
 * server
 *
 * Created by lintao on 2020/4/17 12:00 下午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 * 关系链管理
 */

package TencentIM

import "encoding/json"

//https://cloud.tencent.com/document/product/269/1643
//添加好友，支持批量添加好友。
func (s IMServer) FriendAdd(f Friend) ([]FriendResultItem, error) {
	result, err := s.requestWithPath(FRIEND_ADD, f)
	if err != nil {
		return nil, err
	}
	var resp struct {
		ResultItem []FriendResultItem `json:"ResultItem"`
	}

	if err = json.Unmarshal(result, &resp); err != nil {
		return nil, err
	}

	return resp.ResultItem, nil
}

// https://cloud.tencent.com/document/product/269/8301
// 导入好友
// 支持批量导入单向好友。
// 往同一个用户导入好友时建议采用批量导入的方式，避免并发写导致的写冲突。
func (s IMServer) FriendImport(f ImportFriend) ([]FriendResultItem, error) {
	v, err := s.requestWithPath(FRIEND_IMPORT, f)
	if err != nil {
		return nil, err
	}

	var resp struct {
		ResultItem []FriendResultItem `json:"ResultItem"`
	}

	if err = json.Unmarshal(v, &resp); err != nil {
		return nil, err
	}

	return resp.ResultItem, nil
}

// https://cloud.tencent.com/document/product/269/12525
// 更新好友
// 支持批量更新同一用户的多个好友的关系链数据。
// 更新一个用户多个好友时，建议采用批量方式，避免并发写导致的写冲突。
func (s IMServer) FriendUpdate(f UpdateFriend) ([]FriendResultItem, error) {
	result, err := s.requestWithPath(FRIEND_UPDATE, f)
	if err != nil {
		return nil, err
	}

	var resp struct {
		ResultItem []FriendResultItem `json:"ResultItem"`
	}

	if err = json.Unmarshal(result, &resp); err != nil {
		return nil, err
	}

	return resp.ResultItem, nil
}

// https://cloud.tencent.com/document/product/269/1644
// 删除好友
// 删除好友，支持单向删除好友和双向删除好友。
// From_Account	String	必填	需要删除该 UserID 的好友
// To_Account	Array	必填	待删除的好友的 UserID 列表，单次请求的 To_Account 数不得超过1000
// DeleteType	String	选填	删除模式，详情可参见
func (s IMServer) FriendDelete(fromAccount string, toAccount []string, deleteType DeleteType) ([]FriendResultItem, error) {
	var req struct {
		FromAccount string     `json:"From_Account"`
		ToAccount   []string   `json:"To_Account"`
		DeleteType  DeleteType `json:"DeleteType"`
	}
	req.ToAccount = toAccount
	req.FromAccount = fromAccount
	req.DeleteType = deleteType
	result, err := s.requestWithPath(FRIEND_DELETE, req)
	if err != nil {
		return nil, err
	}
	var resp struct {
		ResultItem []FriendResultItem `json:"ResultItem"`
	}

	if err = json.Unmarshal(result, &resp); err != nil {
		return nil, err
	}

	return resp.ResultItem, nil
}

// https://cloud.tencent.com/document/product/269/1645
// 删除所有好友
// 清除指定用户的标配好友数据和自定义好友数据。
//From_Account	String	必填	指定要清除好友数据的用户的 UserID
//DeleteType	String	选填	删除模式，默认删除单向好友，详情可参见 删除好友
func (s IMServer) FriendDeleteAll(fromAccount string, deleteType DeleteType) error {
	var req struct {
		FromAccount string     `json:"From_Account"`
		DeleteType  DeleteType `json:"DeleteType"`
	}
	req.DeleteType = deleteType
	req.FromAccount = fromAccount
	_, err := s.requestWithPath(FRIEND_DELETE_ALL, req)
	if err != nil {
		return err
	}

	return nil
}

// https://cloud.tencent.com/document/product/269/1646
// 校验好友
// 支持批量校验好友关系。
// From_Account	String	必填	需要校验该 UserID 的好友
// To_Account	Array	必填	请求校验的好友的 UserID 列表，单次请求的 To_Account 数不得超过1000
// CheckType	String	必填	校验模式，详情可参见 校验好友
func (s IMServer) FriendCheck(fromAccount string, toAccount []string, checkType CheckType) ([]CheckResult, error) {
	var req struct {
		FromAccount string    `json:"From_Account"`
		ToAccount   []string  `json:"To_Account"`
		CheckType   CheckType `json:"CheckType"`
	}
	req.FromAccount = fromAccount
	req.ToAccount = toAccount
	req.CheckType = checkType

	result, err := s.requestWithPath(FRIEND_CHECK, req)
	if err != nil {
		return nil, err
	}
	var resp struct {
		InfoItem []CheckResult `json:"InfoItem"`
	}

	if err = json.Unmarshal(result, &resp); err != nil {
		return nil, err
	}

	return resp.InfoItem, nil
}

// https://cloud.tencent.com/document/product/269/1646
// 拉取好友
// 分页拉取全量好友数据。
// 不支持资料数据的拉取。
// 不需要指定请求拉取的字段，默认返回全量的标配好友数据和自定义好友数据。
func (s IMServer) FriendGet(f GetFriend) (GetFriendResult, error) {
	result, err := s.requestWithPath(FRIEND_GET, f)
	if err != nil {
		return GetFriendResult{}, err
	}
	var resp GetFriendResult
	if err = json.Unmarshal(result, &resp); err != nil {
		return GetFriendResult{}, err
	}

	return resp, err
}

// https://cloud.tencent.com/document/product/269/8609
// 拉取指定好友
// 支持拉取指定好友的好友数据和资料数据。
// 建议每次拉取的好友数不超过100，避免因数据量太大导致回包失败
//From_Account	String	必填	指定要拉取好友数据的用户的 UserID
//To_Account	Array	必填	好友的 UserID 列表
//建议每次请求的好友数不超过100，避免因数据量太大导致回包失败
//TagList	Array	必填	指定要拉取的资料字段及好友字段： 标配资料字段 标配好友字段
func (s IMServer) FriendGetList(fromAccount string, toAccount []string, tagList []ProfileTag) ([]InfoItem, error) {
	var req struct {
		FromAccount string       `json:"From_Account"`
		ToAccount   []string     `json:"To_Account"`
		TagList     []ProfileTag `json:"TagList"`
	}
	req.FromAccount = fromAccount
	req.ToAccount = toAccount
	req.TagList = tagList

	result, err := s.requestWithPath(FRIEND_GET_LIST, req)
	if err != nil {
		return nil, err
	}

	var resp struct {
		InfoItem []InfoItem `json:"InfoItem"`
	}

	if err = json.Unmarshal(result, &resp); err != nil {
		return nil, err
	}

	return resp.InfoItem, nil
}

// https://cloud.tencent.com/document/product/269/3718
// 添加黑名单
// 添加黑名单，支持批量添加黑名单。
// From_Account	String	必填	请求为该 UserID 添加黑名单
// To_Account	Array	必填	待添加为黑名单的用户 UserID 列表，单次请求的 To_Account 数不得超过1000
func (s IMServer) BlackListAdd(fromAccount string, toAccount []string) ([]FriendResultItem, error) {
	var req struct {
		FromAccount string   `json:"From_Account"`
		ToAccount   []string `json:"To_Account"`
	}
	req.FromAccount = fromAccount
	req.ToAccount = toAccount
	result, err := s.requestWithPath(BLACK_LIST_ADD, req)
	if err != nil {
		return nil, err
	}
	var resp struct {
		ResultItem []FriendResultItem `json:"ResultItem"`
	}

	if err = json.Unmarshal(result, &resp); err != nil {
		return nil, err
	}

	return resp.ResultItem, nil
}

// https://cloud.tencent.com/document/product/269/3719
// 删除黑名单
// From_Account	String	必填	需要删除该 UserID 的黑名单
// To_Account	Array	必填	待删除的黑名单的 UserID 列表，单次请求的 To_Account 数不得超过1000
func (s IMServer) BlackListDelete(fromAccount string, toAccount []string) ([]FriendResultItem, error) {
	var req struct {
		FromAccount string   `json:"From_Account"`
		ToAccount   []string `json:"To_Account"`
	}
	req.FromAccount = fromAccount
	req.ToAccount = toAccount
	result, err := s.requestWithPath(BLACK_LIST_DELETE, req)
	if err != nil {
		return nil, err
	}
	var resp struct {
		ResultItem []FriendResultItem `json:"ResultItem"`
	}

	if err = json.Unmarshal(result, &resp); err != nil {
		return nil, err
	}

	return resp.ResultItem, nil
}

// https://cloud.tencent.com/document/product/269/3722
// 拉取黑名单
// 支持分页拉取所有黑名单。
func (s IMServer) BlackListGet(b BlackList) (BlackListResult, error) {
	result, err := s.requestWithPath(BLACK_LIST_GET, b)
	if err != nil {
		return BlackListResult{}, err
	}
	var resp BlackListResult
	if err = json.Unmarshal(result, &resp); err != nil {
		return BlackListResult{}, err
	}

	return resp, nil
}

// https://cloud.tencent.com/document/product/269/3725
// 校验黑名单
// 支持批量校验黑名单。
// From_Account	String	必填	需要校验该 UserID 的黑名单
// To_Account	Array	必填	待校验的黑名单的 UserID 列表，单次请求的 To_Account 数不得超过1000
// CheckType	String	必填	校验模式，详情可参见 校验黑名单
func (s IMServer) BlackListCheck(fromAccount string, toAccount []string, checkType BlackListCheckType) ([]CheckResult, error) {

	req := struct {
		FromAccount string             `json:"From_Account"`
		ToAccount   []string           `json:"To_Account"`
		CheckType   BlackListCheckType `json:"CheckType"`
	}{FromAccount: fromAccount, CheckType: checkType, ToAccount: toAccount}

	result, err := s.requestWithPath(BLACK_LIST_CHECK, req)
	if err != nil {
		return nil, err
	}

	var resp struct {
		BlackListCheckItem []CheckResult `json:"BlackListCheckItem"`
	}

	if err = json.Unmarshal(result, &resp); err != nil {
		return nil, err
	}

	return resp.BlackListCheckItem, nil
}

// 添加分组
// 添加分组，支持批量添加分组，并将指定好友加入到新增分组中。
// From_Account	String	必填	需要为该 UserID 添加新分组
// GroupName	Array	必填	新增分组列表
// To_Account	Array	选填	需要加入新增分组的好友的 UserID 列表
func (s IMServer) GroupAdd(fromAccount string, groupName, toAccount []string) ([]FriendResultItem, error) {

	req := struct {
		FromAccount string   `json:"From_Account"`
		GroupName   []string `json:"GroupName"`
		ToAccount   []string `json:"To_Account"`
	}{FromAccount: fromAccount, GroupName: groupName, ToAccount: toAccount}
	result, err := s.requestWithPath(GROUP_ADD, req)
	if err != nil {
		return nil, err
	}

	var resp struct {
		ResultItem []FriendResultItem `json:"ResultItem"`
	}

	if err = json.Unmarshal(result, &resp); err != nil {
		return nil, err
	}

	return resp.ResultItem, nil
}

// 删除分组
// 删除指定分组。
// From_Account	String	必填	需要删除该 UserID 的分组
// GroupName	Array	必填	要删除的分组列表
func (s IMServer) GroupDelete(fromAccount string, groupName []string) error {

	req := struct {
		FromAccount string   `json:"From_Account"`
		GroupName   []string `json:"GroupName"`
	}{FromAccount: fromAccount, GroupName: groupName}

	if _, err := s.requestWithPath(GROUP_DELETE, req); err != nil {
		return err
	}

	return nil
}
