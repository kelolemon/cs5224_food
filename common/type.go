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
