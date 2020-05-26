/*
 *
 * data_manage_types.go
 * TencentIM
 *
 * Created by lintao on 2020/4/22 2:50 下午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package TencentIM

type PortraitTag string

const (
	// 昵称	长度不得超过500个字节
	ProfileIMNick PortraitTag = "Tag_Profile_IM_Nick"
	//Gender_Type_Unknown：没设置性别 Gender_Type_Female：女性 Gender_Type_Male：男性
	ProfileIMGender PortraitTag = "Tag_Profile_IM_Gender"
	//生日
	ProfileIMBirthDay PortraitTag = "Tag_Profile_IM_BirthDay"
	// 所在地
	//长度不得超过16个字节，推荐用法如下：
	//App 本地定义一套数字到地名的映射关系
	//后台实际保存的是4个 uint32_t 类型的数字
	//其中第一个 uint32_t 表示国家
	//第二个 uint32_t 用于表示省份
	//第三个 uint32_t 用于表示城市
	//第四个 uint32_t 用于表示区县
	ProfileIMLocation PortraitTag = "Tag_Profile_IM_Location"
	//加好友验证方式	有	AllowType_Type_NeedConfirm：需要经过自己确认才能添加自己为好友
	//AllowType_Type_AllowAny：允许任何人添加自己为好友
	//AllowType_Type_DenyAny：不允许任何人添加自己为好友
	ProfileIMAllowType PortraitTag = "Tag_Profile_IM_AllowType"
	//个性签名 长度不得超过500个字节
	ProfileIMSelfSignature PortraitTag = "Tag_Profile_IM_SelfSignature"
	//语言
	ProfileIMLanguage PortraitTag = "Tag_Profile_IM_Language"
	//Tag_Profile_IM_Image	string	头像URL	有	长度不得超过500个字节
	ProfileIMImage PortraitTag = "Tag_Profile_IM_Image"
	//	消息设置	有	标志位：
	//Bit0：置0表示接收消息，置1则不接收消息
	ProfileIMMsgSettings = "Tag_Profile_IM_MsgSettings"
	//管理员禁止加好友标识	AdminForbid_Type_None：默认值，允许加好友
	//AdminForbid_Type_SendOut：禁止该用户发起加好友请求
	ProfileIMAdminForbidType PortraitTag = "Tag_Profile_IM_AdminForbidType"
	//等级	uint32	等级	有	通常一个 UINT-8 数据即可保存一个等级信息
	//您可以考虑拆分保存，从而实现多种角色的等级信息
	ProfileIMLevel PortraitTag = "Tag_Profile_IM_Level"
	//角色 通常一个 UINT-8 数据即可保存一个角色信息
	//您可以考虑拆分保存，从而保存多种角色信息
	ProfileIMRole       PortraitTag = "Tag_Profile_IM_Role"
	ProfileIMCustomTest PortraitTag = "Tag_Profile_Custom_Test"
)

// https://cloud.tencent.com/document/product/269/1500#.E6.A0.87.E9.85.8D.E8.B5.84.E6.96.99.E5.AD.97.E6.AE.B5
//To_Account	Array	必填	需要拉取这些 UserID 的资料；
// 注意：每次拉取的用户数不得超过100，避免因回包数据量太大以致回包失败
//TagList	Array	必填	指定要拉取的资料字段的 Tag，支持的字段有：
//1. 标配资料字段，详情可参见 标配资料字段
//2. 自定义资料字段，详情可参见 自定义资料字段
type PortraitReq struct {
	ToAccount []string      `json:"To_Account"`
	TagList   []PortraitTag `json:"TagList"`
}

type PortraitResp struct {
	UserProfileItem []UserProfileItem `json:"UserProfileItem"`
	FailAccount     []string          `json:"Fail_Account"`
	ActionStatus    string            `json:"ActionStatus"`
	ErrorCode       int               `json:"ErrorCode"`
	ErrorInfo       string            `json:"ErrorInfo"`
	ErrorDisplay    string            `json:"ErrorDisplay"`
}

type UserProfileItem struct {
	ToAccount   string        `json:"To_Account"`
	ProfileItem []ProfileItem `json:"ProfileItem,omitempty"`
	ResultCode  int           `json:"ResultCode"`
	ResultInfo  string        `json:"ResultInfo"`
}

type PortraitSet struct {
	FromAccount string        `json:"From_Account"`
	ProfileItem []ProfileItem `json:"ProfileItem"`
}

type ProfileItem struct {
	Tag   PortraitTag `json:"Tag"`
	Value string      `json:"Value"`
}
