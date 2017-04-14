package services

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/mgo.v2/bson"
	"hankliu.com.cn/go-my-blog/models"
	"hankliu.com.cn/go-my-blog/share/utils"
)

// SearchOwnPosts search own posts list, searching by title, filtering by tags and sorting by title or creation time, with pagination
func (BlogService) SearchOwnPosts(c *gin.Context) {
	CPost := &models.Post{}

	var (
		condition = bson.M{}
		page      uint32
		pageSize  uint32 = 10
		orderbys         = []string{}
	)

	page, pageSize = utils.ParsePagingCondition(c)
	GetSearchPostCondition(c, &condition, &orderbys)
	orderbys = append(orderbys, "-createdAt")

	UserID, _ := c.Get("UserID")
	hexUserID := UserID.(bson.ObjectId)
	condition["userId"] = hexUserID

	posts, total := CPost.GetAllByPagination(condition, page, pageSize, orderbys)

	vposts := []models.VPost{}
	for _, post := range posts {
		VPost := &models.VPost{}
		VPost.Copy(&post)
		vposts = append(vposts, *VPost)
	}

	resp := gin.H{
		"_links": gin.H{"self": gin.H{"href": c.Request.URL.RequestURI()}},
		"_meta":  utils.GetResultsMeta(int(page), int(pageSize), total),
		"items":  vposts,
	}

	c.JSON(http.StatusOK, resp)
}
