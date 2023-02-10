package helper

import (
	"food/common"
	"food/dao"
)

func CreateDiner(r common.CreateDinerReq) (res common.CreateDinerResp, err error) {
	diner := common.Diner{
		Name:        r.Name,
		Category:    r.Category,
		Address:     r.Address,
		URL:         r.URL,
		Ratings:     r.Ratings,
		Reviewers:   r.Reviewers,
		Reviews:     r.Reviews,
		ReviewTimes: r.ReviewTimes,
	}
	err = dao.CreateDiner(diner)
	if err != nil {
		return common.CreateDinerResp{
			Code: 1,
			Msg:  err.Error(),
		}, nil
	}
	return common.CreateDinerResp{
		Code: 0,
		Msg:  "OK",
	}, nil
}
