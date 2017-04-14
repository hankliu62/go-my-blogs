package models

import "time"

// VUser user table view model
type VUser struct {
	ID           string    `json:"ids"`
	Name         string    `json:"name"`
	Gender       string    `json:"gender"`
	Introduction string    `json:"introduction"`
	Email        string    `json:"email"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// Copy copy po model token into vo model
func (vu *VUser) Copy(u *User) {
	vu.ID = u.ID.Hex()
	vu.Name = u.Name
	vu.Gender = u.Gender
	vu.Introduction = u.Introduction
	vu.Email = u.Email
	vu.CreatedAt = u.CreatedAt
	vu.UpdatedAt = u.UpdatedAt
}
