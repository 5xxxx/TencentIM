/*
 *
 * ban_speak_test.go
 * tencentIM
 *
 * Created by lintao on 2020/4/20 11:07 上午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package TencentIM

import "testing"

func TestIMServer_GetNoSpeaking(t *testing.T) {

	type args struct {
		account string
	}
	tests := []struct {
		name                       string
		args                       args
		wantC2CmsgNospeakingTime   int64
		wantGroupmsgNospeakingTime int64
		wantErr                    bool
	}{
		{
			name:                       "查询禁言",
			args:                       args{account: "BbyrNMnaV3"},
			wantC2CmsgNospeakingTime:   4294967295,
			wantGroupmsgNospeakingTime: 4294967295,
			wantErr:                    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotC2CmsgNospeakingTime, gotGroupmsgNospeakingTime, err := s.GetNoSpeaking(tt.args.account)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNoSpeaking() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotC2CmsgNospeakingTime != tt.wantC2CmsgNospeakingTime {
				t.Errorf("GetNoSpeaking() gotC2CmsgNospeakingTime = %v, want %v", gotC2CmsgNospeakingTime, tt.wantC2CmsgNospeakingTime)
			}
			if gotGroupmsgNospeakingTime != tt.wantGroupmsgNospeakingTime {
				t.Errorf("GetNoSpeaking() gotGroupmsgNospeakingTime = %v, want %v", gotGroupmsgNospeakingTime, tt.wantGroupmsgNospeakingTime)
			}
		})
	}
}

func TestIMServer_SetNoSpeaking(t *testing.T) {

	type args struct {
		account                string
		C2CmsgNospeakingTime   int64
		groupmsgNospeakingTime int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "设置禁言",
			args: args{
				account:                "BbyrNMnaV3",
				C2CmsgNospeakingTime:   4294967295,
				groupmsgNospeakingTime: 4294967295,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.SetNoSpeaking(tt.args.account, tt.args.C2CmsgNospeakingTime, tt.args.groupmsgNospeakingTime); (err != nil) != tt.wantErr {
				t.Errorf("SetNoSpeaking() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
