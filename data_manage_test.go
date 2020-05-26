/*
 *
 * data_manage_test.go
 * tencentIM
 *
 * Created by lintao on 2020/4/20 4:05 下午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package TencentIM

import (
	"reflect"
	"testing"
)

func TestIMServer_GetPortrait(t *testing.T) {

	type args struct {
		p PortraitReq
	}
	tests := []struct {
		name     string
		args     args
		wantResp PortraitResp
		wantErr  bool
	}{
		{
			name: "获取用户信息",
			args: args{PortraitReq{
				ToAccount: []string{"CaF2hpUL4K"},
				TagList:   []PortraitTag{ProfileIMNick},
			}},
			wantResp: PortraitResp{
				UserProfileItem: []UserProfileItem{{
					ToAccount:   "CaF2hpUL4K",
					ProfileItem: []ProfileItem{{Tag: "Tag_Profile_IM_Nick", Value: ""}},
					ResultCode:  0,
					ResultInfo:  "",
				}},
				ActionStatus: "OK",
				ErrorCode:    0,
				ErrorInfo:    "",
				ErrorDisplay: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := s.GetPortrait(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPortrait() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("GetPortrait() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestIMServer_SetPortrait(t *testing.T) {

	type args struct {
		p PortraitSet
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "设置用户资料",
			args: args{PortraitSet{
				FromAccount: "Alf97JLVzk",
				ProfileItem: []ProfileItem{{
					Tag:   ProfileIMGender,
					Value: "Gender_Type_Male",
				}},
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.SetPortrait(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("SetPortrait() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
