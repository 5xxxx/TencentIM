/*
 *
 * group_manage_test.go
 * TencentIM
 *
 * Created by lintao on 2020/4/21 4:13 下午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package TencentIM

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

var groupId string

func TestIMServer_GetAllGroup(t *testing.T) {

	type args struct {
		limit     int
		next      int
		groupType GroupType
	}
	tests := []struct {
		name string

		args    args
		want    GroupList
		wantErr bool
	}{
		{
			name: "获取 App 中的所有群组",
			args: args{
				limit:     10,
				next:      0,
				groupType: PublicGroup,
			},
			wantErr: false,
			want: GroupList{
				GroupIDList: []GroupId{{GroupID: ""}},
				Next:        0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := s.GetAllGroup(tt.args.limit, tt.args.next, tt.args.groupType)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("GetAllGroup() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestIMServer_CreateGroup(t *testing.T) {

	type args struct {
		g Group
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "创建群组",
			args: args{Group{
				OwnerAccount:    identifier,
				Type:            "Public",
				Name:            "测试一下",
				Introduction:    "群简介，最长240字节，使用 UTF-8 编码，1个汉字占3个字节",
				Notification:    "群公告，最长300字节，使用 UTF-8 编码，1个汉字占3个字节",
				FaceURL:         "https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=3028752206,69792269&fm=26&gp=0.jpg",
				MaxMemberCount:  2000,
				ApplyJoinOption: "FreeAccess",
				//AppDefinedData: []AppDefinedData{{
				//	Key:   "",
				//	Value: "",
				//}},
				MemberList: []MemberList{{
					MemberAccount: "MGYjF3E58N",
					Role:          "Admin",
					//AppMemberDefinedData: []AppDefinedData{{
					//	Key:   "MemberDefined1",
					//	Value: "MemberData1",
					//}},
				}},
			}},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			groupId, err = s.CreateGroup(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if got != tt.want {
			//	t.Errorf("CreateGroup() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestIMServer_GetGroupInfo(t *testing.T) {

	type args struct {
		g GroupInfo
	}
	tests := []struct {
		name    string
		args    args
		want    []Group
		wantErr bool
	}{
		{
			name: "获取群组详细资料",
			args: args{GroupInfo{
				GroupIDList: []string{
					"@TGS#2UNN5GMGB",
					"@TGS#2Y7P5GMGP",
				},
				ResponseFilter: GroupFilter{GroupBaseInfoFilter: []string{"Type"}},
			}},
			wantErr: false,
			want: []Group{{
				Type:    "Public",
				GroupID: "@TGS#2UNN5GMGB",
			},
				{
					Type:    "Public",
					GroupID: "@TGS#2Y7P5GMGP",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetGroupInfo(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGroupInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGroupInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIMServer_GetGroupMemberInfo(t *testing.T) {

	type args struct {
		g GroupMemberInfoReq
	}
	tests := []struct {
		name    string
		args    args
		want    GroupMemberInfo
		wantErr bool
	}{
		{
			name: "获取群组成员详细资料",
			args: args{GroupMemberInfoReq{
				GroupID: "@TGS#2Y7P5GMGP",
				MemberInfoFilter: []string{
					"Role",
					"JoinTime",
					"MsgSeq",
					"MsgFlag",
					"LastSendMsgTime",
					"ShutUpUntil",
					"NameCard"},
				MemberRoleFilter: []string{
					"Owner",
					"Member",
				},

				Limit:  100,
				Offset: 0,
			}},
			want: GroupMemberInfo{
				MemberNum: 2,
				MemberList: []MemberList{{
					MemberAccount:   "admin_01",
					Role:            "Owner",
					JoinTime:        1587459310,
					MsgSeq:          0,
					MsgFlag:         "AcceptAndNotify",
					LastSendMsgTime: 0,
					ShutUpUntil:     0,
				}},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetGroupMemberInfo(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGroupMemberInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGroupMemberInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIMServer_ModifyGroupBaseInfo(t *testing.T) {

	type args struct {
		m ModifyGroup
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "修改群组基础资料",
			args: args{ModifyGroup{
				GroupID:         "@TGS#2UNN5GMGB",
				Name:            "嘿嘿嘿",
				Introduction:    "群简介，最长240字节",
				Notification:    "群公告，最长300字节",
				FaceURL:         "",
				MaxMemberNum:    6000,
				ApplyJoinOption: "FreeAccess",
				ShutUpAllMember: "On",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.ModifyGroupBaseInfo(tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("ModifyGroupBaseInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIMServer_AddGroupMember(t *testing.T) {

	type args struct {
		groupId    string
		slilence   int
		memberList []MemberList
	}
	tests := []struct {
		name    string
		args    args
		want    []MemberList
		wantErr bool
	}{
		{
			name: "增加群组成员",
			args: args{
				groupId:  "@TGS#2UNN5GMGB",
				slilence: 0,
				memberList: []MemberList{
					{
						MemberAccount: "nDSdlrIy3q",
					},
					{
						MemberAccount: "CaF2hpUL4K",
					}},
			},
			want: []MemberList{
				{
					MemberAccount: "nDSdlrIy3q",
					Result:        1,
				},
				{
					MemberAccount: "CaF2hpUL4K",
					Result:        1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.AddGroupMember(tt.args.groupId, tt.args.slilence, tt.args.memberList)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddGroupMember() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddGroupMember() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIMServer_ModifyGroupMemberInfo(t *testing.T) {

	type args struct {
		m ModifyGroupMemberInfo
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "修改群成员资料",
			args: args{ModifyGroupMemberInfo{
				GroupID:       "@TGS#2UNN5GMGB",
				MemberAccount: "nDSdlrIy3q",
				ShutUpTime:    86400,
				Role:          "Admin",
				NameCard:      "鲍勃",
				MsgFlag:       "AcceptAndNotify",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.ModifyGroupMemberInfo(tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("ModifyGroupMemberInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIMServer_DeleteGroupMember(t *testing.T) {

	type args struct {
		groupID            string
		reason             string
		silence            int
		memberToDelAccount []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "删除群组成员",
			args: args{
				groupID:            "@TGS#2UNN5GMGB",
				reason:             "kick reason",
				silence:            1,
				memberToDelAccount: []string{"nDSdlrIy3q", "CaF2hpUL4K"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.DeleteGroupMember(tt.args.groupID, tt.args.reason, tt.args.silence, tt.args.memberToDelAccount); (err != nil) != tt.wantErr {
				t.Errorf("DeleteGroupMember() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIMServer_DestroyGroup(t *testing.T) {
	//todo
	//先调用创建方法拿群号,上传代码前注释掉手动获取的群号
	type args struct {
		groupId string
	}
	tests := []struct {
		name string

		args    args
		wantErr bool
	}{
		{
			name:    "解散群组",
			args:    args{groupId: groupId},
			wantErr: false,
		},
		//{
		//	name:    "解散群组",
		//	args:    args{groupId: "@TGS#27YLAJMG2"},
		//	wantErr: false,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.DestroyGroup(tt.args.groupId); (err != nil) != tt.wantErr {
				t.Errorf("DestroyGroup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIMServer_GetJoinedGroupList(t *testing.T) {

	type args struct {
		j JoinGroupList
	}
	tests := []struct {
		name    string
		args    args
		want    GetJoinedGroupList
		wantErr bool
	}{
		{
			name: "获取用户所加入的群组",
			args: args{JoinGroupList{
				MemberAccount: "YkvlM3SHgI",
				Limit:         10,
				Offset:        0,
			}},
			wantErr: false,
			want: GetJoinedGroupList{
				TotalCount: 1,
				GroupIDList: []GroupIDList{
					{
						GroupID: "qxQT077M51",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.GetJoinedGroupList(tt.args.j)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetJoinedGroupList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetJoinedGroupList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIMServer_GetRoleInGroup(t *testing.T) {

	type args struct {
		groupID     string
		userAccount []string
	}
	tests := []struct {
		name    string
		args    args
		want    []MemberList
		wantErr bool
	}{
		{
			name: "查询用户在群组中的身份",
			args: args{
				groupID:     "qxQT077M51",
				userAccount: []string{"mIpJNhfj4V", "YkvlM3SHgI", "Alf97JLVzk"},
			},
			want: []MemberList{
				{
					MemberAccount: "mIpJNhfj4V",
					Role:          "Owner",
				},
				{
					MemberAccount: "YkvlM3SHgI",
					Role:          "Member",
				},
				{
					MemberAccount: "Alf97JLVzk",
					Role:          "Member",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetRoleInGroup(tt.args.groupID, tt.args.userAccount)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRoleInGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRoleInGroup() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIMServer_ForbidSendMsg(t *testing.T) {

	type args struct {
		groupID        string
		membersAccount []string
		shutUpTime     int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "批量禁言",
			args: args{
				groupID:        "qxQT077M51",
				membersAccount: []string{"YkvlM3SHgI", "Alf97JLVzk"},
				shutUpTime:     60,
			},
			wantErr: false,
		},
		{
			name: "批量取消禁言",
			args: args{
				groupID:        "qxQT077M51",
				membersAccount: []string{"YkvlM3SHgI", "Alf97JLVzk"},
				shutUpTime:     0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.ForbidSendMsg(tt.args.groupID, tt.args.membersAccount, tt.args.shutUpTime); (err != nil) != tt.wantErr {
				t.Errorf("ForbidSendMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIMServer_GetGroupShuttedUin(t *testing.T) {

	type args struct {
		groupId string
	}
	tests := []struct {
		name    string
		args    args
		want    []ShuttedUinList
		wantErr bool
	}{
		{
			name:    "获取群组被禁言用户列表",
			args:    args{groupId: "qxQT077M51"},
			want:    []ShuttedUinList{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetGroupShuttedUin(tt.args.groupId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGroupShuttedUin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGroupShuttedUin() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIMServer_SendGroupMsg(t *testing.T) {
	rand.Seed(time.Now().Unix())
	type args struct {
		m GroupMsg
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "在群组中发送普通消息",
			args: args{GroupMsg{
				GroupID:     "qxQT077M51",
				FromAccount: identifier,
				Random:      rand.Intn(50000),
				MsgBody:     []MsgBody{NewTextElem("hello")},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.SendGroupMsg(tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("SendGroupMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIMServer_SendGroupSystemNotification(t *testing.T) {

	type args struct {
		groupID          string
		content          string
		toMembersAccount []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "在群组中发送系统通知",
			args: args{
				groupID:          "qxQT077M51",
				content:          "hello",
				toMembersAccount: []string{"mIpJNhfj4V", "YkvlM3SHgI"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.SendGroupSystemNotification(tt.args.groupID, tt.args.content, tt.args.toMembersAccount); (err != nil) != tt.wantErr {
				t.Errorf("SendGroupSystemNotification() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIMServer_GroupMsgRecall(t *testing.T) {

	type args struct {
		groupID    string
		msgSeqList []MsgSeqList
	}
	tests := []struct {
		name    string
		args    args
		want    []RecallRetList
		wantErr bool
	}{
		{
			name: "撤回群组消息",
			args: args{
				groupID:    "qxQT077M51",
				msgSeqList: []MsgSeqList{{MsgSeq: 482070324}},
			},
			wantErr: false,
			want: []RecallRetList{{
				MsgSeq:  482070324,
				RetCode: 10030,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.GroupMsgRecall(tt.args.groupID, tt.args.msgSeqList)
			if (err != nil) != tt.wantErr {
				t.Errorf("GroupMsgRecall() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupMsgRecall() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIMServer_ChangeGroupOwner(t *testing.T) {

	type args struct {
		groupId         string
		newOwnerAcoount string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "转让群主",
			args: args{
				groupId:         "qxQT077M51",
				newOwnerAcoount: "mIpJNhfj4V",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := s.ChangeGroupOwner(tt.args.groupId, tt.args.newOwnerAcoount); (err != nil) != tt.wantErr {
				t.Errorf("ChangeGroupOwner() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIMServer_ImportGroup(t *testing.T) {
	groupName := fmt.Sprint(time.Now().Unix())
	type args struct {
		g Group
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "导入群基础资料",
			args: args{Group{
				OwnerAccount:    "mIpJNhfj4V",
				Type:            "Public",
				GroupID:         groupName,
				Name:            "调配",
				Introduction:    "This is group Introduction",
				Notification:    "This is group Notification",
				FaceURL:         "http://this.is.face.url",
				MaxMemberCount:  500,
				ApplyJoinOption: "FreeAccess",
			}},
			wantErr: false,
			want:    groupName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.ImportGroup(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("ImportGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ImportGroup() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIMServer_ImportGroupMsg(t *testing.T) {
	sendTime := time.Now().Unix()
	type args struct {
		groupId string
		msgs    []MsgList
	}
	tests := []struct {
		name    string
		args    args
		want    []ImportMsgResult
		wantErr bool
	}{
		{
			name: "导入群消息",
			args: args{
				groupId: "qxQT077M51",
				msgs: []MsgList{{
					FromAccount: "mIpJNhfj4V",
					SendTime:    sendTime,
					Random:      rand.Intn(50000),
					MsgBody:     []MsgBody{NewTextElem("hello")},
				}},
			},
			want: []ImportMsgResult{{
				Result:  0,
				MsgTime: sendTime,
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.ImportGroupMsg(tt.args.groupId, tt.args.msgs)
			if (err != nil) != tt.wantErr {
				t.Errorf("ImportGroupMsg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ImportGroupMsg() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIMServer_ImportGroupMember(t *testing.T) {

	type args struct {
		groupId string
		list    []MemberList
	}
	tests := []struct {
		name    string
		args    args
		want    []MemberList
		wantErr bool
	}{
		{
			name: "导入群成员",
			args: args{
				groupId: "qxQT077M51",
				list: []MemberList{{
					MemberAccount: "mIpJNhfj4V",
					Role:          "Admin",
					JoinTime:      time.Now().Unix(),
					UnreadMsgNum:  5,
					MsgFlag:       "AcceptAndNotify",
				}},
			},
			wantErr: false,
			want:    []MemberList{{MemberAccount: "mIpJNhfj4V", Result: 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.ImportGroupMember(tt.args.groupId, tt.args.list)
			if (err != nil) != tt.wantErr {
				t.Errorf("ImportGroupMember() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ImportGroupMember() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIMServer_SetUnreadMsgNum(t *testing.T) {

	type args struct {
		group         string
		memberAccount string
		unReadMsgNum  int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "设置成员未读消息计数",
			args: args{
				group:         "qxQT077M51",
				memberAccount: "mIpJNhfj4V",
				unReadMsgNum:  50,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.SetUnreadMsgNum(tt.args.group, tt.args.memberAccount, tt.args.unReadMsgNum); (err != nil) != tt.wantErr {
				t.Errorf("SetUnreadMsgNum() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIMServer_DeleteGroupMsgBySender(t *testing.T) {

	type args struct {
		groupId       string
		senderAccount string
	}
	tests := []struct {
		name string

		args    args
		wantErr bool
	}{
		{
			name: "删除指定用户发送的消息",
			args: args{
				groupId:       "qxQT077M51",
				senderAccount: "mIpJNhfj4V",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := s.DeleteGroupMsgBySender(tt.args.groupId, tt.args.senderAccount); (err != nil) != tt.wantErr {
				t.Errorf("DeleteGroupMsgBySender() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIMServer_GroupMsgGetSimple(t *testing.T) {

	type args struct {
		groupId      string
		reqMsgSeq    int
		reqMsgNumber int
	}
	tests := []struct {
		name    string
		args    args
		want    RoamingMessage
		wantErr bool
	}{
		{
			name: "拉取群漫游消息",
			args: args{
				groupId:      "qxQT077M51",
				reqMsgSeq:    7803321,
				reqMsgNumber: 20,
			},
			want: RoamingMessage{
				IsFinished: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.GroupMsgGetSimple(tt.args.groupId, tt.args.reqMsgSeq, tt.args.reqMsgNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("GroupMsgGetSimple() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupMsgGetSimple() got = %v, want %v", got, tt.want)
			}
		})
	}
}
