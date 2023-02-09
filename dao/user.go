package dao

import (
	"errors"
	"fmt"
	"food/client"
	"food/common"

	"gorm.io/gorm"
)

func CreateUser(r common.User) (code int32, err error) {
	var user common.User
	res := client.MysqlDB.First(&user, "id = ?", r.ID)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		res = client.MysqlDB.Create(&r)
		if res.Error != nil {
			return 1, res.Error
		}
		return 0, nil
	}

	if res.Error != nil {
		return 1, res.Error
	}

	return 1, fmt.Errorf("username is existed")
}

func GetUser(id string) (user common.User, err error) {
	res := client.MysqlDB.First(&user, "id = ?", id)
	if res.Error != nil {
		return common.User{}, res.Error
	}
	return user, nil
}
