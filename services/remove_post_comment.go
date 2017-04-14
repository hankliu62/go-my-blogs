package services

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/mgo.v2/bson"
	"hankliu.com.cn/go-my-blog/models"
)

// RemoveComment remove comment for post
func (BlogService) RemoveComment(c *gin.Context) {
	id := c.Param("id")
	hexID := bson.ObjectIdHex(id)

	commentID := c.Param("commentId")
	hexCommentID := bson.ObjectIdHex(commentID)

	UserID, _ := c.Get("UserID")
	hexUserID := UserID.(bson.ObjectId)

	CPost := &models.Post{
		ID: hexID,
	}

	CComment := &models.Comment{
		ID:     hexCommentID,
		UserID: hexUserID,
	}

	if err := CPost.RemoveComment(CComment); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "remove comment success",
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
