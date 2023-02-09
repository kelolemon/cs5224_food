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
	if err != nil {
		return common.CreateNewLoginSessionResp{
			Code:    1,
			Msg:     err.Error(),
			Session: "",
		}, nil
	}

	if user.Password != r.Password {
		return common.CreateNewLoginSessionResp{
			Code:    1,
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
		Code:    0,
		Msg:     "OK",
		Session: session,
	}, nil
}
