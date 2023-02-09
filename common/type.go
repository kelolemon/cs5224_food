package common

type CreateNewUserReq struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	RealName string `json:"real_name,omitempty"`
	Email    string `json:"email,omitempty"`
	Mobile   string `json:"mobile,omitempty"`
	Age      int    `json:"age,omitempty"`
	Gender   int    `json:"gender,omitempty"`
}

// CreateNewUserResp Type
// Code = 0, Msg = "successful"
// Code = 1, Msg = error msg, such as "username used"
type CreateNewUserResp struct {
	Msg  string `json:"msg"`
	Code int32  `json:"code"`
}

type CreateNewLoginSessionReq struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// CreateNewLoginSessionResp
// Code = 1 means something wrong, error can be seen at message, such as Msg = "wrong password"
// Code = 0 means correct, session is at Session
type CreateNewLoginSessionResp struct {
	Msg     string `json:"msg,omitempty"`
	Session string `json:"session,omitempty"`
	Code    int32  `json:"code,omitempty"`
}

type GetUserInfoReq struct {
	Session string `form:"session"`
}

type GetUserInfoResp struct {
	Code int32  `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
	User User   `json:"user"`
}
