package dao

import (
	"food/client"
	"food/common"
)

func CreateDialog(dialog common.Dialog) (err error) {
	err = client.MysqlDB.Create(&dialog).Error
	return err
}

func FindDialog(dinerID int32) (dialogs []common.Dialog, err error) {
	err = client.MysqlDB.Where("diner_id = ?", dinerID).Find(&dialogs).Error
	return dialogs, err
}
