package helper

import (
	"food/client"
	"food/common"
	"food/dao"
	"github.com/go-redis/redis"
)

func GetUserInfo(r common.GetUserInfoReq) (res common.GetUserInfoResp, err error) {
	id, err := client.RedisCli.Get(r.Session).Result()
	if err == redis.Nil {
		return common.GetUserInfoResp{
			Code: 1,
			Msg:  "session is wrong or expired",
			User: common.User{},
		}, nil
	}

	if err != nil {
		return common.GetUserInfoResp{}, err
	}

	user, err := dao.GetUser(id)
	if err != nil {
		return common.GetUserInfoResp{}, err
	}

	return common.GetUserInfoResp{
		Code: 0,
		Msg:  "OK",
		User: user,
	}, nil
}
