package services

import (
	"net/http"

	"encoding/json"

	"fmt"

	"strings"

	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/mgo.v2/bson"
	"hankliu.com.cn/go-my-blog/constants"
	"hankliu.com.cn/go-my-blog/models"
	"hankliu.com.cn/go-my-blog/share/utils"
)

// SearchPosts search posts list, searching by title, filtering by tags and sorting by title or creation time, with pagination
func (BlogService) SearchPosts(c *gin.Context) {
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

// GetSearchPostCondition get condition
func GetSearchPostCondition(c *gin.Context, condition *bson.M, orderbys *[]string) {
	if searchWord, ok := c.GetQuery("searchWord"); ok {
		fuzzyValue := bson.RegEx{
			Pattern: utils.FormatRegexpStr(searchWord),
			Options: "i",
		}

		(*condition)["title"] = fuzzyValue
	}

	if queryOrderBy, ok := c.GetQuery("orderBy"); ok {
		orderBy := map[string]interface{}{}
		if err := json.Unmarshal([]byte(queryOrderBy), &orderBy); err == nil {
			for sortKey, sortDirection := range orderBy {
				sortDirectionStr := strings.TrimSpace(sortDirection.(string))
				if strings.Compare(sortDirectionStr, constants.DESC) == 0 {
					sortKey = fmt.Sprintf("-%s", sortKey)
				}
				*orderbys = append(*orderbys, sortKey)
			}
		}
	}

	if tags, ok := c.GetQueryArray("tags"); ok && len(tags) > 0 {
		(*condition)["tags"] = bson.M{"$in": tags}
	}
}
