package services

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/mgo.v2/bson"
	"hankliu.com.cn/go-my-blog/models"
)

// AppendComment append comment for post
func (BlogService) AppendComment(c *gin.Context) {
	id := c.Param("id")
	hexID := bson.ObjectIdHex(id)

	CPost := &models.Post{
		ID: hexID,
	}

	VComment := &models.VComment{}
	c.BindJSON(VComment)

	CComment := &models.Comment{
		Content: VComment.Content,
	}

	if len(VComment.UserID) > 0 {
		CComment.UserID = bson.ObjectIdHex(VComment.UserID)
	}

	if err := CPost.AppendComment(CComment); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "create comment success",
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
