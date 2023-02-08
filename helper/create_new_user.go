package helper

import (
	"food/common"
	"food/dao"
)

func CreateNewUser(r common.CreateNewUserReq) (res common.CreateNewUserResp, err error) {
	code, err := dao.CreateUser(common.User{
		ID:       r.Username,
		Name:     r.RealName,
		Password: r.Password,
		Email:    r.Email,
		Mobile:   r.Mobile,
		Age:      r.Age,
		Gender:   r.Gender,
	})
	if err != nil {
		return common.CreateNewUserResp{
			Msg:  err.Error(),
			Code: code,
		}, nil
	}
	return common.CreateNewUserResp{}, nil
}
