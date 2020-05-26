/*
 *
 * chat_test.go
 * tencentIM
 *
 * Created by lintao on 2020/4/20 2:17 下午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package TencentIM

import (
	"testing"
	"time"
)

func TestIMServer_SendMsg(t *testing.T) {
	localtion := NewLocationElem("cc", 12, 5)
	text := NewTextElem("hello")
	face := NewFaceElem(1, "ha")
	custom := NewCustomElem("a", "b", "c", "d")
	type args struct {
		b SingleChatMsg
	}
	tests := []struct {
		name        string
		args        args
		wantMsgTime int64
		wantMsgKey  string
		wantErr     bool
	}{
		{
			name: "发送消息文本消息",
			args: args{SingleChatMsg{
				ToAccount: "MGYjF3E58N",
				ChatMsg: ChatMsg{
					SyncOtherMachine: 1,
					MsgRandom:        1287657,
					MsgTimeStamp:     time.Now().Unix(),
					MsgBody:          []MsgBody{text},
				},
			}},
			wantMsgKey:  "",
			wantMsgTime: 0,
			wantErr:     false,
		},
		{
			name: "发送位置消息",
			args: args{SingleChatMsg{
				ToAccount: "MGYjF3E58N",
				ChatMsg: ChatMsg{
					SyncOtherMachine: 1,
					MsgRandom:        1287657,
					MsgTimeStamp:     time.Now().Unix(),
					MsgBody:          []MsgBody{localtion},
				},
			}},
			wantMsgKey:  "",
			wantMsgTime: 0,
			wantErr:     false,
		},
		{
			name: "发送表情消息",
			args: args{SingleChatMsg{
				ToAccount: "MGYjF3E58N",
				ChatMsg: ChatMsg{
					SyncOtherMachine: 1,
					MsgRandom:        1287657,
					MsgTimeStamp:     time.Now().Unix(),
					MsgBody:          []MsgBody{face},
				},
			}},
			wantMsgKey:  "",
			wantMsgTime: 0,
			wantErr:     false,
		},
		{
			name: "发送自定义消息消息",
			args: args{SingleChatMsg{
				ToAccount: "MGYjF3E58N",
				ChatMsg: ChatMsg{
					SyncOtherMachine: 1,
					MsgRandom:        1287657,
					MsgTimeStamp:     time.Now().Unix(),
					MsgBody:          []MsgBody{custom},
				},
			}},
			wantMsgKey:  "",
			wantMsgTime: 0,
			wantErr:     false,
		},
		{
			name: "某一帐号向其他帐号发消息，接收方看到发送者不是管理员",
			args: args{SingleChatMsg{
				ToAccount: "nDSdlrIy3q",
				ChatMsg: ChatMsg{
					SyncOtherMachine: 1,
					MsgRandom:        1287657,
					MsgTimeStamp:     time.Now().Unix(),
					MsgBody:          []MsgBody{text},
				},
			}},
			wantMsgKey:  "",
			wantMsgTime: 0,
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := s.SendMsg(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendMsg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if gotMsgTime != tt.wantMsgTime {
			//	t.Errorf("SendMsg() gotMsgTime = %v, want %v", gotMsgTime, tt.wantMsgTime)
			//}
			//if gotMsgKey != tt.wantMsgKey {
			//	t.Errorf("SendMsg() gotMsgKey = %v, want %v", gotMsgKey, tt.wantMsgKey)
			//}
		})
	}
}

func TestIMServer_BATCHChatMsg(t *testing.T) {
	text := NewTextElem("hello")
	type args struct {
		b BatchChatMsg
	}
	tests := []struct {
		name        string
		args        args
		wantMsgKey  string
		wantErrList []ErrorList
		wantErr     bool
	}{
		{
			name: "批量发单聊消息",
			args: args{BatchChatMsg{
				ToAccount: []string{"MGYjF3E58N", "nDSdlrIy3q", "CaF2hpUL4K"},
				ChatMsg: ChatMsg{
					SyncOtherMachine: 20,
					MsgRandom:        1287657,
					MsgTimeStamp:     time.Now().Unix(),
					MsgBody:          []MsgBody{text},
				},
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := s.BATCHChatMsg(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("BATCHChatMsg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if gotMsgKey != tt.wantMsgKey {
			//	t.Errorf("BATCHChatMsg() gotMsgKey = %v, want %v", gotMsgKey, tt.wantMsgKey)
			//}
			//if !reflect.DeepEqual(gotErrList, tt.wantErrList) {
			//	t.Errorf("BATCHChatMsg() gotErrList = %v, want %v", gotErrList, tt.wantErrList)
			//}
		})
	}
}

func TestIMServer_ImportMsg(t *testing.T) {
	text := NewTextElem("hello")
	type args struct {
		b SingleChatMsg
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "导入单聊消息",
			args: args{SingleChatMsg{
				ToAccount: "nDSdlrIy3q",
				ChatMsg: ChatMsg{
					SyncFromOldSystem: 1,
					MsgRandom:         1287657,
					MsgTimeStamp:      time.Now().Unix(),
					MsgBody:           []MsgBody{text},
				},
			}},
			wantErr: false,
		},
		{
			name: "导入单聊消息",
			args: args{SingleChatMsg{
				FromAccount: "CaF2hpUL4K",
				ToAccount:   "nDSdlrIy3q",
				ChatMsg: ChatMsg{
					SyncFromOldSystem: 1,
					MsgRandom:         1287657,
					MsgTimeStamp:      time.Now().Unix(),
					MsgBody:           []MsgBody{text},
				},
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.ImportMsg(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("ImportMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIMServer_GetRoamMsg(t *testing.T) {

	type args struct {
		m RoamMsgReq
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "查询单聊消息",
			args: args{RoamMsgReq{
				FromAccount: "MGYjF3E58N",
				ToAccount:   "nDSdlrIy3q",
				MaxCnt:      100,
				MinTime:     1522048000,
				MaxTime:     1527318400,
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.GetRoamMsg(tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("GetRoamMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIMServer_MsgWithdraw(t *testing.T) {

	type args struct {
		m MsgWithdraw
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{MsgWithdraw{
				FromAccount: identifier,
				ToAccount:   "MGYjF3E58N",
				MsgKey:      "262699204_1287657_1587369373",
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.MsgWithdraw(tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("MsgWithdraw() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
