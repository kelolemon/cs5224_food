package dao

import (
	"food/client"
	"food/common"
)

func GetAds() (ads []common.Ad, err error) {
	err = client.MysqlDB.Find(&ads).Error

	if err != nil {
		return nil, err
	}

	return ads, nil
}
