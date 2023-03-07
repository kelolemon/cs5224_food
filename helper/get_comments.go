package helper

import (
	"food/common"
	"food/dao"
)

func GetComments(r common.GetCommentReq) (res common.GetCommentResp, err error) {
	comments, err := dao.FindComment(r.DinerID)
	if err != nil {
		return common.GetCommentResp{}, err
	}
	return common.GetCommentResp{
		Code:     0,
		Msg:      "OK",
		Comments: comments,
	}, nil
}
