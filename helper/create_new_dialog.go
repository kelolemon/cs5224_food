package helper

import (
	"food/common"
	"food/dao"
)

func CreateDialog(r common.CreateDialogReq) (res common.CreateDialogResp, err error) {
	dialog := common.Dialog{
		FatherID: r.FatherID,
		UserID:   r.Username,
		DinerID:  r.DinerID,
		Msg:      r.Msg,
	}

	err = dao.CreateDialog(dialog)
	if err != nil {
		return common.CreateDialogResp{}, err
	}
	return common.CreateDialogResp{
		Code: 0,
		Msg:  "OK",
	}, nil
}
