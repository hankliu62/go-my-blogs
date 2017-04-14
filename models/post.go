package models

import (
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"

	"hankliu.com.cn/go-my-blog/codes"
	"hankliu.com.cn/go-my-blog/constants"
	ext "hankliu.com.cn/go-my-blog/share/extension"
	"hankliu.com.cn/go-my-blog/share/utils"
)

const (
	// C_POST post table name
	C_POST = "post"
)

// Comment comment table structural
type Comment struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Content   string        `bson:"content"`
	UserID    bson.ObjectId `bson:"userId,omitempty"`
	PostID    bson.ObjectId `bson:"postId"`
	CreatedAt time.Time     `bson:"createdAt"`
	UpdatedAt time.Time     `bson:"updatedAt"`
	IsDeleted bool          `bson:"isDelete"`
}

// Post post table structural
type Post struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Title     string        `bson:"title"`
	Content   string        `bson:"content"`
	Tags      []string      `bson:"tags"`
	UserID    bson.ObjectId `bson:"userId"`
	CreatedAt time.Time     `bson:"createdAt"`
	UpdatedAt time.Time     `bson:"updatedAt"`
	Comments  []Comment     `bson:"comments"`
	IsDeleted bool          `bson:"isDelete"`
}

// GetByPk get post by id
func (p *Post) GetByPk(ID bson.ObjectId) error {
	return ext.MongoDBRepository.FindByPK(C_POST, ID, p)
}

// GetByCondition get post by condition
func (p *Post) GetByCondition(condition bson.M) error {
	condition["isDelete"] = false
	return ext.MongoDBRepository.FindOne(C_POST, condition, p)
}

// GetAllByCondition get posts by condition
func (*Post) GetAllByCondition(condition bson.M) error {
	var posts []Post
	condition["isDelete"] = false
	return ext.MongoDBRepository.FindAll(C_POST, condition, []string{}, int(constants.MAX_LIMIT), &posts)
}

// GetAllByPagination get posts by pagination
func (*Post) GetAllByPagination(condition bson.M, page uint32, pageSize uint32, orderbys []string) ([]Post, int) {
	var posts []Post
	sortFields := utils.NormalizeOrderBy(orderbys)
	pageCondition := ext.PagingCondition{
		Selector:  condition,
		PageIndex: int(page),
		PageSize:  int(pageSize),
		Sortor:    sortFields,
	}
	total, _ := ext.MongoDBRepository.FindByPagination(C_POST, pageCondition, &posts)
	return posts, total
}

// Create creat a post into db
func (p *Post) Create() error {
	p.ID = bson.NewObjectId()
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	p.Comments = []Comment{}
	p.IsDeleted = false

	if err := p.validatePostParams(); err != nil {
		return err
	}

	return ext.MongoDBRepository.Inset(C_POST, p)
}

func (p *Post) validatePostParams() error {
	if len(strings.TrimSpace(p.Title)) == 0 || len(strings.TrimSpace(p.Content)) == 0 || len(p.Tags) == 0 {
		return codes.NewError(codes.CommonMissingRequiredFields)
	}

	return nil
}

// Update update post
func (p *Post) Update(selector bson.M, updator bson.M) error {
	return ext.MongoDBRepository.UpdateOne(C_POST, selector, updator)
}

// AppendComment add comment for post
func (p *Post) AppendComment(c *Comment) error {
	c.ID = bson.NewObjectId()
	c.PostID = p.ID
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	c.IsDeleted = false

	selector := bson.M{"_id": p.ID}
	updator := bson.M{"$push": bson.M{
		"comments": *c,
	}}
	return ext.MongoDBRepository.UpdateOne(C_POST, selector, updator)
}

// RemoveComment remove some commet for post
func (p *Post) RemoveComment(c *Comment) error {
	selector := bson.M{"_id": p.ID}
	updator := bson.M{"$pull": bson.M{
		"comments": bson.M{
			"_id":    c.ID,
			"userId": c.UserID,
		},
	}}
	return ext.MongoDBRepository.UpdateOne(C_POST, selector, updator)
}

// Delete delete a post
func (p *Post) Delete(selector bson.M) error {
	return ext.MongoDBRepository.RemoveOne(C_POST, selector)
}

// Copy copy vo model post into po model
func (c *Comment) Copy(vc *VComment) {
	c.ID = bson.ObjectIdHex(vc.ID)
	c.Content = vc.Content
	c.UserID = bson.ObjectIdHex(vc.UserID)
	c.CreatedAt = vc.CreatedAt
	c.UpdatedAt = vc.UpdatedAt
}
