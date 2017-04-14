package services

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/mgo.v2/bson"
	"hankliu.com.cn/go-my-blog/models"
)

// CreatePost create post Api
func (BlogService) CreatePost(c *gin.Context) {
	CPost := &models.Post{}
	c.BindJSON(CPost)

	UserID, _ := c.Get("UserID")
	CPost.UserID = UserID.(bson.ObjectId)
	if err := CPost.Create(); err == nil {
		VPost := &models.VPost{}
		VPost.Copy(CPost)
		c.JSON(http.StatusOK, VPost)
	} else {
		c.JSON(http.StatusInternalServerError, &gin.H{
			"message": err.Error(),
			"name":    "Server error",
			"status":  http.StatusInternalServerError,
			"code":    0,
		})
	}
}
