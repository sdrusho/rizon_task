package apperrors

import (
	"errors"
)

var (
	ErrUserNotFoundInDB = errors.New("no user found with this email")
	ErrWithPassword     = errors.New("Invalid current password")
	ErrUserNotFound     = errors.New("no users found")
	ErrUserNotFoundByID = errors.New(("no user found with this user ID"))
)
