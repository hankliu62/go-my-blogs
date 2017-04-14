package services

import (
	"net/http"
	"time"

	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/mgo.v2/bson"
	"hankliu.com.cn/go-my-blog/models"
)

// UpdatePost update post
func (BlogService) UpdatePost(c *gin.Context) {
	id := c.Param("id")

	CPost := &models.Post{}
	c.BindJSON(CPost)

	UserID, _ := c.Get("UserID")
	hexUserID := UserID.(bson.ObjectId)
	hexID := bson.ObjectIdHex(id)

	selector := bson.M{"_id": hexID, "userId": hexUserID}

	updatedAt := time.Now()
	updatorFields := bson.M{
		"title":     CPost.Title,
		"content":   CPost.Content,
		"tags":      CPost.Tags,
		"Comments":  CPost.Comments,
		"updatedAt": updatedAt,
	}
	updator := bson.M{"$set": updatorFields}

	if err := CPost.Update(selector, updator); err == nil {
		CPost.ID = hexID
		CPost.UpdatedAt = updatedAt
		CPost.UserID = hexUserID
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
