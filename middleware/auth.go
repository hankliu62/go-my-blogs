package middleware

import (
	"net/http"

	"strings"

	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/mgo.v2/bson"
	"hankliu.com.cn/go-my-blog/models"
	"hankliu.com.cn/go-my-blog/share/utils"
)

// AuthLogin check user token is access
func AuthLogin(handle gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")

		if len(strings.TrimSpace(token)) >= 0 {
			CToken := &models.Token{}
			condition := bson.M{
				"accesstoken": token,
				"isDelete":    false,
			}
			err := CToken.GetByCondition(condition)

			if err == nil {
				isAccess := utils.IsToday(CToken.ExpireTime)

				if isAccess {
					c.Set("UserID", CToken.UserID)
					handle(c)
					return
				}
			}
		}

		c.JSON(http.StatusUnauthorized, &gin.H{
			"message": "You have not logined",
			"status":  http.StatusUnauthorized,
			"code":    0,
		})
	}
}
