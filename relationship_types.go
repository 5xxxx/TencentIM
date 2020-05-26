/*
 *
 * relationship_types.go
 * TencentIM
 *
 * Created by lintao on 2020/4/22 2:55 下午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package TencentIM

type Relation string

const (
	//单向模式下:From_Account 的好友表中没有 To_Account，但无法确定 To_Account 的好友表中是否有 From_Account
	//双向模式下:From_Account 的好友表中没有 To_Account，To_Account 的好友表中也没有 From_Account
	FriendNoRelation Relation = "CheckResult_Type_NoRelation"
	//From_Account 单向模式下:的好友表中有 To_Account，但无法确定 To_Account 的好友表中是否有 From_Account
	//From_Account 双向模式下:的好友表中有 To_Account，但 To_Account 的好友表中没有 From_Account
	FriendAWithB Relation = "CheckResult_Type_AWithB"
	//From_Account 的好友表中有 To_Account，To_Account 的好友表中也有 From_Account
	FriendBothWay Relation = "CheckResult_Type_BothWay"
	//From_Account 的好友表中没有 To_Account，但 To_Account 的好友表中有 From_Account
	FriendBWithA Relation = "CheckResult_Type_BWithA"
)

type AddFriendType string

const (
	AddFriendSingle AddFriendType = "Add_Type_Single"
	AddFriendBoth   AddFriendType = "Add_Type_Both"
)

type Friend struct {
	FromAccount   string          `json:"From_Account"`
	AddFriendItem []AddFriendItem `json:"AddFriendItem"`
	AddType       AddFriendType   `json:"AddType"`       //加好友方式（默认双向加好友方式）： Add_Type_Single 表示单向加好友 Add_Type_Both 表示双向加好友
	ForceAddFlags int             `json:"ForceAddFlags"` //管理员强制加好友标记：1表示强制加好友，0表示常规加好友方式
}

type AddFriendItem struct {
	ToAccount string `json:"To_Account"`
	Remark    string `json:"Remark"`
	GroupName string `json:"GroupName"`
	//加好友来源：
	//1. 加好友来源字段包含前缀和关键字两部分；
	//2. 加好友来源字段的前缀是：AddSource_Type_ ；
	//3. 关键字：必须是英文字母，且长度不得超过 8 字节，建议用一个英文单词或该英文单词的缩写；
	//4. 示例：加好友来源的关键字是 Android，则加好友来源字段是：AddSource_Type_Android
	AddSource  string `json:"AddSource"`
	AddWording string `json:"AddWording"`
}

// From_Account	String	必填	需要为该 UserID 添加好友
// AddFriendItem	Array	必填	好友结构体对象
type ImportFriend struct {
	AddFriendItem []ImportFriendItem `json:"AddFriendItem"`
	FromAccount   string             `json:"From_Account"`
}

//To_Account	String	必填	好友的 UserID
//Remark	String	选填	From_Account 对 To_Account 的好友备注，详情可参见 标配好友字段
//RemarkTime	Integer	选填	From_Account 对 To_Account 的好友备注时间
//GroupName	Array	选填	From_Account 对 To_Account 的分组信息，详情可参见 标配好友字段
//AddSource	String	必填	加好友来源字段，详情可参见 标配好友字段
//AddWording	String	选填	From_Account 和 To_Account 形成好友关系时的附言信息，详情可参见 标配好友字段
//AddTime	Integer	选填	From_Account 和 To_Account 形成好友关系的时间
//CustomItem	Array	选填	From_Account 对 To_Account 的自定义好友数据，每一个成员都包含一个 Tag 字段和一个 Value 字段，详情可参见 自定义好友字段
//Tag	String	选填	自定义好友字段的名称，使用前请通过即时通信 IM 控制台 >【应用配置】>【功能配置】申请自定义好友字段
//Value	String/Integer	选填	自定义好友字段的值
type ImportFriendItem struct {
	ToAccount  string   `json:"To_Account"`
	AddTime    int64    `json:"AddTime"`
	Remark     string   `json:"Remark"`
	AddSource  string   `json:"AddSource"`
	GroupName  []string `json:"GroupName"`
	AddWording string   `json:"AddWording"`
}

//To_Account	String	请求添加的好友的 UserID
//ResultCode	Integer	To_Account 的处理结果，0表示成功，非0表示失败，非0取值的详细描述请参见 错误码说明
//ResultInfo	String	To_Account 的错误描述信息，成功时该字段为空
type FriendResultItem struct {
	ToAccount  string `json:"To_Account"`
	ResultCode int    `json:"ResultCode"`
	ResultInfo string `json:"ResultInfo"`
}

type UpdateFriend struct {
	FromAccount string       `json:"From_Account"`
	UpdateItem  []UpdateItem `json:"UpdateItem"`
}

type UpdateItem struct {
	ToAccount string     `json:"To_Account"`
	SnsItem   []TagValue `json:"SnsItem"`
}

//{
//"From_Account":"id",
//"To_Account":["id1","id2","id3"],
//"DeleteType":"Delete_Type_Single"
//}

type DeleteType string

const (
	// 单向删除好友
	// 只将 To_Account 从 From_Account 的好友表中删除，但不会将 From_Account 从 To_Account 的好友表中删除
	DeleteSingle DeleteType = "Delete_Type_Single"
	// 双向删除好友
	// 将 To_Account 从 From_Account 的好友表中删除，同时将 From_Account 从 To_Account 的好友表中删除
	DeleteBoth DeleteType = "Delete_Type_Both"
)

type CheckType string

const (
	// 单向校验好友关系
	// 只会检查 From_Account 的好友表中是否有 To_Account，不会检查 To_Account 的好友表中是否有 From_Account
	CheckSingle CheckType = "CheckResult_Type_Single"
	// 双向校验好友关系
	// 既会检查 From_Account 的好友表中是否有 To_Account，也会检查 To_Account 的好友表中是否有 From_Account
	CheckBoth CheckType = "CheckResult_Type_Both"
)

//To_Account	String	请求校验的用户的 UserID
//Relation	String	校验成功时 To_Account 与 From_Account 之间的黑名单关系，详情可参见 校验黑名单
//ResultCode	Integer	To_Account 的处理结果，0表示成功，非0表示失败，非0取值的详细描述请参见 错误码说明
//ResultInfo	String	To_Account 的错误描述信息，成功时该字段为空
type CheckResult struct {
	ToAccount  string   `json:"To_Account"`
	Relation   Relation `json:"Relation"`
	ResultCode int      `json:"ResultCode"`
	ResultInfo string   `json:"ResultInfo"`
}

//From_Account	String	必填	指定要拉取好友数据的用户的 UserID
//StartIndex	Integer	必填	分页的起始位置
//StandardSequence	Integer	选填	上次拉好友数据时返回的 StandardSequence，如果 StandardSequence 字段的值与后台一致，后台不会返回标配好友数据
//CustomSequence	Integer	选填	上次拉好友数据时返回的 CustomSequence，如果 CustomSequence 字段的值与后台一致，后台不会返回自定义好友数据
type GetFriend struct {
	FromAccount      string `json:"From_Account"`
	StartIndex       int    `json:"StartIndex"`
	StandardSequence int    `json:"StandardSequence"`
	CustomSequence   int    `json:"CustomSequence"`
}

type GetFriendResult struct {
	UserDataItem     []UserDataItem `json:"UserDataItem"`
	StandardSequence int            `json:"StandardSequence"`
	CustomSequence   int            `json:"CustomSequence"`
	FriendNum        int            `json:"FriendNum"`
	CompleteFlag     int            `json:"CompleteFlag"`
	NextStartIndex   int            `json:"NextStartIndex"`
}

type UserDataItem struct {
	ToAccount string     `json:"To_Account"`
	ValueItem []TagValue `json:"ValueItem"`
}

type TagValue struct {
	Tag   string `json:"Tag"`
	Value string `json:"Value"`
}

type ProfileTag string

const (
	//昵称
	//长度不得超过500个字节
	TagProfileIMNick = "Tag_Profile_IM_Nick"
	//性别
	//Gender_Type_Unknown：没设置性别
	//Gender_Type_Female：女性
	//Gender_Type_Male：男性
	TagProfileIMGender = "Tag_Profile_IM_Gender"
	//生日
	//推荐用法：20190419
	TagProfileIMBirthDay = "Tag_Profile_IM_BirthDay"
	//所在地
	//长度不得超过16个字节，推荐用法如下：
	//App 本地定义一套数字到地名的映射关系
	//后台实际保存的是4个 uint32_t 类型的数字
	//其中第一个 uint32_t 表示国家
	//第二个 uint32_t 用于表示省份
	//第三个 uint32_t 用于表示城市
	//第四个 uint32_t 用于表示区县
	TagProfileIMLocation = "Tag_Profile_IM_Location"
	//个性签名
	//长度不得超过500个字节
	TagProfileIMSelfSignature = "Tag_Profile_IM_SelfSignature"
	//加好友验证方式
	//AllowType_Type_NeedConfirm：需要经过自己确认才能添加自己为好友
	//AllowType_Type_AllowAny：允许任何人添加自己为好友
	//AllowType_Type_DenyAny：不允许任何人添加自己为好友
	TagProfileIMAllowType = "Tag_Profile_IM_AllowType"
	//语言
	TagProfileIMLanguage = "Tag_Profile_IM_Language"
	//头像URL
	//长度不得超过500个字节
	TagProfileIMImage = "Tag_Profile_IM_Image"
	//消息设置
	//标志位： Bit0：置0表示接收消息，置1则不接收消息
	TagProfileIMMsgSettings = "Tag_Profile_IM_MsgSettings"
	//管理员禁止加好友标识
	//AdminForbid_Type_None：默认值，允许加好友 AdminForbid_Type_SendOut：禁止该用户发起加好友请求
	TagProfileIMAdminForbidType = "Tag_Profile_IM_AdminForbidType"
	//等级
	//通常一个 UINT-8 数据即可保存一个等级信息 您可以考虑拆分保存，从而实现多种角色的等级信息
	TagProfileIMLevel = "Tag_Profile_IM_Level"
	//角色
	//通常一个 UINT-8 数据即可保存一个角色信息 您可以考虑拆分保存，从而保存多种角色信息
	TagProfileIMRole = "Tag_Profile_IM_Role"
	//好友分组：
	//1. 最多支持 32 个分组；
	//2. 不允许分组名为空；
	//3. 分组名长度不得超过 30 个字节；
	//4. 同一个好友可以有多个不同的分组
	TagSNSIMGroup = "Tag_SNS_IM_Group"
	//好友备注：
	//1. 备注长度最长不得超过 96 个字节
	TagSNSIMRemark = "Tag_SNS_IM_Remark"
	//加好友来源：
	//1. 加好友来源字段包含前缀和关键字两部分；
	//2. 加好友来源字段的前缀是：AddSource_Type_ ；
	//3. 关键字：必须是英文字母，且长度不得超过 8 字节，建议用一个英文单词或该英文单词的缩写；
	//4. 示例：加好友来源的关键字是 Android，则加好友来源字段是：AddSource_Type_Android
	TagSNSIMAddSource = "Tag_SNS_IM_AddSource"
	//加好友附言：
	//1. 加好友附言的长度最长不得超过 256 个字节
	TagSNSIMAddWording = "Tag_SNS_IM_AddWording"
)

//To_Account	String	好友的 UserID
//SnsProfileItem	Array	保存好友数据的数组，数组每一个元素都包含一个 Tag 字段和一个 Value 字段
//Tag	String	字段的名称
type InfoItem struct {
	ToAccount      string     `json:"To_Account"`
	SnsProfileItem []TagValue `json:"SnsProfileItem"`
	ResultCode     int        `json:"ResultCode"`
	ResultInfo     string     `json:"ResultInfo"`
}

//From_Account	String	必填	需要拉取该 UserID 的黑名单
//StartIndex	Integer	必填	拉取的起始位置
//MaxLimited	Integer	必填	每页最多拉取的黑名单数
//LastSequence	Integer	必填	上一次拉黑名单时后台返回给客户端的 Seq，初次拉取时为0
type BlackList struct {
	FromAccount  string `json:"From_Account"`
	StartIndex   int    `json:"StartIndex"`
	MaxLimited   int    `json:"MaxLimited"`
	LastSequence int    `json:"LastSequence"`
}

//BlackListItem	Array	黑名单对象数组，每一个黑名单对象都包括了 To_Account 和 AddBlackTimeStamp
//To_Account	String	黑名单的 UserID
//AddBlackTimeStamp	Integer	添加黑名单的时间
//StartIndex	Integer	下页拉取的起始位置，0表示已拉完
//CurruentSequence	Integer	黑名单最新的 Seq
type BlackListResult struct {
	BlackListItem []struct {
		ToAccount         string `json:"To_Account"`
		AddBlackTimeStamp int    `json:"AddBlackTimeStamp"`
	} `json:"BlackListItem"`
	StartIndex       int `json:"StartIndex"`
	CurruentSequence int `json:"CurruentSequence"`
}

type BlackListCheckType string

//https://cloud.tencent.com/document/product/269/1501#.E6.A0.A1.E9.AA.8C.E9.BB.91.E5.90.8D.E5.8D.95
const (
	//只会检查 From_Account 的黑名单中是否有 To_Account，不会检查 To_Account 的黑名单中是否有 From_Account
	BlackCheckResultTypeSingle BlackListCheckType = "BlackCheckResult_Type_Single"
	//既会检查 From_Account 的黑名单中是否有 To_Account，也会检查 To_Account 的黑名单中是否有 From_Account
	BlackCheckResultTypeBoth BlackListCheckType = "BlackCheckResult_Type_Both"
	//From_Account 的黑名单中有 To_Account，但无法确定 To_Account 的黑名单中是否有 From_Account
	//From_Account 的黑名单中有 To_Account，但 To_Account 的黑名单中没有 From_Account
	BlackCheckResult_Type_AWithB BlackListCheckType = "BlackCheckResult_Type_AWithB"
	//From_Account 的黑名单中没有 To_Account，但无法确定 To_Account 的黑名单中是否有 From_Account
	//From_Account 的黑名单中没有 To_Account，To_Account 的黑名单中也没有 From_Account
	BlackCheckResult_Type_NO BlackListCheckType = "BlackCheckResult_Type_NO"
	//From_Account 的黑名单中有 To_Account，To_Account 的黑名单中也有 From_Account
	BlackCheckResult_Type_BothWay BlackListCheckType = "BlackCheckResult_Type_BothWay"
	//From_Account 的黑名单中没有 To_Account，但 To_Account 的黑名单中有 From_Account
	BlackCheckResult_Type_BWithA BlackListCheckType = "BlackCheckResult_Type_BWithA"
)
