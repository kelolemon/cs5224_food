package helper

import (
	"food/client"
	"food/common"
	"food/dao"
	"github.com/go-redis/redis"
)

func GetUserInfo(r common.GetUserInfoReq) (res common.GetUserInfoResp, err error) {
	id, err := client.RedisCli.Get(r.Session).Result()
	code := int32(0)
	if err == redis.Nil {
		code = 1
		return common.GetUserInfoResp{
			Code: code,
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
		Code: code,
		Msg:  "OK",
		User: user,
	}, nil
}
