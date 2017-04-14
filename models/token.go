package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	ext "hankliu.com.cn/go-my-blog/share/extension"
	"hankliu.com.cn/go-my-blog/share/utils"
)

const (
	// C_TOKEN token table name
	C_TOKEN = "token"
)

// Token token table structural
type Token struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	AccessToken string        `bson:"accesstoken"`
	ExpireTime  time.Time     `bson:"expireTime"`
	UserID      bson.ObjectId `bson:"userId"`
	CreatedAt   time.Time     `bson:"createdAt"`
	UpdatedAt   time.Time     `bson:"updatedAt"`
	IsDeleted   bool          `bson:"isDelete"`
}

// Create create token
func (t *Token) Create() error {
	t.ID = bson.NewObjectId()
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
	t.ExpireTime = utils.GetEndTimeOfDay(time.Now())
	t.IsDeleted = false

	return ext.MongoDBRepository.Inset(C_TOKEN, t)
}

// GetByCondition get token by condition
func (t *Token) GetByCondition(condition bson.M) error {
	return ext.MongoDBRepository.FindOne(C_TOKEN, condition, t)
}

// UpdateAll update tokens
func (t *Token) UpdateAll(selector bson.M, updator bson.M) (int, error) {
	return ext.MongoDBRepository.UpdateAll(C_TOKEN, selector, updator)
}
