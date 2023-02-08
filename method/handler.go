package method

import (
	"encoding/json"
	"food/common"
	"food/helper"
	"log"

	"github.com/gin-gonic/gin"
)

// CreateNewUser POST method
func CreateNewUser(c *gin.Context) {
	raw, _ := c.GetRawData()
	log.Printf("[info] create new user, request body = %v", string(raw))
	var createNewUserReq common.CreateNewUserReq
	err := json.Unmarshal(raw, &createNewUserReq)
	if err != nil {
		log.Printf("[warn] request json converted error, err = %v, request body = %v", err, string(raw))
		c.JSON(400, gin.H{
			"message": "bad request",
			"err":     err,
		})
		return
	}
	res, err := helper.CreateNewUser(createNewUserReq)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal error",
			"err":     err,
		})
		return
	}
	c.JSON(200, res)
}
