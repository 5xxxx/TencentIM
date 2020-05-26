/*
 *
 * group_manage.go
 * server
 *
 * Created by lintao on 2020/4/17 1:39 下午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package TencentIM

import "encoding/json"

// https://cloud.tencent.com/document/product/269/1614
// 获取 App 中的所有群组
// App 管理员可以通过该接口获取 App 中所有群组的 ID。
// 如果 App 中的总群数量超过10000个，最多只会返回10000个群组 ID（如果需要完整获取，必须使用分页拉取的形式）。
// 可以使用 Limit 和 Next 两个值用于控制分页拉取：
// Limit 限制回包中 GroupIdList 中群组的个数，不得超过10000。
// Next 控制分页。对于分页请求，第一次填0，后面的请求填上一次返回的 Next 字段，当返回的 Next 为0，代表所有的群都已拉取到。
// 例如：假设需要分页拉取，每页展示 20 个，则第一页的请求参数应当为{“Limit” : 20, “Next” : 0}，第二页的请求参数应当为{“Limit” : 20, “Next” : 上次返回的Next字段}，依此类推。
func (s IMServer) GetAllGroup(limit, next int, groupType GroupType) (GroupList, error) {

	req := struct {
		Limit     int       `json:"Limit"`
		Next      int       `json:"Next"`
		GroupType GroupType `json:"GroupType"`
	}{Limit: limit, Next: next, GroupType: groupType}

	result, err := s.requestWithPath(GET_APPID_GROUP_LIST, req)
	if err != nil {
		return GroupList{}, err
	}

	var resp GroupList
	if err = json.Unmarshal(result, &resp); err != nil {
		return GroupList{}, nil
	}

	return resp, nil
}

// https://cloud.tencent.com/document/product/269/1615
// 创建群组
// App 管理员可以通过该接口创建群组。
func (s IMServer) CreateGroup(g Group) (string, error) {
	result, err := s.requestWithPath(CREATE_GROUP, g)
	if err != nil {
		return "", err
	}

	var resp struct {
		GroupID string `json:"GroupId"`
	}

	if err = json.Unmarshal(result, &resp); err != nil {
		return "", err
	}

	return resp.GroupID, nil
}

// https://cloud.tencent.com/document/product/269/1616
// 获取群组详细资料
// App 管理员可以根据群组 ID 获取群组的详细信息。
func (s IMServer) GetGroupInfo(g GroupInfo) ([]Group, error) {
	reslut, err := s.requestWithPath(GET_GROUP_INFO, g)
	if err != nil {
		return nil, err
	}
	var resp struct {
		GroupInfo []Group `json:"GroupInfo"`
	}

	if err = json.Unmarshal(reslut, &resp); err != nil {
		return nil, err
	}

	return resp.GroupInfo, nil
}

// https://cloud.tencent.com/document/product/269/1617
// 获取群组成员详细资料
// App 管理员可以根据群组 ID 获取群组成员的资料。
// 可以使用 Limit 和 Offset 两个值用于控制分页拉取：
// Limit 限制回包中 MemberList 数组中成员的个数，不得超过6000。
// Offset 控制从群成员中的第多少个成员开始拉取信息。对于分页请求（页码数字从1开始），每一页的 Offset 值应当为：（页码数– 1）×每页展示的群成员数量。
// 例如：假设需要分页拉取，每页展示 20 个，则第一页的请求参数应当为：{“Limit” : 20, “Offset” : 0}，第二页的请求参数应当为{“Limit” : 20, “Offset” : 20}，依此类推。
func (s IMServer) GetGroupMemberInfo(g GroupMemberInfoReq) (GroupMemberInfo, error) {
	result, err := s.requestWithPath(GET_GROUP_MEMBER_INFO, g)
	if err != nil {
		return GroupMemberInfo{}, err
	}

	var resp GroupMemberInfo
	if err = json.Unmarshal(result, &resp); err != nil {
		return GroupMemberInfo{}, err
	}

	return resp, nil
}

// https://cloud.tencent.com/document/product/269/1620
// 修改群组基础资料
// App 管理员可以通过该接口修改指定群组的基础信息。
// GroupId	String	必填	需要修改基础信息的群组的 ID
//Name	String	选填	群名称，最长30字节
//Introduction	String	选填	群简介，最长240字节
//Notification	String	选填	群公告，最长300字节
//FaceUrl	String	选填	群头像 URL，最长100字节
//MaxMemberNum	Integer	选填	最大群成员数量，最大为6000
//ApplyJoinOption	String	选填	申请加群处理方式。包含 FreeAccess（自由加入），NeedPermission（需要验证），DisableApply（禁止加群）
func (s IMServer) ModifyGroupBaseInfo(m ModifyGroup) error {
	if _, err := s.requestWithPath(MODIFY_GROUP_BASE_INFO, m); err != nil {
		return err
	}
	return nil
}

// https://cloud.tencent.com/document/product/269/1621
// 增加群组成员
// App 管理员可以通过该接口向指定的群中添加新成员。
// GroupId	        String		操作的群 ID
// Silence	        Integer		是否静默加人。0：非静默加人；1：静默加人。不填该字段默认为0
// MemberList	    Array		待添加的群成员数组
// Member_Account	String		待添加的群成员 UserID
func (s IMServer) AddGroupMember(groupId string, slilence int, memberList []MemberList) ([]MemberList, error) {
	req := struct {
		GroupID    string       `json:"GroupId"`
		Silence    int          `json:"Silence"`
		MemberList []MemberList `json:"MemberList"`
	}{
		GroupID:    groupId,
		Silence:    slilence,
		MemberList: memberList,
	}

	var resp struct {
		MemberList []MemberList `json:"MemberList" `
	}

	result, err := s.requestWithPath(ADD_GROUP_MEMBER, req)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(result, &resp); err != nil {
		return nil, err
	}

	return resp.MemberList, nil

}

// https://cloud.tencent.com/document/product/269/1622
// 删除群组成员
// App 管理员可以通过该接口删除群成员。
// GroupId	String	必填	操作的群 ID
// Silence	Integer	选填	是否静默删人。0表示非静默删人，1表示静默删人。静默即删除成员时不通知群里所有成员，只通知被删除群成员。不填写该字段时默认为0
// Reason	String	选填	踢出用户原因
// MemberToDel_Account	Array	必填	待删除的群成员
func (s IMServer) DeleteGroupMember(groupID, reason string, silence int, memberToDelAccount []string) error {
	req := struct {
		GroupID            string   `json:"GroupId"`
		Reason             string   `json:"Reason"`
		Silence            int      `json:"Silence"`
		MemberToDelAccount []string `json:"MemberToDel_Account"`
	}{
		GroupID:            groupID,
		Reason:             reason,
		Silence:            silence,
		MemberToDelAccount: memberToDelAccount,
	}
	if _, err := s.requestWithPath(DELETE_GROUP_MEMBER, req); err != nil {
		return err
	}

	return nil
}

// https://cloud.tencent.com/document/product/269/1623
// 修改群成员资料
// App 管理员可以通过该接口修改群成员资料。
func (s IMServer) ModifyGroupMemberInfo(m ModifyGroupMemberInfo) error {
	if _, err := s.requestWithPath(MODIFY_GROUP_MEMBER_INFO, m); err != nil {
		return err
	}

	return nil
}

// https://cloud.tencent.com/document/product/269/1624
// 解散群组
// App 管理员通过该接口解散群。
// GroupId	String	必填	操作的群 ID
func (s IMServer) DestroyGroup(groupId string) error {
	if _, err := s.requestWithPath(DESTROY_GROUP, []byte(`{"GroupId": "`+groupId+`"}`)); err != nil {
		return err
	}

	return nil
}

// https://cloud.tencent.com/document/product/269/1625
// 获取用户所加入的群组
// App 管理员可以通过本接口获取某一用户加入的公开群、聊天室、已激活的私有群信息。默认不获取用户加入的未激活私有群以及音视频聊天室和在线成员广播大群信息。
func (s IMServer) GetJoinedGroupList(j JoinGroupList) (GetJoinedGroupList, error) {
	result, err := s.requestWithPath(GET_JOINED_GROUP_LIST, j)
	if err != nil {
		return GetJoinedGroupList{}, err
	}

	var resp GetJoinedGroupList
	if err = json.Unmarshal(result, &resp); err != nil {
		return GetJoinedGroupList{}, err
	}

	return resp, nil
}

// https://cloud.tencent.com/document/product/269/1626
// 查询用户在群组中的身份
// GroupId	String	必填	需要查询的群组 ID
// User_Account	Array	必填	表示需要查询的用户帐号，最多支持500个帐号
// UserIdList	Array	拉取到的成员在群内的身份信息，可能的身份包括：Owner：群主，Admin：群管理员，Member：群成员，NotMember：非群成员
func (s IMServer) GetRoleInGroup(groupID string, userAccount []string) ([]MemberList, error) {
	req := struct {
		GroupID     string   `json:"GroupId"`
		UserAccount []string `json:"User_Account"`
	}{GroupID: groupID, UserAccount: userAccount}

	result, err := s.requestWithPath(GET_ROLE_IN_GROUP, req)
	if err != nil {
		return nil, err
	}

	var resp struct {
		MemberList []MemberList `json:"UserIdList" `
	}

	if err = json.Unmarshal(result, &resp); err != nil {
		return nil, err
	}

	return resp.MemberList, nil
}

// https://cloud.tencent.com/document/product/269/1627
// 批量禁言和取消禁言
// App 管理员禁止指定群组中某些用户在一段时间内发言。
// App 管理员取消对某些用户的禁言。
// 被禁言用户退出群组之后再进入同一群组，禁言仍然有效
// GroupId		  必填	需要查询的群组 ID
// Members_Account 必填	需要禁言的用户帐号，最多支持500个帐号
// ShutUpTime	  必填	需禁言时间，单位为秒，为0时表示取消禁言
func (s IMServer) ForbidSendMsg(groupID string, membersAccount []string, shutUpTime int) error {
	req := struct {
		GroupID        string   `json:"GroupId"`
		MembersAccount []string `json:"Members_Account"`
		ShutUpTime     int      `json:"ShutUpTime"`
	}{
		GroupID:        groupID,
		MembersAccount: membersAccount,
		ShutUpTime:     shutUpTime,
	}

	if _, err := s.requestWithPath(FORBID_SEND_MSG, req); err != nil {
		return err
	}
	return nil
}

// https://cloud.tencent.com/document/product/269/2925
// 获取群组被禁言用户列表
// App 管理员可以根据群组 ID 获取群组中被禁言的用户列表。
// GroupId	String	必填	需要获取被禁言成员列表的群组 ID。
// ShuttedUinList
// 返回结果为禁言用户信息数组，内容包括被禁言的成员 ID，及其被禁言到的时间（使用 UTC 时间，即世界协调时间）
func (s IMServer) GetGroupShuttedUin(groupId string) ([]ShuttedUinList, error) {

	result, err := s.requestWithPath(GET_GROUP_SHUTTED_UIN, []byte(`{
     "GroupId":"`+groupId+`"
		}`))
	if err != nil {
		return nil, err
	}

	var resp struct {
		ShuttedUinList []ShuttedUinList `json:"ShuttedUinList"`
	}

	if err = json.Unmarshal(result, &resp); err != nil {
		return nil, err
	}

	return resp.ShuttedUinList, nil
}

// https://cloud.tencent.com/document/product/269/1629
// 在群组中发送普通消息
// App 管理员可以通过该接口在群组中发送普通消息
func (s IMServer) SendGroupMsg(m GroupMsg) error {

	if _, err := s.requestWithPath(SEND_GROUP_MSG, m); err != nil {
		return err
	}

	return nil
}

// https://cloud.tencent.com/document/product/269/1630
// 在群组中发送系统通知
// GroupId	String	必填	向哪个群组发送系统通知
// ToMembers_Account	Array	选填	接收者群成员列表，请填写接收者 UserID，不填或为空表示全员下发
// Content	String	必填	系统通知的内容
func (s IMServer) SendGroupSystemNotification(groupID string, content string, toMembersAccount []string) error {
	req := struct {
		GroupID          string   `json:"GroupId"`
		ToMembersAccount []string `json:"ToMembers_Account"`
		Content          string   `json:"Content"`
	}{
		GroupID:          groupID,
		Content:          content,
		ToMembersAccount: toMembersAccount,
	}

	if _, err := s.requestWithPath(SEND_GROUP_SYSTEM_NOTIFICATION, req); err != nil {
		return err
	}
	return nil
}

// https://cloud.tencent.com/document/product/269/12341
// 撤回群组消息
// GroupId		必填	操作的群 ID
// MsgSeqList	必填	被撤回的消息 seq 列表，一次请求最多可以撤回10条消息 seq
// MsgSeq		必填	请求撤回的消息 seq
func (s IMServer) GroupMsgRecall(groupID string, msgSeqList []MsgSeqList) ([]RecallRetList, error) {
	req := struct {
		GroupID    string       `json:"GroupId"`
		MsgSeqList []MsgSeqList `json:"MsgSeqList"`
	}{GroupID: groupID, MsgSeqList: msgSeqList}

	result, err := s.requestWithPath(GROUP_MSG_RECALL, req)
	if err != nil {
		return nil, err
	}

	var resp struct {
		RecallRetList []RecallRetList `json:"RecallRetList"`
	}

	if err = json.Unmarshal(result, &resp); err != nil {
		return nil, err
	}

	return resp.RecallRetList, nil
}

// https://cloud.tencent.com/document/product/269/1633
// 转让群主
// App 管理员可以通过该接口将群主转移给他人。
// 没有群主的群也可以转让，新群主不能为空。
// GroupId				必填	要被转移的群组 ID
// NewOwner_Account		必填	新群主 ID
func (s IMServer) ChangeGroupOwner(groupId string, newOwnerAcoount string) error {
	req := struct {
		GroupID         string `json:"GroupId"`
		NewOwnerAccount string `json:"NewOwner_Account"`
	}{GroupID: groupId, NewOwnerAccount: newOwnerAcoount}

	if _, err := s.requestWithPath(CHANGE_GROUP_OWNER, req); err != nil {
		return err
	}

	return nil
}

// https://cloud.tencent.com/document/product/269/1634
// 导入群基础资料
// App 管理员可以通过该接口导入群组，不会触发回调、不会下发通知；当 App 需要从其他即时通信系统迁移到即时通信 IM 时，使用该协议导入存量群组数据
func (s IMServer) ImportGroup(g Group) (string, error) {
	result, err := s.requestWithPath(IMPORT_GROUP, g)
	if err != nil {
		return "", err
	}
	var resp struct {
		GroupID string `json:"GroupId"`
	}

	if err = json.Unmarshal(result, &resp); err != nil {
		return "", err
	}

	return resp.GroupID, nil
}

// https://cloud.tencent.com/document/product/269/1635
// 导入群消息
// 该 API 接口的作用是导入群组的消息，不会触发回调、不会下发通知。
// 当 App 需要从其他即时通信系统迁移到即时通信 IM 时，使用该协议导入存量群消息数据
func (s IMServer) ImportGroupMsg(groupId string, msgs []MsgList) ([]ImportMsgResult, error) {
	req := struct {
		GroupID string    `json:"GroupId"`
		MsgList []MsgList `json:"MsgList"`
	}{
		GroupID: groupId,
		MsgList: msgs,
	}
	result, err := s.requestWithPath(IMPORT_GROUP_MSG, req)
	if err != nil {
		return nil, err
	}

	var resp struct {
		ImportMsgResult []ImportMsgResult `json:"ImportMsgResult"`
	}
	if err = json.Unmarshal(result, &resp); err != nil {
		return nil, err
	}

	return resp.ImportMsgResult, err
}

// https://cloud.tencent.com/document/product/269/1636
// 导入群成员
// 该 API 接口的作用是导入群组成员，不会触发回调、不会下发通知。
// 当 App 需要从其他即时通信系统迁移到即时通信 IM 时，使用该协议导入存量群成员数据
// GroupId			必填	操作的群 ID
// MemberList		必填	待添加的群成员数组。
func (s IMServer) ImportGroupMember(groupId string, list []MemberList) ([]MemberList, error) {
	req := struct {
		GroupID    string       `json:"GroupId"`
		MemberList []MemberList `json:"MemberList"`
	}{GroupID: groupId, MemberList: list}

	result, err := s.requestWithPath(IMPORT_GROUP_MEMBER, req)
	if err != nil {
		return nil, err
	}

	var resp struct {
		MemberList []MemberList `json:"MemberList"`
	}

	if err = json.Unmarshal(result, &resp); err != nil {
		return nil, err
	}

	return resp.MemberList, nil
}

// https://cloud.tencent.com/document/product/269/1637
// 设置成员未读消息计数
// App 管理员使用该接口设置群组成员未读消息数，不会触发回调、不会下发通知。
// 当 App 需要从其他即时通信系统迁移到即时通信 IM 时，使用该协议设置群成员的未读消息计数
// GroupId			必填	操作的群 ID
// Member_Account	必填	要操作的群成员
// UnreadMsgNum		必填	成员未读消息数
func (s IMServer) SetUnreadMsgNum(group, memberAccount string, unReadMsgNum int) error {

	req := struct {
		GroupID       string `json:"GroupId"`
		MemberAccount string `json:"Member_Account"`
		UnreadMsgNum  int    `json:"UnreadMsgNum"`
	}{
		GroupID:       group,
		MemberAccount: memberAccount,
		UnreadMsgNum:  unReadMsgNum,
	}

	if _, err := s.requestWithPath(SET_UNREAD_MSG_NUM, req); err != nil {
		return err
	}

	return nil

}

// https://cloud.tencent.com/document/product/269/2359
// 删除指定用户发送的消息
// 该 API 接口的作用是删除最近1000条消息中指定用户发送的消息。
// GroupId				必填	要删除消息的群 ID
// Sender_Account		必填	被删除消息的发送者 ID
func (s IMServer) DeleteGroupMsgBySender(groupId string, senderAccount string) error {
	req := struct {
		GroupID       string `json:"GroupId"`
		SenderAccount string `json:"Sender_Account"`
	}{GroupID: groupId, SenderAccount: senderAccount}

	if _, err := s.requestWithPath(DELETE_GROUP_MSG_BY_SENDER, req); err != nil {
		return err
	}

	return nil
}

// https://cloud.tencent.com/document/product/269/2738
// 拉取群漫游消息
// 即时通信 IM 的群消息是按 Seq 排序的，按照 server 收到群消息的顺序分配 Seq，先发的群消息 Seq 小，后发的 Seq 大。
// 如果用户想拉取一个群的全量消息，首次拉取时不用填拉取 Seq，Server 会自动返回最新的消息，以后拉取时拉取 Seq 填上次返回的最小 Seq 减1。
// 如果返回消息的 IsPlaceMsg 为1，表示这个 Seq 的消息或者过期、或者存储失败、或者被删除了。
//  GroupId				必填	要拉取漫游消息的群组 ID
//	ReqMsgNumber		必填	拉取的漫游消息的条数，目前一次请求最多返回20条漫游消息，所以这里最好小于等于20
//	ReqMsgSeq			选填	拉取消息的最大 seq
func (s IMServer) GroupMsgGetSimple(groupId string, reqMsgSeq int, reqMsgNumber int) (RoamingMessage, error) {
	req := struct {
		GroupID      string `json:"GroupId"`
		ReqMsgSeq    int    `json:"ReqMsgSeq"`
		ReqMsgNumber int    `json:"ReqMsgNumber"`
	}{GroupID: groupId, ReqMsgNumber: reqMsgNumber, ReqMsgSeq: reqMsgSeq}

	result, err := s.requestWithPath(GROUP_MSG_GET_SIMPLE, req)
	if err != nil {
		return RoamingMessage{}, err
	}

	var resp RoamingMessage
	if err = json.Unmarshal(result, &resp); err != nil {
		return RoamingMessage{}, err
	}

	return resp, nil
}
