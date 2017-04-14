package services

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/mgo.v2/bson"

	"time"

	"hankliu.com.cn/go-my-blog/models"
	"hankliu.com.cn/go-my-blog/share/utils"
)

var (
	// CUser User struct instance
	CUser = &models.User{}
)

func removePassToken(ct *models.Token) (int, error) {
	updator := bson.M{"$set": bson.M{"isDelete": true}}
	selector := bson.M{
		"userId":    ct.UserID,
		"createdAt": bson.M{"$gte": utils.GetStartTimeOfDay(time.Now())},
		"updatedAt": bson.M{"$lt": ct.UpdatedAt},
		"isDelete":  false,
	}

	CToken := &models.Token{}
	count, err := CToken.UpdateAll(selector, updator)
	return count, err
}

// Login check user is in database
func (BlogService) Login(c *gin.Context) {
	VUser := &models.User{}
	c.BindJSON(VUser)

	email := VUser.Email
	password := VUser.Password

	err := CUser.GetByEmail(email)

	if err != nil {
		c.JSON(http.StatusBadRequest, &gin.H{
			"message": map[string]interface{}{"email": "The user does not exist"},
			"name":    "Form tip error",
			"status":  http.StatusBadRequest,
			"code":    0,
		})
		return
	}

	isValid := CUser.CheckPassword(password)

	if isValid {
		CToken := &models.Token{
			AccessToken: utils.GenerateUUIDString(),
			UserID:      CUser.ID,
		}
		err = CToken.Create()

		if err == nil {
			VTOKEN := &models.VToken{}
			VTOKEN.Copy(CToken, CUser)
			go removePassToken(CToken)
			c.JSON(http.StatusOK, VTOKEN)
		} else {
			c.JSON(http.StatusInternalServerError, &gin.H{
				"message": "Server error",
				"name":    "Server error",
				"status":  http.StatusInternalServerError,
				"code":    0,
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, &gin.H{
			"message": map[string]interface{}{"password": "Password is incorrect"},
			"name":    "Form tip error",
			"status":  http.StatusBadRequest,
			"code":    0,
		})
	}
}
