package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
	ext "hankliu.com.cn/go-my-blog/share/extension"
	"hankliu.com.cn/go-my-blog/share/utils"
)

const (
	// C_USER user table name
	C_USER = "user"
)

// User user table structural
type User struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	Name         string        `bson:"name"`
	Gender       string        `bson:"gender"`
	Introduction string        `bson:"introduction"`
	Email        string        `bson:"email"`
	Password     string        `bson:"password"`
	Salt         string        `bson:"salt"`
	CreatedAt    time.Time     `bson:"createdAt"`
	UpdatedAt    time.Time     `bson:"updatedAt"`
	IsDeleted    bool          `bson:"isDelete"`
}

// GetByEmail get user by email from db
func (u *User) GetByEmail(email string) error {
	condition := bson.M{
		"email": email,
	}

	return ext.MongoDBRepository.FindOne(C_USER, condition, u)
}

// GetByPK get user by userId
func (u *User) GetByPK(id bson.ObjectId) error {
	return ext.MongoDBRepository.FindByPK(C_USER, id, u)
}

// CheckPassword validate the user has correct password
func (u *User) CheckPassword(password string) bool {
	return utils.Md5Encode(password) == u.Password
}

// GetEncodePassword get encode password
func (*User) GetEncodePassword(password string) string {
	return utils.Md5Encode(password)
}

// Create create user
func (u *User) Create() error {
	u.ID = bson.NewObjectId()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	u.IsDeleted = false

	if len(u.Password) > 0 {
		u.Password = u.GetEncodePassword(u.Password)
		u.Salt = utils.RandomStr(6)
	}

	return ext.MongoDBRepository.Inset(C_USER, u)
}

// Update update user
func (*User) Update(selector bson.M, updator bson.M) error {
	return ext.MongoDBRepository.UpdateOne(C_USER, selector, updator)
}
