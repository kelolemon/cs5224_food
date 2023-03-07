package helper

import (
	"food/common"
	"food/dao"
)

func CreateComment(r common.CreateCommentReq) (res common.CreateCommentResp, err error) {
	comment := common.Comment{
		UserID:  r.Username,
		DinerID: r.DinerID,
		Msg:     r.Msg,
	}
	err = dao.CreateComment(comment)
	if err != nil {
		return common.CreateCommentResp{}, err
	}
	return common.CreateCommentResp{
		Code: 0,
		Msg:  "OK",
	}, nil
}
