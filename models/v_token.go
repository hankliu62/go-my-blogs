package models

import (
	"time"
)

// VToken token table view model
type VToken struct {
	ID          string    `json:"id"`
	AccessToken string    `json:"accesstoken"`
	ExpireTime  time.Time `json:"expireTime"`
	User        *VUser    `json:"user"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// Copy copy po model token into vo model
func (vt *VToken) Copy(t *Token, u *User) {
	vt.ID = t.ID.Hex()
	vt.AccessToken = t.AccessToken
	vt.ExpireTime = t.ExpireTime
	vt.CreatedAt = t.CreatedAt
	vt.UpdatedAt = t.UpdatedAt
	vu := &VUser{}
	vu.Copy(u)
	vt.User = vu
}
