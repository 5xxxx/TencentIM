/*
 *
 * types.go
 * server
 *
 * Created by lintao on 2020/4/17 10:14 上午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package TencentIM

const (
	//TIMTextElem	文本消息。
	TIMTextElem = "TIMTextElem"
	//TIMLocationElem	地理位置消息。
	TIMLocationElem = "TIMLocationElem"
	//TIMFaceElem	表情消息。
	TIMFaceElem = "TIMFaceElem"
	//TIMCustomElem	自定义消息，当接收方为 iOS 系统且应用处在后台时，此消息类型可携带除文本以外的字段到 APNs。一条组合消息中只能包含一个 TIMCustomElem 自定义消息元素
	TIMCustomElem = "TIMCustomElem"
)

// 文本消息元素
// Text	String	消息内容。当接收方为 iOS 或 Android 后台在线时，作为离线推送的文本展示。
func NewTextElem(text string) MsgBody {
	return MsgBody{
		MsgContent: MsgContent{
			Text: text,
		},
		MsgType: TIMTextElem,
	}
}

// 地理位置消息元素
// Desc	String	地理位置描述信息。
// Latitude	Number	纬度。
// Longitude	Number	经度。
func NewLocationElem(desc string, latitude, longtitude float64) MsgBody {
	return MsgBody{
		MsgContent: MsgContent{
			Desc:      desc,
			Latitude:  latitude,
			Longitude: longtitude,
		},
		MsgType: TIMLocationElem,
	}
}

// 自定义消息元素
// Index	Number	表情索引，用户自定义。
// Data	String	额外数据。
func NewFaceElem(index int, data string) MsgBody {
	return MsgBody{
		MsgContent: MsgContent{
			Data:  data,
			Index: index,
		},
		MsgType: TIMFaceElem,
	}
}

// 自定义消息元素
// Data	String	自定义消息数据。 不作为 APNs 的 payload 字段下发，故从 payload 中无法获取 Data 字段。
// Desc	String	自定义消息描述信息。当接收方为 iOS 或 Android 后台在线时，做离线推送文本展示。
// 若发送自定义消息的同时设置了 OfflinePushInfo.Desc 字段，此字段会被覆盖，请优先填 OfflinePushInfo.Desc 字段。
// 当消息中只有一个 TIMCustomElem 自定义消息元素时，如果 Desc 字段和 OfflinePushInfo.Desc 字段都不填写，将收不到该条消息的离线推送，需要填写 OfflinePushInfo.Desc 字段才能收到该消息的离线推送。
// Ext	String	扩展字段。当接收方为 iOS 系统且应用处在后台时，此字段作为 APNs 请求包 Payloads 中的 Ext 键值下发，Ext 的协议格式由业务方确定，APNs 只做透传。
// Sound	String	自定义 APNs 推送铃音。
func NewCustomElem(data, desc, ext, sound string) MsgBody {
	return MsgBody{
		MsgContent: MsgContent{
			Data:  data,
			Desc:  desc,
			Ext:   ext,
			Sound: sound,
		},
		MsgType: TIMCustomElem,
	}
}

//From_Account	String	选填	消息发送方 UserID（用于指定发送消息方帐号）
//To_Account	String	必填	消息接收方 UserID
type SingleChatMsg struct {
	FromAccount string `json:"From_Account"`
	ToAccount   string `json:"To_Account"`
	ChatMsg
}

type BatchChatMsg struct {
	FromAccount string   `json:"From_Account"`
	ToAccount   []string `json:"To_Account"`
	ChatMsg
}

//SyncOtherMachine	Integer
//选填	1：把消息同步到 From_Account 在线终端和漫游上；
//      2：消息不同步至 From_Account；
//若不填写默认情况下会将消息存 From_Account 漫游
//MsgLifeTime	Integer	选填	消息离线保存时长（单位：秒），最长为7天（604800秒）
//若设置该字段为0，则消息只发在线用户，不保存离线
//若设置该字段超过7天（604800秒），仍只保存7天
//若不设置该字段，则默认保存7天
//MsgRandom	Integer	必填	消息随机数，由随机函数产生，用于后台定位问题
//MsgTimeStamp	Integer	选填	消息时间戳，UNIX 时间戳（单位：秒）
//MsgType	String	必填	TIM 消息对象类型，目前支持的消息对象包括：TIMTextElem(文本消息)，TIMFaceElem(表情消息)，TIMLocationElem(位置消息)，TIMCustomElem(自定义消息)
//MsgContent	Object	必填	对于每种 MsgType 用不同的 MsgContent 格式，具体可参考 消息格式描述
//OfflinePushInfo	Object	选填	离线推送信息配置，具体可参考 消息格式描述
type ChatMsg struct {
	SyncOtherMachine  int             `json:"SyncOtherMachine"`
	SyncFromOldSystem int             `json:"SyncFromOldSystem" `
	MsgLifeTime       int             `json:"MsgLifeTime"`
	MsgRandom         int             `json:"MsgRandom"`
	MsgTimeStamp      int64           `json:"MsgTimeStamp"`
	MsgBody           []MsgBody       `json:"MsgBody"`
	OfflinePushInfo   OfflinePushInfo `json:"OfflinePushInfo"`
}

//https://cloud.tencent.com/document/product/269/2720#.E7.A6.BB.E7.BA.BF.E6.8E.A8.E9.80.81-offlinepushinfo-.E8.AF.B4.E6.98.8E
//PushFlag	Integer	选填	0表示推送，1表示不离线推送。
//Title	String	选填	离线推送标题。该字段为 iOS 和 Android 共用。
//Desc	String	选填	离线推送内容。该字段会覆盖上面各种消息元素 TIMMsgElement 的离线推送展示文本。
//若发送的消息只有一个 TIMCustomElem 自定义消息元素，该 Desc 字段会覆盖 TIMCustomElem 中的 Desc 字段。如果两个 Desc 字段都不填，将收不到该自定义消息的离线推送。
//Ext	String	选填	离线推送透传内容。
//AndroidInfo.Sound	String	选填	Android 离线推送声音文件路径。
//AndroidInfo.OPPOChannelID	String	选填	OPPO 手机 Android 8.0 以上的 NotificationChannel 通知适配字段。
//ApnsInfo.BadgeMode	Integer	选填	这个字段缺省或者为0表示需要计数，为1表示本条消息不需要计数，即右上角图标数字不增加。
//ApnsInfo.Title	String	选填	该字段用于标识 APNs 推送的标题，若填写则会覆盖最上层 Title。
//ApnsInfo.SubTitle	String	选填	该字段用于标识 APNs 推送的子标题。
//ApnsInfo.Image	String	选填	该字段用于标识 APNs 携带的图片地址，当客户端拿到该字段时，可以通过下载图片资源的方式将图片展示在弹窗上。
type OfflinePushInfo struct {
	PushFlag    int         `json:"PushFlag"`
	Desc        string      `json:"Desc"`
	Ext         string      `json:"Ext"`
	AndroidInfo AndroidInfo `json:"AndroidInfo"`
	ApnsInfo    ApnsInfo    `json:"ApnsInfo"`
}

type ErrorList struct {
	ToAccount string `json:"To_Account"`
	ErrorCode int    `json:"ErrorCode"`
}

type AndroidInfo struct {
	Sound string `json:"Sound"`
}

type ApnsInfo struct {
	Sound     string `json:"Sound"`
	BadgeMode int    `json:"BadgeMode"`
	Title     string `json:"Title"`
	SubTitle  string `json:"SubTitle"`
	Image     string `json:"Image"`
}

type RoamMsgReq struct {
	FromAccount string `json:"From_Account"`
	ToAccount   string `json:"To_Account"`
	MaxCnt      int    `json:"MaxCnt"`
	MinTime     int    `json:"MinTime"`
	MaxTime     int    `json:"MaxTime"`
	LastMsgKey  string `json:"LastMsgKey"`
}

type RoamMsg struct {
	FromAccount  string `json:"From_Account"`
	ToAccount    string `json:"To_Account"`
	MsgSeq       int    `json:"MsgSeq"`
	MsgRandom    int    `json:"MsgRandom"`
	MsgTimeStamp int    `json:"MsgTimeStamp"`
	MsgFlagBits  int    `json:"MsgFlagBits"`
	MsgKey       string `json:"MsgKey"`
	MsgBody      []struct {
		MsgType    string `json:"MsgType"`
		MsgContent struct {
			Text string `json:"Text"`
		} `json:"MsgContent"`
	} `json:"MsgBody"`
}

type RoamMsgResp struct {
	ActionStatus string    `json:"ActionStatus"`
	ErrorInfo    string    `json:"ErrorInfo"`
	ErrorCode    int       `json:"ErrorCode"`
	Complete     int       `json:"Complete"`
	MsgCnt       int       `json:"MsgCnt"`
	LastMsgTime  int       `json:"LastMsgTime"`
	LastMsgKey   string    `json:"LastMsgKey"`
	MsgList      []RoamMsg `json:"MsgList"`
}

type MsgWithdraw struct {
	FromAccount string `json:"From_Account"`
	ToAccount   string `json:"To_Account"`
	MsgKey      string `json:"MsgKey"`
}

type MsgBody struct {
	MsgContent MsgContent `json:"MsgContent"`
	MsgType    string     `json:"MsgType"`
}

type MsgContent struct {
	Data      string  `json:"Data"`
	Desc      string  `json:"Desc"`
	Ext       string  `json:"Ext"`
	Text      string  `json:"Text"`
	Sound     string  `json:"Sound"`
	Index     int     `json:"Index"`
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
}
