package statecode

const (
	//language
	LANG_ZH    = 111
	LANG_EN    = 112
	LANG_ZH_TW = 113

	//common
	COMMON_SUCCESS        = 0
	COMMON_ERR_SERVER_ERR = 1000
	PARAMETER_EMPTY_ERR   = 1001

	//token
	TOKEN_EMPTY = 1101 //token empty
	TOKEN_ERR   = 1102 //token error

	// muti-sign
	P_NAME_EMPTY  = 1201 //p_name empty
	CHAINID_EMPTY = 1202 //chain id empty
	CHAINID_ERR   = 1203 //chain id error

	// user login
	NAME_EMPTY           = 1301 //username empty
	PASSWORD_EMPTY       = 1302 //password empty
	NAME_OR_PASSWORD_ERR = 1303 //name or password error

)

var Msg = map[int]map[int]string{

	0: map[int]string{
		LANG_ZH:    "成功",
		LANG_ZH_TW: "成功",
		LANG_EN:    "success",
	},
	1000: map[int]string{
		LANG_ZH:    "服务器繁忙，请稍后重试",
		LANG_ZH_TW: "服務器繁忙，請稍後重試",
		LANG_EN:    "server is busy, please try again later",
	},
	1001: map[int]string{
		LANG_ZH:    "参数不能为空",
		LANG_ZH_TW: "参数不能為空",
		LANG_EN:    "parameter is empty",
	},
	1101: map[int]string{
		LANG_ZH:    "token 不能为空",
		LANG_ZH_TW: "token 不能為空",
		LANG_EN:    "token required",
	},
	1102: map[int]string{
		LANG_ZH:    "token错误",
		LANG_ZH_TW: "token錯誤",
		LANG_EN:    "token invalid",
	},
	1201: map[int]string{
		LANG_ZH:    "sp_name 不能为空",
		LANG_ZH_TW: "sp_name 不能為空",
		LANG_EN:    "sp_name required",
	},
	1202: map[int]string{
		LANG_ZH:    "chain_id 不能为空",
		LANG_ZH_TW: "chain_id 不能為空",
		LANG_EN:    "chain_id required",
	},
	1203: map[int]string{
		LANG_ZH:    "chain_id 错误",
		LANG_ZH_TW: "chain_id 錯誤",
		LANG_EN:    "chain_id error",
	},
	1301: map[int]string{
		LANG_ZH:    "name 不能为空",
		LANG_ZH_TW: "name 不能為空",
		LANG_EN:    "name required",
	},
	1302: map[int]string{
		LANG_ZH:    "password 不能为空",
		LANG_ZH_TW: "password 不能為空",
		LANG_EN:    "password required",
	},
	1303: map[int]string{
		LANG_ZH:    "用户名或密码错误",
		LANG_ZH_TW: "用戶名或密碼錯誤",
		LANG_EN:    "name or password error",
	},
}

func GetMsg(c int, lang int) string {
	_, ok := Msg[c]
	if ok {
		msg, ok := Msg[c][lang]
		if ok {
			return msg
		}
		return Msg[COMMON_ERR_SERVER_ERR][lang]
	}
	return Msg[COMMON_ERR_SERVER_ERR][lang]
}
