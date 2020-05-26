/*
 *
 * relationship_manage_test.go
 * tencentIM
 *
 * Created by lintao on 2020/4/20 4:24 下午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package TencentIM

import (
	"reflect"
	"testing"
	"time"
)

func TestIMServer_FriendAdd(t *testing.T) {

	type args struct {
		f Friend
	}
	tests := []struct {
		name    string
		args    args
		want    []FriendResultItem
		wantErr bool
	}{
		{
			name: "测试添加好友",
			args: args{Friend{
				FromAccount: "CaF2hpUL4K",
				AddFriendItem: []AddFriendItem{{
					ToAccount:  "mIpJNhfj4V", //测试时切换账号
					Remark:     "小明",
					GroupName:  "好基友",
					AddSource:  "AddSource_Type_xxa",
					AddWording: "你好啊",
				}},
				AddType:       "Add_Type_Both",
				ForceAddFlags: 1,
			}},
			wantErr: false,
			want: []FriendResultItem{{
				ToAccount:  "mIpJNhfj4V",
				ResultCode: 0,
				ResultInfo: "",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.FriendAdd(tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("FriendAdd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FriendAdd() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIMServer_FriendGetList(t *testing.T) {

	type args struct {
		fromAccount string
		toAccount   []string
		tagList     []ProfileTag
	}
	tests := []struct {
		name    string
		args    args
		want    []InfoItem
		wantErr bool
	}{
		{
			name: "拉取指定好友",
			args: args{
				fromAccount: "CaF2hpUL4K",
				toAccount:   []string{"mIpJNhfj4V"},
				tagList:     []ProfileTag{TagProfileIMNick},
			},
			wantErr: false,
			want: []InfoItem{{
				ToAccount: "mIpJNhfj4V",
				SnsProfileItem: []TagValue{{
					Tag:   "Tag_Profile_IM_Nick",
					Value: "",
				}},
				ResultCode: 0,
				ResultInfo: "",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.FriendGetList(tt.args.fromAccount, tt.args.toAccount, tt.args.tagList)
			if (err != nil) != tt.wantErr {
				t.Errorf("FriendGetList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FriendGetList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIMServer_FriendImport(t *testing.T) {

	type args struct {
		f ImportFriend
	}
	tests := []struct {
		name    string
		args    args
		want    []FriendResultItem
		wantErr bool
	}{
		{
			name: "导入好友",
			args: args{ImportFriend{
				AddFriendItem: []ImportFriendItem{{
					ToAccount:  "mIpJNhfj4V",
					Remark:     "你好",
					AddSource:  "AddSource_Type_iOS",
					AddWording: "还不知道",
					AddTime:    time.Now().Unix(),
					GroupName:  []string{"同事"},
				}},
				FromAccount: identifier,
			}},
			want: []FriendResultItem{{
				ToAccount:  "mIpJNhfj4V",
				ResultCode: 0,
				ResultInfo: "",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.FriendImport(tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("FriendImport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FriendImport() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIMServer_FriendUpdate(t *testing.T) {

	type args struct {
		f UpdateFriend
	}
	tests := []struct {
		name string

		args    args
		want    []FriendResultItem
		wantErr bool
	}{
		{
			name: "更新好友",
			args: args{UpdateFriend{
				FromAccount: "CaF2hpUL4K",
				UpdateItem: []UpdateItem{{
					ToAccount: "mIpJNhfj4V",
					SnsItem: []TagValue{{
						Tag:   "Tag_SNS_IM_Remark",
						Value: "你好",
					}},
				}},
			}},
			wantErr: false,
			want: []FriendResultItem{
				{
					ToAccount:  "mIpJNhfj4V",
					ResultCode: 0,
					ResultInfo: "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.FriendUpdate(tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("FriendUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FriendUpdate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIMServer_FriendDelete(t *testing.T) {

	type args struct {
		fromAccount string
		toAccount   []string
		deleteType  DeleteType
	}
	tests := []struct {
		name    string
		args    args
		want    []FriendResultItem
		wantErr bool
	}{
		{
			name: "删除好友",
			args: args{
				fromAccount: "CaF2hpUL4K",
				toAccount:   []string{"mIpJNhfj4V"},
				deleteType:  DeleteBoth,
			},
			want: []FriendResultItem{{
				ToAccount:  "mIpJNhfj4V",
				ResultCode: 0,
				ResultInfo: "OK",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.FriendDelete(tt.args.fromAccount, tt.args.toAccount, tt.args.deleteType)
			if (err != nil) != tt.wantErr {
				t.Errorf("FriendDelete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FriendDelete() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIMServer_FriendDeleteAll(t *testing.T) {

	type args struct {
		fromAccount string
		deleteType  DeleteType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "删除全部好友",
			args: args{
				fromAccount: "mIpJNhfj4V",
				deleteType:  DeleteBoth,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.FriendDeleteAll(tt.args.fromAccount, tt.args.deleteType); (err != nil) != tt.wantErr {
				t.Errorf("FriendDeleteAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIMServer_FriendCheck(t *testing.T) {

	type args struct {
		fromAccount string
		toAccount   []string
		checkType   CheckType
	}
	tests := []struct {
		name    string
		args    args
		want    []CheckResult
		wantErr bool
	}{
		{
			name: "好友检测",
			args: args{
				fromAccount: "CaF2hpUL4K",
				toAccount:   []string{"mIpJNhfj4V"},
				checkType:   CheckBoth,
			},
			wantErr: false,
			want: []CheckResult{{
				ToAccount:  "mIpJNhfj4V",
				Relation:   "CheckResult_Type_NoRelation",
				ResultCode: 0,
				ResultInfo: "",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.FriendCheck(tt.args.fromAccount, tt.args.toAccount, tt.args.checkType)
			if (err != nil) != tt.wantErr {
				t.Errorf("FriendCheck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FriendCheck() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIMServer_FriendGet(t *testing.T) {

	type args struct {
		f GetFriend
	}
	tests := []struct {
		name    string
		args    args
		want    GetFriendResult
		wantErr bool
	}{
		{
			name: "拉取好友",
			args: args{GetFriend{
				FromAccount:      "mIpJNhfj4V",
				StartIndex:       0,
				StandardSequence: 0,
				CustomSequence:   0,
			}},
			want: GetFriendResult{
				UserDataItem:     nil,
				StandardSequence: 0,
				CustomSequence:   0,
				FriendNum:        0,
				CompleteFlag:     1,
				NextStartIndex:   0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.FriendGet(tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("FriendGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FriendGet() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIMServer_BlackListAdd(t *testing.T) {

	type args struct {
		fromAccount string
		toAccount   []string
	}
	tests := []struct {
		name string

		args    args
		want    []FriendResultItem
		wantErr bool
	}{
		{
			name: "添加黑名单",
			args: args{
				fromAccount: "CaF2hpUL4K",
				toAccount:   []string{"Alf97JLVzk"},
			},
			want: []FriendResultItem{{
				ToAccount:  "Alf97JLVzk",
				ResultCode: 0,
				ResultInfo: "OK",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.BlackListAdd(tt.args.fromAccount, tt.args.toAccount)
			if (err != nil) != tt.wantErr {
				t.Errorf("BlackListAdd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BlackListAdd() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIMServer_BlackListGet(t *testing.T) {

	type args struct {
		b BlackList
	}
	tests := []struct {
		name    string
		args    args
		want    BlackListResult
		wantErr bool
	}{
		{
			name: "拉取黑名单",
			args: args{BlackList{
				FromAccount:  "CaF2hpUL4K",
				StartIndex:   0,
				MaxLimited:   10,
				LastSequence: 0,
			}},
			want: BlackListResult{
				BlackListItem: []struct {
					ToAccount         string `json:"To_Account"`
					AddBlackTimeStamp int    `json:"AddBlackTimeStamp"`
				}{{
					ToAccount:         "Alf97JLVzk",
					AddBlackTimeStamp: 1587454824,
				}},
				StartIndex:       0,
				CurruentSequence: 11,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := s.BlackListGet(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("BlackListGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("BlackListGet() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestIMServer_BlackListDelete(t *testing.T) {

	type args struct {
		fromAccount string
		toAccount   []string
	}
	tests := []struct {
		name    string
		args    args
		want    []FriendResultItem
		wantErr bool
	}{
		{
			name: "删除黑名单",
			args: args{
				fromAccount: "CaF2hpUL4K",
				toAccount:   []string{"Alf97JLVzk"},
			},
			wantErr: false,
			want: []FriendResultItem{{
				ToAccount:  "Alf97JLVzk",
				ResultCode: 0,
				ResultInfo: "OK",
			}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.BlackListDelete(tt.args.fromAccount, tt.args.toAccount)
			if (err != nil) != tt.wantErr {
				t.Errorf("BlackListDelete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BlackListDelete() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIMServer_BlackListCheck(t *testing.T) {

	type args struct {
		fromAccount string
		toAccount   []string
		checkType   BlackListCheckType
	}
	tests := []struct {
		name string

		args    args
		want    []CheckResult
		wantErr bool
	}{
		{
			name: "校验黑名单",
			args: args{
				fromAccount: "CaF2hpUL4K",
				toAccount:   []string{"Alf97JLVzk"},
				checkType:   BlackCheckResultTypeBoth,
			},
			wantErr: false,
			want: []CheckResult{{
				ToAccount:  "Alf97JLVzk",
				Relation:   "BlackCheckResult_Type_NO",
				ResultCode: 0,
				ResultInfo: "",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.BlackListCheck(tt.args.fromAccount, tt.args.toAccount, tt.args.checkType)
			if (err != nil) != tt.wantErr {
				t.Errorf("BlackListCheck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BlackListCheck() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIMServer_GroupDelete(t *testing.T) {

	type args struct {
		fromAccount string
		groupName   []string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "删除群组",
			args: args{
				fromAccount: "CaF2hpUL4K",
				groupName:   []string{"亲戚"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.GroupDelete(tt.args.fromAccount, tt.args.groupName); (err != nil) != tt.wantErr {
				t.Errorf("GroupDelete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestIMServer_GroupAdd(t *testing.T) {

	type args struct {
		fromAccount string
		groupName   []string
		toAccount   []string
	}
	tests := []struct {
		name    string
		args    args
		want    []FriendResultItem
		wantErr bool
	}{
		{
			name: "添加分组",
			args: args{
				fromAccount: "CaF2hpUL4K",
				groupName:   []string{"亲戚"},
				toAccount:   []string{"Alf97JLVzk"},
			},
			want: []FriendResultItem{{
				ToAccount:  "Alf97JLVzk",
				ResultCode: 30001,
				ResultInfo: "Err_SNS_GroupAdd_Friend_Not_Exist",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.GroupAdd(tt.args.fromAccount, tt.args.groupName, tt.args.toAccount)
			if (err != nil) != tt.wantErr {
				t.Errorf("GroupAdd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupAdd() got = %v, want %v", got, tt.want)
			}
		})
	}
}
