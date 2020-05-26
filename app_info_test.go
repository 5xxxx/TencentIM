/*
 *
 * app_info_test.go
 * tencentIM
 *
 * Created by lintao on 2020/4/20 10:11 上午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package TencentIM

import (
	"reflect"
	"testing"
)

func TestIMServer_GetAppInfo(t *testing.T) {

	type args struct {
		field []AppInfoField
	}
	tests := []struct {
		name    string
		args    args
		want    []Info
		wantErr bool
	}{
		{
			name:    "获取运营数据",
			args:    args{[]AppInfoField{}},
			want:    []Info{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.GetAppInfo(tt.args.field...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAppInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("GetAppInfo() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestIMServer_GetHistory(t *testing.T) {

	type args struct {
		chatType string
		msgTime  string
	}
	tests := []struct {
		name    string
		args    args
		want    []Msg
		wantErr bool
	}{
		{
			name: "下载聊天记录",
			args: args{
				chatType: "C2C",
				msgTime:  "2020040121",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetHistory(tt.args.chatType, tt.args.msgTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHistory() got = %v, want %v", got, tt.want)
			}
		})
	}
}
