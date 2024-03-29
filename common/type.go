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
	Msg     string `json:"msg"`
	Session string `json:"session"`
	Code    int32  `json:"code"`
}

type GetUserInfoReq struct {
	Session string `form:"session"`
}

type GetUserInfoResp struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	User User   `json:"user"`
}

type GetAllDinersResp struct {
	Code   int32   `json:"code"`
	Msg    string  `json:"msg"`
	Diners []Diner `json:"diners"`
}

type CreateDinerReq struct {
	Name        string `json:"name"`
	Category    string `json:"category"`
	Address     string `json:"address"`
	URL         string `json:"url"`
	Ratings     string `json:"ratings"`
	Reviews     string `json:"reviews"`
	Reviewers   string `json:"reviewers"`
	ReviewTimes string `json:"review_times"`
	ImageURL    string `json:"image_url"`
}

type CreateDinerResp struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

type CreateCommentReq struct {
	Username string `json:"username"`
	DinerID  int32  `json:"diner_id"`
	Msg      string `json:"comment"`
}

type CreateCommentResp struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

type GetCommentReq struct {
	DinerID int32 `form:"diner_id"`
}

type GetCommentResp struct {
	Code     int32     `json:"code"`
	Msg      string    `json:"msg"`
	Comments []Comment `json:"comments"`
}

type CreateDialogReq struct {
	Username string `json:"username"`
	FatherID int32  `json:"father_id"`
	DinerID  int32  `json:"diner_id"`
	Msg      string `json:"comment"`
}

type CreateDialogResp struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

type GetDialogReq struct {
	DinerID int32 `form:"diner_id"`
}

type GetDialogResp struct {
	Code    int32    `json:"code"`
	Msg     string   `json:"msg"`
	Dialogs []Dialog `json:"dialogs"`
}
