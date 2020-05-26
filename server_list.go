/*
 *
 * server_list.go
 * server
 *
 * Created by lintao on 2020/4/16 2:24 下午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package TencentIM

type IMPath string

const (
	VERSION  = "v4"
	BASE_URL = "https://console.tim.qq.com"
	//帐号管理
	MULTIACCOUNTIMPORT IMPath = "/im_open_login_svc/multiaccount_import" //导入多个帐号
	ACCOUNTIMPORT      IMPath = "/im_open_login_svc/account_import"      //导入单个帐号
	ACCOUNTDELETE      IMPath = "/im_open_login_svc/account_delete"      //删除帐号	v4/im_open_login_svc/account_delete
	ACCOUNTCHECK       IMPath = "/im_open_login_svc/account_check"       //查询帐号	v4/im_open_login_svc/account_check
	KICK               IMPath = "/im_open_login_svc/kick"                //失效帐号登录态	v4/im_open_login_svc/kick

	//单聊消息
	QUERYSTATE        IMPath = "/openim/querystate"        //查询帐号在线状态	v4/openim/querystate
	SENDMSG           IMPath = "/openim/sendmsg"           //单发单聊消息	v4/openim/sendmsg
	BATCHSENDMSG      IMPath = "/openim/batchsendmsg"      //批量发单聊消息	v4/openim/batchsendmsg
	IMPORTMSG         IMPath = "/openim/importmsg"         //导入单聊消息	v4/openim/importmsg
	ADMIN_GETROAMMSG  IMPath = "/openim/admin_getroammsg"  //查询单聊消息	v4/openim/admin_getroammsg
	ADMIN_MSGWITHDRAW IMPath = "/openim/admin_msgwithdraw" //撤回单聊消息
	//资料管理
	PORTRAIT_GET IMPath = "/profile/portrait_get" //拉取资料	v4/profile/portrait_get
	PORTRAIT_SET IMPath = "/profile/portrait_set" //设置资料	v4/profile/portrait_set

	//关系链管理
	FRIEND_ADD        IMPath = "/sns/friend_add"        //添加好友	v4/sns/friend_add
	FRIEND_IMPORT     IMPath = "/sns/friend_import"     //导入好友	v4/sns/friend_import
	FRIEND_DELETE     IMPath = "/sns/friend_delete"     //删除好友	v4/sns/friend_delete
	FRIEND_UPDATE     IMPath = "/sns/friend_update"     //更新好友
	FRIEND_DELETE_ALL IMPath = "/sns/friend_delete_all" //删除所有好友	v4/sns/friend_delete_all
	FRIEND_CHECK      IMPath = "/sns/friend_check"      //校验好友	v4/sns/friend_check
	FRIEND_GET        IMPath = "/sns/friend_get"        //拉取好友	v4/sns/friend_get
	FRIEND_GET_LIST   IMPath = "/sns/friend_get_list"   //拉取指定好友	v4/sns/friend_get_list
	BLACK_LIST_ADD    IMPath = "/sns/black_list_add"    //添加黑名单	v4/sns/black_list_add
	BLACK_LIST_DELETE IMPath = "/sns/black_list_delete" //删除黑名单	v4/sns/black_list_delete
	BLACK_LIST_GET    IMPath = "/sns/black_list_get"    //拉取黑名单	v4/sns/black_list_get
	BLACK_LIST_CHECK  IMPath = "/sns/black_list_check"  //校验黑名单	v4/sns/black_list_check
	GROUP_ADD         IMPath = "/sns/group_add"         //添加分组	v4/sns/group_add
	GROUP_DELETE      IMPath = "/sns/group_delete"      //删除分组	v4/sns/group_delete

	//群组管理

	//获取 App 中的所有群组	v4/group_open_http_svc/get_appid_group_list
	GET_APPID_GROUP_LIST           IMPath = "/group_open_http_svc/get_appid_group_list"
	CREATE_GROUP                   IMPath = "/group_open_http_svc/create_group"                   //创建群组	v4/group_open_http_svc/create_group
	GET_GROUP_INFO                 IMPath = "/group_open_http_svc/get_group_info"                 //获取群组详细资料	v4/group_open_http_svc/get_group_info
	GET_GROUP_MEMBER_INFO          IMPath = "/group_open_http_svc/get_group_member_info"          //获取群成员详细资料	v4/group_open_http_svc/get_group_member_info
	MODIFY_GROUP_BASE_INFO         IMPath = "/group_open_http_svc/modify_group_base_info"         //修改群组基础资料	v4/group_open_http_svc/modify_group_base_info
	ADD_GROUP_MEMBER               IMPath = "/group_open_http_svc/add_group_member"               //增加群组成员	v4/group_open_http_svc/add_group_member
	DELETE_GROUP_MEMBER            IMPath = "/group_open_http_svc/delete_group_member"            //删除群组成员	v4/group_open_http_svc/delete_group_member
	MODIFY_GROUP_MEMBER_INFO       IMPath = "/group_open_http_svc/modify_group_member_info"       //修改群组成员资料	v4/group_open_http_svc/modify_group_member_info
	DESTROY_GROUP                  IMPath = "/group_open_http_svc/destroy_group"                  //解散群组	v4/group_open_http_svc/destroy_group
	GET_JOINED_GROUP_LIST          IMPath = "/group_open_http_svc/get_joined_group_list"          //获取用户所加入的群组	v4/group_open_http_svc/get_joined_group_list
	GET_ROLE_IN_GROUP              IMPath = "/group_open_http_svc/get_role_in_group"              //查询用户在群组中的身份	v4/group_open_http_svc/get_role_in_group
	FORBID_SEND_MSG                IMPath = "/group_open_http_svc/forbid_send_msg"                //批量禁言和取消禁言	v4/group_open_http_svc/forbid_send_msg
	GET_GROUP_SHUTTED_UIN          IMPath = "/group_open_http_svc/get_group_shutted_uin"          //获取群组被禁言用户列表	v4/group_open_http_svc/get_group_shutted_uin
	SEND_GROUP_MSG                 IMPath = "/group_open_http_svc/send_group_msg"                 //在群组中发送普通消息	v4/group_open_http_svc/send_group_msg
	SEND_GROUP_SYSTEM_NOTIFICATION IMPath = "/group_open_http_svc/send_group_system_notification" //在群组中发送系统通知	v4/group_open_http_svc/send_group_system_notification
	GROUP_MSG_RECALL               IMPath = "/group_open_http_svc/group_msg_recall"               //群组消息撤回	v4/group_open_http_svc/group_msg_recall
	CHANGE_GROUP_OWNER             IMPath = "/group_open_http_svc/change_group_owner"             //转让群组	v4/group_open_http_svc/change_group_owner
	IMPORT_GROUP                   IMPath = "/group_open_http_svc/import_group"                   //导入群基础资料	v4/group_open_http_svc/import_group
	IMPORT_GROUP_MSG               IMPath = "/group_open_http_svc/import_group_msg"               //导入群消息	v4/group_open_http_svc/import_group_msg
	IMPORT_GROUP_MEMBER            IMPath = "/group_open_http_svc/import_group_member"            //导入群成员	v4/group_open_http_svc/import_group_member
	SET_UNREAD_MSG_NUM             IMPath = "/group_open_http_svc/set_unread_msg_num"             //设置成员未读消息计数	v4/group_open_http_svc/set_unread_msg_num
	DELETE_GROUP_MSG_BY_SENDER     IMPath = "/group_open_http_svc/delete_group_msg_by_sender"     //删除指定用户发送的消息	v4/group_open_http_svc/delete_group_msg_by_sender
	GROUP_MSG_GET_SIMPLE           IMPath = "/group_open_http_svc/group_msg_get_simple"           //拉取群漫游消息	v4/group_open_http_svc/group_msg_get_simple

	//全局禁言管理
	SETNOSPEAKING IMPath = "/openconfigsvr/setnospeaking" //设置全局禁言	v4/openconfigsvr/setnospeaking
	GETNOSPEAKING IMPath = "/openconfigsvr/getnospeaking" //查询全局禁言	v4/openconfigsvr/getnospeaking

	//运营管理
	GET_HISTORY IMPath = "/open_msg_svc/get_history" //下载消息记录	v4/open_msg_svc/get_history
	GETAPPINFO  IMPath = "/openconfigsvr/getappinfo" //拉取运营数据	v4/openconfigsvr/getappinfo
)
