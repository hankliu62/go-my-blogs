package services

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/mgo.v2/bson"
	"hankliu.com.cn/go-my-blog/models"
)

// GetPost view post
func (BlogService) GetPost(c *gin.Context) {
	id := c.Param("id")
	hexID := bson.ObjectIdHex(id)

	CPost := &models.Post{}

	if err := CPost.GetByPk(hexID); err == nil {
		VPost := &models.VPost{}
		VPost.Copy(CPost)
		c.JSON(http.StatusOK, VPost)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"name":    "Server error",
			"status":  http.StatusInternalServerError,
			"code":    0,
		})
	}
}
