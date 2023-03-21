package helper

import (
	"food/common"
	"food/dao"
)

func GetDialogs(r common.GetDialogReq) (res common.GetDialogResp, err error) {
	dialogs, err := dao.FindDialog(r.DinerID)
	if err != nil {
		return common.GetDialogResp{}, err
	}
	return common.GetDialogResp{
		Code:    0,
		Msg:     "OK",
		Dialogs: dialogs,
	}, nil
}
