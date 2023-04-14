package method

import (
	"encoding/json"
	"food/common"
	"food/dao"
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

// CreateLoginSession POST method
func CreateLoginSession(c *gin.Context) {
	raw, _ := c.GetRawData()
	log.Printf("[info] create new login session, request body = %v", string(raw))
	var createNewLoginSessionReq common.CreateNewLoginSessionReq
	err := json.Unmarshal(raw, &createNewLoginSessionReq)
	if err != nil {
		log.Printf("[warn] request json converted error, err = %v, request body = %v", err, string(raw))
		c.JSON(400, gin.H{
			"message": "bad request",
			"err":     err,
		})
		return
	}
	res, err := helper.CreateNewLoginSession(createNewLoginSessionReq)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal error",
			"err":     err,
		})
		return
	}
	c.JSON(200, res)
}

func GetUser(c *gin.Context) {
	var getUserReq common.GetUserInfoReq
	err := c.ShouldBindQuery(&getUserReq)
	if err != nil {
		log.Printf("get param error, err=%v", err)
		c.JSON(400, gin.H{
			"message": "bad request",
			"err":     err,
		})
		return
	}
	log.Printf("[info] get user info for session: %v", getUserReq.Session)
	res, err := helper.GetUserInfo(getUserReq)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal error",
			"err":     err,
		})
		return
	}
	c.JSON(200, res)
}

func GetAllDiners(c *gin.Context) {
	res, err := helper.GetAllDiners()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal error",
			"err":     err,
		})
		return
	}
	c.JSON(200, res)
}

func CreateDiner(c *gin.Context) {
	raw, _ := c.GetRawData()
	log.Printf("[info] create dinner, request body = %v", string(raw))
	var createDinerReq common.CreateDinerReq
	err := json.Unmarshal(raw, &createDinerReq)
	if err != nil {
		log.Printf("[warn] request json converted error, err = %v, request body = %v", err, string(raw))
		c.JSON(400, gin.H{
			"message": "bad request",
			"err":     err,
		})
		return
	}
	res, err := helper.CreateDiner(createDinerReq)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal error",
			"err":     err,
		})
		return
	}
	c.JSON(200, res)
}

func CreateComment(c *gin.Context) {
	raw, _ := c.GetRawData()
	log.Printf("[info] create comment, request body = %v", string(raw))
	var createCommentReq common.CreateCommentReq
	err := json.Unmarshal(raw, &createCommentReq)
	if err != nil {
		log.Printf("[warn] request json converted error, err = %v, request body = %v", err, string(raw))
		c.JSON(400, gin.H{
			"message": "bad request",
			"err":     err,
		})
		return
	}
	res, err := helper.CreateComment(createCommentReq)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal error",
			"err":     err,
		})
		return
	}
	c.JSON(200, res)
}

func GetComments(c *gin.Context) {
	var getCommentReq common.GetCommentReq
	err := c.ShouldBindQuery(&getCommentReq)
	if err != nil {
		log.Printf("get param error, err=%v", err)
		c.JSON(400, gin.H{
			"message": "bad request",
			"err":     err,
		})
		return
	}
	res, err := helper.GetComments(getCommentReq)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal error",
			"err":     err,
		})
		return
	}
	c.JSON(200, res)
}

func CreateDialog(c *gin.Context) {
	raw, _ := c.GetRawData()
	log.Printf("[info] create dialog, request body = %v", string(raw))
	var createDialogReq common.CreateDialogReq
	err := json.Unmarshal(raw, &createDialogReq)
	if err != nil {
		log.Printf("[warn] request json converted error, err = %v, request body = %v", err, string(raw))
		c.JSON(400, gin.H{
			"message": "bad request",
			"err":     err,
		})
		return
	}
	res, err := helper.CreateDialog(createDialogReq)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal error",
			"err":     err,
		})
		return
	}
	c.JSON(200, res)
}

func GetDialogs(c *gin.Context) {
	var getDialogReq common.GetDialogReq
	err := c.ShouldBindQuery(&getDialogReq)
	if err != nil {
		log.Printf("get param error, err=%v", err)
		c.JSON(400, gin.H{
			"message": "bad request",
			"err":     err,
		})
		return
	}
	res, err := helper.GetDialogs(getDialogReq)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal error",
			"err":     err,
		})
		return
	}
	c.JSON(200, res)
}

func GetAds(c *gin.Context) {
	res, err := dao.GetAds()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal error",
			"err":     err,
		})
		return
	}
	c.JSON(200, res)
}
