package helper

import (
	"food/client"
	"food/common"
	"food/dao"
	"time"

	"github.com/google/uuid"
)

func CreateNewLoginSession(r common.CreateNewLoginSessionReq) (res common.CreateNewLoginSessionResp, err error) {
	user, err := dao.GetUser(r.Username)
	code := int32(0)
	if err != nil {
		code = 1
		return common.CreateNewLoginSessionResp{
			Code:    code,
			Msg:     err.Error(),
			Session: "",
		}, nil
	}

	if user.Password != r.Password {
		code = 1
		return common.CreateNewLoginSessionResp{
			Code:    code,
			Msg:     "password not correct",
			Session: "",
		}, nil
	}

	session := uuid.New().String()
	err = client.RedisCli.Set(session, user.ID, 12*time.Hour).Err()
	if err != nil {
		return common.CreateNewLoginSessionResp{}, err
	}

	return common.CreateNewLoginSessionResp{
		Code:    code,
		Msg:     "OK",
		Session: session,
	}, nil
}
