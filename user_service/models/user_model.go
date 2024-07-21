package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Username         string         `gorm:"unique not null" json:"username"`
	Password         string         `gorm:"not null" json:"password"`
	Email            string         `gorm:"unique not null" json:"email"`
	FullName         string         `gorm:"not null" json:"fullName"`
	Avatar           string         `json:"avatar"`
	Intersets        pq.Int32Array  `gorm:"type:int[]; not null" json:"interests"`
	Communities      pq.Int32Array  `gorm:"type:int[]" json:"communities"`
	Posts            pq.Int32Array  `gorm:"type:int[]" json:"posts"`
	Rating           int            `json:"rating"`
	Level            int            `json:"level"`
	AboutMe          string         `json:"about_me"`
	Links            pq.StringArray `gorm:"type:text[]" json:"links"`
	PublicProfileURL string         `json:"publicProfileUrl"`
}

type UserDto struct {
	ID               uint           `json:"id"`
	Username         string         `json:"username"`
	Email            string         `json:"email"`
	FullName         string         `json:"fullName"`
	Avatar           string         `json:"avatar"`
	Intersets        pq.Int32Array  `json:"interests"`
	Communities      pq.Int32Array  `json:"communities"`
	Posts            pq.Int32Array  `json:"posts"`
	Rating           int            `json:"rating"`
	Level            int            `json:"level"`
	AboutMe          string         `json:"about_me"`
	Links            pq.StringArray `json:"links"`
	PublicProfileURL string         `json:"publicProfileUrl"`
}

// convert user model to user dto
func (userModel UserModel) ToUserDto() UserDto {
	return UserDto{
		ID:               userModel.ID,
		Username:         userModel.Username,
		Email:            userModel.Email,
		FullName:         userModel.FullName,
		Avatar:           userModel.Avatar,
		Intersets:        userModel.Intersets,
		Communities:      userModel.Communities,
		Posts:            userModel.Posts,
		Rating:           userModel.Rating,
		Level:            userModel.Level,
		AboutMe:          userModel.AboutMe,
		Links:            userModel.Links,
		PublicProfileURL: userModel.PublicProfileURL,
	}
}
