package services

import (
	"net/http"

	"strings"

	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/mgo.v2/bson"
	"hankliu.com.cn/go-my-blog/codes"
	"hankliu.com.cn/go-my-blog/constants"
	"hankliu.com.cn/go-my-blog/errors"
	"hankliu.com.cn/go-my-blog/models"
)

// UpdatePwdForm update password form struct
type UpdatePwdForm struct {
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

// UpdatePassword update user password
func (BlogService) UpdatePassword(c *gin.Context) {
	form := &UpdatePwdForm{}
	c.BindJSON(form)

	UserID, _ := c.Get("UserID")
	hexUserID := UserID.(bson.ObjectId)

	if err := validatePasswords(form); err != nil {
		c.JSON(http.StatusBadRequest, getStatusBadRequestResp(err))
		return
	}

	CUser = &models.User{}
	if err := CUser.GetByPK(hexUserID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"name":    "Server error",
			"status":  http.StatusInternalServerError,
			"code":    0,
		})
		return
	}

	if !CUser.CheckPassword(form.Password) {
		c.JSON(http.StatusBadRequest, getStatusBadRequestResp(codes.NewError(codes.OldPasswordNotMatched)))
		return
	}

	selector := bson.M{"_id": hexUserID}
	updator := bson.M{"$set": bson.M{"password": CUser.GetEncodePassword(form.NewPassword)}}

	if err := CUser.Update(selector, updator); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"name":    "Server error",
			"status":  http.StatusInternalServerError,
			"code":    0,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "change password success",
		"status":  http.StatusInternalServerError,
		"code":    0,
	})
}

func validatePasswords(pwdFrom *UpdatePwdForm) error {
	if len(strings.TrimSpace(pwdFrom.Password)) <= 0 || len(strings.TrimSpace(pwdFrom.NewPassword)) <= 0 {
		return codes.NewError(codes.CommonMissingRequiredFields)
	}

	if !constants.PasswordReg.MatchString(pwdFrom.NewPassword) {
		return codes.NewError(codes.InvalidPasswordFormat)
	}

	return nil
}

func getStatusBadRequestResp(err error) gin.H {
	bServerError, ok := err.(*errors.BServerError)

	resp := gin.H{
		"message": err.Error(),
		"name":    "Bad request",
		"status":  http.StatusBadRequest,
		"code":    0,
	}
	if ok {
		resp["code"] = bServerError.Code
	}

	return resp
}
