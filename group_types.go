/*
 *
 * group_types.go
 * TencentIM
 *
 * Created by lintao on 2020/4/22 10:43 上午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package TencentIM

//如果仅需要返回特定群组形态的群组，可以通过 GroupType 进行过滤，但此时返回的 TotalCount 的含义就变成了 App 中属于该群组形态的群组总数。不填为获取所有类型的群组。
//群组形态包括 Public（公开群），Private（私密群），ChatRoom（聊天室），AVChatRoom（音视频聊天室）和 BChatRoom（在线成员广播大群）
type GroupType string

const (
	PublicGroup   GroupType = "Public"
	PrivateGroup  GroupType = "Private"
	ChatRoomGroup GroupType = "ChatRoom"
	AVChatRoom    GroupType = "AVChatRoom"
	BChatRoom     GroupType = "BChatRoom"
)

//Owner_Account	String	选填	群主 ID，自动添加到群成员中。如果不填，群没有群主
//Type	String	必填	群组形态，包括 Public（公开群），Private（私密群），ChatRoom（聊天室），AVChatRoom（音视频聊天室），BChatRoom（在线成员广播大群）
//GroupId	String	选填	为了使得群组 ID 更加简单，便于记忆传播，腾讯云支持 App 在通过 REST API 创建群组时 自定义群组 ID
//Name	String	必填	群名称，最长30字节，使用 UTF-8 编码，1个汉字占3个字节
//Introduction	String	选填	群简介，最长240字节，使用 UTF-8 编码，1个汉字占3个字节
//Notification	String	选填	群公告，最长300字节，使用 UTF-8 编码，1个汉字占3个字节
//FaceUrl	String	选填	群头像 URL，最长100字节
//MaxMemberCount	Integer	选填	最大群成员数量，缺省时的默认值：私有群是200，公开群是2000，聊天室是6000，音视频聊天室和在线成员广播大群无限制
//ApplyJoinOption	String	选填	申请加群处理方式。包含 FreeAccess（自由加入），NeedPermission（需要验证），DisableApply（禁止加群），不填默认为 NeedPermission（需要验证）
//仅当创建支持申请加群的 群组 时，该字段有效
//AppDefinedData	Array	选填	群组维度的自定义字段，默认情况是没有的，需要开通，详情请参阅 自定义字段
//MemberList	Array	选填	初始群成员列表，最多500个；成员信息字段详情请参阅 群成员资料
//AppMemberDefinedData	Array	选填	群成员维度的自定义字段，默认情况是没有的，需要开通，详情请参阅 自定义字段
type Group struct {
	OwnerAccount    string           `json:"Owner_Account"`
	Type            string           `json:"Type"`
	GroupID         string           `json:"GroupId"`
	Name            string           `json:"Name"`
	Introduction    string           `json:"Introduction"`
	Notification    string           `json:"Notification"`
	FaceURL         string           `json:"FaceUrl"`
	MaxMemberCount  int              `json:"MaxMemberCount"`
	CreateTime      int              `json:"CreateTime"`
	ApplyJoinOption string           `json:"ApplyJoinOption"`
	AppDefinedData  []AppDefinedData `json:"AppDefinedData"`
	MemberList      []MemberList     `json:"MemberList"`
	ErrorCode       int              `json:"ErrorCode"`
	ErrorInfo       string           `json:"ErrorInfo"`
}

type MemberList struct {
	MemberAccount        string           `json:"Member_Account"`
	Role                 string           `json:"Role"`
	JoinTime             int64            `json:"JoinTime"`
	MsgSeq               int              `json:"MsgSeq"`
	MsgFlag              string           `json:"MsgFlag"`
	LastSendMsgTime      int              `json:"LastSendMsgTime"`
	ShutUpUntil          int              `json:"ShutUpUntil"`
	AppMemberDefinedData []AppDefinedData `json:"AppMemberDefinedData"`
	Result               int              `json:"Result"`
	UnreadMsgNum         int              `json:"UnreadMsgNum" `
}

type AppDefinedData struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

type GroupId struct {
	GroupID string `json:"GroupId"`
}

type GroupList struct {
	GroupIDList []GroupId `json:"GroupIdList"`
	Next        int64     `json:"Next"`
}

type GroupInfo struct {
	GroupIDList    []string    `json:"GroupIdList"`
	ResponseFilter GroupFilter `json:"ResponseFilter"`
}

type GroupFilter struct {
	GroupBaseInfoFilter             []string `json:"GroupBaseInfoFilter"`
	MemberInfoFilter                []string `json:"MemberInfoFilter"`
	AppDefinedDataFilterGroup       []string `json:"AppDefinedDataFilter_Group"`
	AppDefinedDataFilterGroupMember []string `json:"AppDefinedDataFilter_GroupMember"`
}

type GroupMemberInfoReq struct {
	GroupID                         string   `json:"GroupId"`
	MemberInfoFilter                []string `json:"MemberInfoFilter"`
	MemberRoleFilter                []string `json:"MemberRoleFilter"`
	AppDefinedDataFilterGroupMember []string `json:"AppDefinedDataFilter_GroupMember"`
	Limit                           int      `json:"Limit"`
	Offset                          int      `json:"Offset"`
}

type GroupMemberInfo struct {
	MemberNum  int          `json:"MemberNum"`
	MemberList []MemberList `json:"MemberList"`
}

type ApplyJoinOption string

const (
	ApplyJoinFreeAccess     ApplyJoinOption = "FreeAccess"
	ApplyJoinNeedPermission ApplyJoinOption = "NeedPermission"
	ApplyJoinDisableApply   ApplyJoinOption = "DisableApply"
)

type ModifyGroup struct {
	GroupID         string           `json:"GroupId"`
	Name            string           `json:"Name"`
	Introduction    string           `json:"Introduction"`
	Notification    string           `json:"Notification"`
	FaceURL         string           `json:"FaceUrl"`
	MaxMemberNum    int              `json:"MaxMemberNum"`
	ApplyJoinOption ApplyJoinOption  `json:"ApplyJoinOption"`
	ShutUpAllMember string           `json:"ShutUpAllMember"`
	AppDefinedData  []AppDefinedData `json:"AppDefinedData"`
}

// GroupId	String	必填	操作的群 ID
// Member_Account	String	必填	要操作的群成员
// Role	String	选填	成员身份，Admin/Member 分别为设置/取消管理员
// MsgFlag	String	选填	消息屏蔽类型
// NameCard	String	选填	群名片（最大不超过50个字节）
// AppMemberDefinedData	Array	选填	群成员维度的自定义字段，默认情况是没有的，需要开通，详情请参阅 群组系统
// ShutUpTime	Integer	选填	需禁言时间，单位为秒，0表示取消禁言
type ModifyGroupMemberInfo struct {
	GroupID              string           `json:"GroupId"`
	MemberAccount        string           `json:"Member_Account"`
	ShutUpTime           int              `json:"ShutUpTime"`
	Role                 string           `json:"Role"`
	NameCard             string           `json:"NameCard"`
	MsgFlag              string           `json:"MsgFlag"`
	AppMemberDefinedData []AppDefinedData `json:"AppMemberDefinedData"`
}

//Member_Account	String	必填	需要查询的用户帐号
//WithHugeGroups	Integer	选填	是否获取用户加入的音视频聊天室和在线成员广播大群，0表示不获取，1表示获取。默认为0
//WithNoActiveGroups	Integer	选填	是否获取用户加入的未激活私有群信息，0表示不获取，1表示获取。默认为0
//Limit	Integer	选填	单次拉取的群组数量，如果不填代表所有群组，分页方式与 获取 App 中的所有群组 相同
//Offset	Integer	选填	从第多少个群组开始拉取，分页方式与 获取 App 中的所有群组 相同
//GroupType	String	选填	拉取哪种群组形态，例如 Private，Public，ChatRoom 或 AVChatRoom，不填为拉取所有
type JoinGroupList struct {
	MemberAccount  string         `json:"Member_Account"`
	Limit          int            `json:"Limit"`
	Offset         int            `json:"Offset"`
	GroupType      string         `json:"GroupType"`
	ResponseFilter ResponseFilter `json:"ResponseFilter"`
}

//ResponseFilter	Object	选填	分别包含 GroupBaseInfoFilter 和 SelfInfoFilter
//两个过滤器； GroupBaseInfoFilter 表示需要拉取哪些基础信息字段，详情请参阅 群组系统；
//SelfInfoFilter 表示需要拉取用户在每个群组中的哪些个人资料，详情请参阅 群组系统
// https://cloud.tencent.com/document/product/269/1502
type ResponseFilter struct {
	GroupBaseInfoFilter []string `json:"GroupBaseInfoFilter"`
	SelfInfoFilter      []string `json:"SelfInfoFilter"`
}

type GetJoinedGroupList struct {
	TotalCount  int           `json:"TotalCount"`
	GroupIDList []GroupIDList `json:"GroupIdList"`
}

type GroupIDList struct {
	GroupID     string   `json:"GroupId"`
	Type        string   `json:"Type"`
	LastMsgTime int      `json:"LastMsgTime"`
	MemberCount int      `json:"MemberCount"`
	SelfInfo    SelfInfo `json:"SelfInfo"`
}

type SelfInfo struct {
	Role         string `json:"Role"`
	MsgFlag      string `json:"MsgFlag"`
	UnreadMsgNum int    `json:"UnreadMsgNum"`
}

type ShuttedUinList struct {
	MemberAccount string `json:"Member_Account"`
	ShuttedUntil  int    `json:"ShuttedUntil"`
}

// GroupId	String	必填	向哪个群组发送消息
// Random	Integer	必填	32位随机数。如果5分钟内两条消息的随机值相同，后一条消息将被当做重复消息而丢弃
// MsgPriority	String	选填	消息的优先级
// MsgBody	Array	必填	消息体，详细可参阅 消息格式描述
// From_Account	String	选填	消息来源帐号，选填。如果不填写该字段，则默认消息的发送者为调用该接口时使用的 App 管理员帐号。除此之外，App 亦可通过该字段“伪造”消息的发送者，从而实现一些特殊的功能需求。需要注意的是，如果指定该字段，必须要确保字段中的帐号是存在的
// ForbidCallbackControl	Array	选填	消息回调禁止开关，只对单条消息有效，ForbidBeforeSendMsgCallback 表示禁止发消息前回调，ForbidAfterSendMsgCallback 表示禁止发消息后回调
// OnlineOnlyFlag	Integer	选填	1表示消息仅发送在线成员，默认0表示发送所有成员，音视频聊天室（AVChatRoom）和在线成员广播大群（BChatRoom）不支持该参数
type GroupMsg struct {
	GroupID               string          `json:"GroupId"`
	Random                int             `json:"Random"`
	FromAccount           string          `json:"From_Account"`
	MsgPriority           string          `json:"MsgPriority"`
	OnlineOnlyFlag        int             `json:"OnlineOnlyFlag"`
	ForbidCallbackControl []string        `json:"ForbidCallbackControl"`
	MsgBody               []MsgBody       `json:"MsgBody"`
	OfflinePushInfo       OfflinePushInfo `json:"OfflinePushInfo"`
}

type MsgSeqList struct {
	MsgSeq int `json:"MsgSeq"`
}

type RecallRetList struct {
	MsgSeq  int `json:"MsgSeq"`
	RetCode int `json:"RetCode"`
}

type MsgList struct {
	FromAccount string    `json:"From_Account"`
	SendTime    int64     `json:"SendTime"`
	Random      int       `json:"Random,omitempty"`
	MsgBody     []MsgBody `json:"MsgBody"`
}

type ImportMsgResult struct {
	Result  int   `json:"Result"`
	MsgSeq  int   `json:"MsgSeq"`
	MsgTime int64 `json:"MsgTime"`
}

//From_Account	String	消息的发送者
//IsPlaceMsg	Integer	是否是空洞消息，当消息被删除或者消息过期后，MsgBody 为空，这个字段为1
//MsgRandom	Integer	消息随机值，用来对消息去重，有客户端发消息时填写，如果没有填，服务端会自动生成一个
//MsgSeq	Integer	消息 seq，用来标识唯一消息，值越小发送的越早
//MsgTimeStamp	Integer	消息被发送的时间戳，server 的时间
//MsgBody	Array	消息内容，详情请参见 消息内容 MsgBody 说明
type RspMsgList struct {
	FromAccount  string    `json:"From_Account"`
	IsPlaceMsg   int       `json:"IsPlaceMsg"`
	MsgBody      []MsgBody `json:"MsgBody"`
	MsgRandom    int       `json:"MsgRandom"`
	MsgSeq       int       `json:"MsgSeq"`
	MsgTimeStamp int       `json:"MsgTimeStamp"`
}

//GroupId	String	请求中的群组 ID
//IsFinished	Integer	是否返回了请求区间的全部消息
//当消息长度太长或者区间太大（超过20）导致无法返回全部消息时，值为0
//当消息长度太长或者区间太大（超过20）且所有消息都过期时，值为2
//RspMsgList	Array	返回的消息列表
type RoamingMessage struct {
	GroupID    string       `json:"GroupId"`
	IsFinished int          `json:"IsFinished"`
	RspMsgList []RspMsgList `json:"RspMsgList"`
}
