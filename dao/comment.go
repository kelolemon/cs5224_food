package dao

import (
	"food/client"
	"food/common"
)

func CreateComment(comment common.Comment) (err error) {
	err = client.MysqlDB.Create(&comment).Error
	return err
}

func FindComment(dinerID int32) (comments []common.Comment, err error) {
	err = client.MysqlDB.Where("diner_id = ?", dinerID).Find(&comments).Error
	return comments, err
}
