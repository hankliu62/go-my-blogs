package models

import "time"

// VComment post table comment property view model
type VComment struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	UserID    string    `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// VPost post table view model
type VPost struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	Tags      []string   `json:"tags"`
	UserID    string     `json:"userId"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	Comments  []VComment `json:"comments"`
}

// Copy copy po model post into vo model
func (vc *VComment) Copy(c *Comment) {
	vc.ID = c.ID.Hex()
	vc.Content = c.Content
	vc.UserID = c.UserID.Hex()
	vc.CreatedAt = c.CreatedAt
	vc.UpdatedAt = c.UpdatedAt
}

// Copy copy po model post into vo model
func (vp *VPost) Copy(p *Post) {
	vp.ID = p.ID.Hex()
	vp.Title = p.Title
	vp.Content = p.Content
	vp.Tags = p.Tags
	vp.UserID = p.UserID.Hex()
	vp.CreatedAt = p.CreatedAt
	vp.UpdatedAt = p.UpdatedAt
	vp.Comments = []VComment{}

	for _, c := range p.Comments {
		uv := &VComment{}
		uv.Copy(&c)
		vp.Comments = append(vp.Comments, *uv)
	}
}
