package dao

import (
	"food/client"
	"food/common"
)

func GetAllDinners() (diners []common.Diner, err error) {
	err = client.MysqlDB.Find(&diners).Error

	if err != nil {
		return nil, err
	}

	return diners, nil
}

func CreateDiner(diner common.Diner) (err error) {
	err = client.MysqlDB.Create(&diner).Error
	return err
}
