package helper

import (
	"food/common"
	"food/dao"
)

func GetAllDiners() (res common.GetAllDinersResp, err error) {
	diners, err := dao.GetAllDinners()
	code := int32(0)

	if err != nil {
		code = 1
		return common.GetAllDinersResp{
			Code:   code,
			Msg:    err.Error(),
			Diners: nil,
		}, nil
	}

	return common.GetAllDinersResp{
		Code:   code,
		Msg:    "OK",
		Diners: diners,
	}, nil
}
