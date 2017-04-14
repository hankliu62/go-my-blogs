package services

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/mgo.v2/bson"
	"hankliu.com.cn/go-my-blog/models"
)

// DeletePost delete post api
func (BlogService) DeletePost(c *gin.Context) {
	id := c.Param("id")
	UserID, _ := c.Get("UserID")
	hexID := bson.ObjectIdHex(id)
	hexUserID := UserID.(bson.ObjectId)

	CPost := &models.Post{}
	selector := bson.M{"_id": hexID, "userId": hexUserID}
	if err := CPost.Delete(selector); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "delete post success",
			"status":  http.StatusOK,
			"code":    0,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"name":    "Server error",
			"status":  http.StatusInternalServerError,
			"code":    0,
		})
	}
}
