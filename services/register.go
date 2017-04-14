package services

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
	"hankliu.com.cn/go-my-blog/models"
	"hankliu.com.cn/go-my-blog/share/utils"
)

// Register register user and return token
func (BlogService) Register(c *gin.Context) {
	CUser = &models.User{}
	c.BindJSON(CUser)

	err := CUser.Create()

	errorResp := &gin.H{
		"message": "Server error",
		"name":    "Server error",
		"status":  http.StatusInternalServerError,
		"code":    0,
	}
	if err == nil {
		CToken := &models.Token{
			AccessToken: utils.GenerateAccessToken(),
			UserID:      CUser.ID,
		}
		err = CToken.Create()

		if err == nil {
			VTOKEN := &models.VToken{}
			VTOKEN.Copy(CToken, CUser)
			c.JSON(http.StatusOK, VTOKEN)
		} else {
			c.JSON(http.StatusInternalServerError, errorResp)
		}
	} else {
		c.JSON(http.StatusInternalServerError, errorResp)
	}
}
