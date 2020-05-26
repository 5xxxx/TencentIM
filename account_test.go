/*
 *
 * account_test.go
 * server
 *
 * Created by lintao on 2020/4/17 2:53 下午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package TencentIM

import (
	"reflect"
	"testing"
)

const secretKey = "xxxx"
const identifier = "xx"
const appid = 000
const expire = 180 * 86400

var s = imServer()

func TestIMServer_DeleteAccount(t *testing.T) {

	accounts := make([]string, 101)
	for i := 0; i < 100; i++ {
		accounts = append(accounts, "@TLS#NOT_FOUND")
	}
	type args struct {
		accounts []string
	}
	tests := []struct {
		name    string
		args    args
		want    []ResultItem
		wantErr bool
	}{
		{
			name: "批量删除不存在的用户",
			args: args{[]string{"dBG2akQ3VE", "dRJISRqwIv"}},
			want: []ResultItem{{
				ResultCode: 70107,
				ResultInfo: "Err_TLS_PT_Open_Login_Account_Not_Exist",
				UserID:     "hSOHtaKTsW",
			},
				{
					ResultCode: 70107,
					ResultInfo: "Err_TLS_PT_Open_Login_Account_Not_Exist",
					UserID:     "szUcAhUyIS",
				},
			},
			wantErr: false,
		},
		{
			name:    "删除超过100个用户",
			args:    args{accounts: accounts},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.DeleteAccount(tt.args.accounts)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIMServer_AccountImport(t *testing.T) {

	type args struct {
		identifier string
		nick       string
		faceUrl    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "account import",
			args: args{
				identifier: "szUcAhUyIS",
				nick:       "lin",
				faceUrl:    "https://pics4.baidu.com/feed/f31fbe096b63f624880b032148196dfe1a4ca326.jpeg?token=e034964f56c79baf13989204b2b29f73",
			},
			wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.AccountImport(tt.args.identifier, tt.args.nick, tt.args.faceUrl); (err != nil) != tt.wantErr {

				t.Errorf("AccountImport() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIMServer_MultiAccountImports(t *testing.T) {

	accounts := make([]string, 101)
	for i := 0; i < 100; i++ {
		accounts = append(accounts, "hSOHtaKTsW")
	}

	type args struct {
		account []string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name:    "导入多个用户",
			args:    args{account: []string{"hSOHtaKTsW", "szUcAhUyIS"}},
			want:    []string{},
			wantErr: true,
		},
		{
			name:    "导入超过100个用户",
			args:    args{account: accounts},
			want:    []string{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.MultiAccountImports(tt.args.account)
			if (err != nil) != tt.wantErr {
				t.Errorf("MultiaccountImports() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("MultiaccountImports() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestIMServer_AccountCheck(t *testing.T) {
	accounts := make([]string, 101)
	for i := 0; i < 100; i++ {
		accounts = append(accounts, "hSOHtaKTsW")
	}
	type args struct {
		accounts []string
	}
	tests := []struct {
		name    string
		args    args
		want    []ResultItem
		wantErr bool
	}{
		{
			name: "查询帐号",
			args: args{accounts: []string{"hSOHtaKTsW", "szUcAhUyIS"}},
			want: []ResultItem{{
				ResultCode:    0,
				ResultInfo:    "",
				UserID:        "hSOHtaKTsW",
				AccountStatus: "NotImported",
			},
				{
					ResultCode:    0,
					ResultInfo:    "",
					UserID:        "szUcAhUyIS",
					AccountStatus: "NotImported",
				}},
			wantErr: false,
		},
		{
			name:    "查询帐号 超过100个的情况",
			args:    args{accounts: accounts},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.AccountCheck(tt.args.accounts)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountCheck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountCheck() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIMServer_Kick(t *testing.T) {

	type args struct {
		userId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "失效帐号登录态",
			args: args{
				userId: "hSOHtaKTsW",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.Kick(tt.args.userId); (err != nil) != tt.wantErr {
				t.Errorf("Kick() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIMServer_QueryState(t *testing.T) {

	type args struct {
		accounts     []string
		isNeedDetail int
	}
	tests := []struct {
		name    string
		args    args
		want    []ResultItem
		wantErr bool
	}{
		{
			name: "查询用户状态",
			args: args{
				accounts:     []string{"hSOHtaKTsW", "szUcAhUyIS"},
				isNeedDetail: 0,
			},
			want:    nil,
			wantErr: true,
		},

		{
			name: "查询用户状态",
			args: args{
				accounts:     []string{"MGYjF3E58N", "CaF2hpUL4K"},
				isNeedDetail: 0,
			},
			want: []ResultItem{{
				ToAccount: "MGYjF3E58N",
				State:     "Offline",
			}, {
				ToAccount: "CaF2hpUL4K",
				State:     "Offline",
			}},
			wantErr: false,
		},
		{
			name: "查询用户状态",
			args: args{
				accounts:     []string{"MGYjF3E58N", "CaF2hpUL4K"},
				isNeedDetail: 1,
			},
			want: []ResultItem{{
				ToAccount: "MGYjF3E58N",
				State:     "Offline",
			}, {
				ToAccount: "CaF2hpUL4K",
				State:     "Offline",
			}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.QueryState(tt.args.accounts, tt.args.isNeedDetail)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryState() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func imServer() IMServer {
	s, err := NewIMServer(appid, expire, identifier, secretKey)
	if err != nil {
		panic(err)
	}
	return s
}
