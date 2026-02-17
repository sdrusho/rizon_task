package model

import (
	"ms-feedback/internal/db/generated/user"
	"ms-feedback/pkg/utils"
)

type User struct {
	ID           int64  `json:"id"`
	UserID       string `json:"userId"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	DeviceId     string `json:"deviceId"`
	Status       string `json:"status"`
	IsDeleted    *bool  `json:"isDeleted"`
	CreatedAt    string `json:"createdAt"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func ToUser(u user.RizonDbUser) User {
	return User{
		ID:        u.ID,
		UserID:    u.Userid.String(),
		Name:      u.Name,
		Email:     u.Email,
		DeviceId:  u.Deviceid.String,
		Status:    u.Status.String,
		CreatedAt: utils.FormatTimestamp(u.Createdat),
		IsDeleted: &u.Isdeleted,
	}
}

func SetAuthToUser(accessToken, refreshToken string, user User) User {
	user.AccessToken = accessToken
	user.RefreshToken = refreshToken
	return user
}

func ToUserFromRow(u user.GetUserByIDRow) User {
	return User{
		ID:        u.ID,
		UserID:    u.Userid.String(),
		Name:      u.Name,
		Email:     u.Email,
		Status:    u.Status.String,
		CreatedAt: utils.FormatTimestamp(u.Createdat),
		IsDeleted: &u.Isdeleted,
	}
}

func ToUserFromEmailRow(u user.GetUserByEmailRow) User {
	return User{
		ID:        u.ID,
		UserID:    u.Userid.String(),
		Name:      u.Name,
		Email:     u.Email,
		Status:    u.Status.String,
		CreatedAt: utils.FormatTimestamp(u.Createdat),
		IsDeleted: &u.Isdeleted,
	}
}
