/*
 *
 * app_info_types.go
 * TencentIM
 *
 * Created by lintao on 2020/4/22 2:37 下午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package TencentIM

type AppInfo struct {
	ErrorCode int    `json:"ErrorCode"`
	ErrorInfo string `json:"ErrorInfo"`
	Result    []Info `json:"Result"`
}

type Info struct {
	APNSMsgNum           string `json:"APNSMsgNum"`
	ActiveUserNum        string `json:"ActiveUserNum"`
	AppID                string `json:"AppId"`
	AppName              string `json:"AppName"`
	C2CAPNSMsgNum        string `json:"C2CAPNSMsgNum"`
	C2CDownMsgNum        string `json:"C2CDownMsgNum"`
	C2CSendMsgUserNum    string `json:"C2CSendMsgUserNum"`
	C2CUpMsgNum          string `json:"C2CUpMsgNum"`
	CallBackReq          string `json:"CallBackReq"`
	CallBackRsp          string `json:"CallBackRsp"`
	ChainDecrease        string `json:"ChainDecrease"`
	ChainIncrease        string `json:"ChainIncrease"`
	Company              string `json:"Company"`
	Date                 string `json:"Date"`
	DownMsgNum           string `json:"DownMsgNum"`
	GroupAPNSMsgNum      string `json:"GroupAPNSMsgNum"`
	GroupAllGroupNum     string `json:"GroupAllGroupNum"`
	GroupDestroyGroupNum string `json:"GroupDestroyGroupNum"`
	GroupDownMsgNum      string `json:"GroupDownMsgNum"`
	GroupJoinGroupTimes  string `json:"GroupJoinGroupTimes"`
	GroupNewGroupNum     string `json:"GroupNewGroupNum"`
	GroupQuitGroupTimes  string `json:"GroupQuitGroupTimes"`
	GroupSendMsgGroupNum string `json:"GroupSendMsgGroupNum"`
	GroupSendMsgUserNum  string `json:"GroupSendMsgUserNum"`
	GroupUpMsgNum        string `json:"GroupUpMsgNum"`
	LoginTimes           string `json:"LoginTimes"`
	LoginUserNum         string `json:"LoginUserNum"`
	MaxOnlineNum         string `json:"MaxOnlineNum"`
	RegistUserNumOneDay  string `json:"RegistUserNumOneDay"`
	RegistUserNumTotal   string `json:"RegistUserNumTotal"`
	SendMsgUserNum       string `json:"SendMsgUserNum"`
	UpMsgNum             string `json:"UpMsgNum"`
}

type AppInfoField string

const (
	AppName              AppInfoField = "AppName"              //应用名称
	AppId                AppInfoField = "AppId"                // 应用 SDKAppID
	Company              AppInfoField = "Company"              //所属客户名称
	ActiveUserNum        AppInfoField = "ActiveUserNum"        //活跃用户数
	RegistUserNumOneDay  AppInfoField = "RegistUserNumOneDay"  //新增注册人数
	RegistUserNumTotal   AppInfoField = "RegistUserNumTotal"   //累计注册人数
	LoginTimes           AppInfoField = "LoginTimes"           //登录次数
	LoginUserNum         AppInfoField = "LoginUserNum"         //登录人数
	UpMsgNum             AppInfoField = "UpMsgNum"             //上行消息数
	SendMsgUserNum       AppInfoField = "SendMsgUserNum"       //发消息人数
	APNSMsgNum           AppInfoField = "APNSMsgNum"           //APNs 推送数
	C2CUpMsgNum          AppInfoField = "C2CUpMsgNum"          //上行消息数（C2C)
	C2CSendMsgUserNum    AppInfoField = "C2CSendMsgUserNum"    //发消息人数（C2C）
	C2CAPNSMsgNum        AppInfoField = "C2CAPNSMsgNum"        //APNs 推送数（C2C）
	MaxOnlineNum         AppInfoField = "MaxOnlineNum"         //最高在线人数
	ChainIncrease        AppInfoField = "ChainIncrease"        //关系链对数增加量
	ChainDecrease        AppInfoField = "ChainDecrease"        //关系链对数删除量
	GroupUpMsgNum        AppInfoField = "GroupUpMsgNum"        //上行消息数（群）
	GroupSendMsgUserNum  AppInfoField = "GroupSendMsgUserNum"  //发消息人数（群）
	GroupAPNSMsgNum      AppInfoField = "GroupAPNSMsgNum"      //APNs 推送数（群）
	GroupSendMsgGroupNum AppInfoField = "GroupSendMsgGroupNum" //发消息群组数
	GroupJoinGroupTimes  AppInfoField = "GroupJoinGroupTimes"  //入群总数
	GroupQuitGroupTimes  AppInfoField = "GroupQuitGroupTimes"  //退群总数
	GroupNewGroupNum     AppInfoField = "GroupNewGroupNum"     //新增群组数
	GroupAllGroupNum     AppInfoField = "GroupAllGroupNum"     //累计群组数
	GroupDestroyGroupNum AppInfoField = "GroupDestroyGroupNum" //解散群个数
	CallBackReq          AppInfoField = "CallBackReq"          //回调请求数
	CallBackRsp          AppInfoField = "CallBackRsp"          //回调应答数
	Date                 AppInfoField = "Date"                 //日期
)

type AppHistory struct {
	File         []Msg  `json:"File"`
	ActionStatus string `json:"ActionStatus"`
	ErrorInfo    string `json:"ErrorInfo"`
	ErrorCode    int    `json:"ErrorCode"`
}

//URL	String	消息记录文件下载地址
//ExpireTime	String	下载地址过期时间，请在过期前进行下载，若地址失效，请通过该接口重新获取
//FileSize	Integer	GZip 压缩前的文件大小（单位 Byte）
//FileMD5	String	GZip 压缩前的文件 MD5
//GzipSize	Integer	GZip 压缩后的文件大小（单位 Byte）
//GzipMD5	String	GZip 压缩后的文件 MD5
type Msg struct {
	URL        string `json:"URL"`
	ExpireTime string `json:"ExpireTime"`
	FileSize   int    `json:"FileSize"`
	FileMD5    string `json:"FileMD5"`
	GzipSize   int    `json:"GzipSize"`
	GzipMD5    string `json:"GzipMD5"`
}
