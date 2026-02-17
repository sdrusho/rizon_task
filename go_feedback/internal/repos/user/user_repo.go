package userrepo

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"ms-feedback/internal/db/generated/user"
	"ms-feedback/internal/model"
	"ms-feedback/pkg/utils"
	"strings"
)

type UserRepository interface {
	CreateUser(ctx context.Context, email, name, status string, deviceId *string) (user.RizonDbUser, error)
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	GetUserById(ctx context.Context, userID string) (model.User, error)
}

type userRepositoryImpl struct {
	queries *user.Queries
	db      *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{
		queries: user.New(db),
		db:      db,
	}
}

func (r *userRepositoryImpl) CreateUser(ctx context.Context, email, name, status string, deviceId *string) (user.RizonDbUser, error) {
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)

	statusNullString := utils.StringToNullString(status)
	deviceIdNullString := utils.StringPointerToNullString(deviceId)
	params := user.CreateUserParams{
		Email:    email,
		Name:     name,
		Status:   statusNullString,
		Deviceid: deviceIdNullString,
	}

	createdUser, err := r.queries.CreateUser(ctx, params)
	if err != nil {
		return user.RizonDbUser{}, fmt.Errorf("failed to create user: %w", err)
	}

	return createdUser, nil
}

func (r *userRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	foundUser, err := r.queries.GetUserByEmail(ctx, strings.TrimSpace(email))
	if err != nil {
		log.Printf("Get user by email %s, returned %v\n", email, err.Error())
		return model.User{}, fmt.Errorf("database operation failed: %w", err)
	}
	log.Printf("Found user %v\n", foundUser)
	return model.ToUserFromEmailRow(foundUser), nil
}

func (r *userRepositoryImpl) GetUserById(ctx context.Context, userID string) (model.User, error) {
	uuidUserID, err := utils.ParseAndConvertUUID(userID)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to parse userId: %w", err)
	}

	foundUser, err := r.queries.GetUserByID(ctx, uuidUserID.UUID)
	if err != nil {
		log.Printf("Get user by ID %s, returned %v\n", userID, err.Error())
		return model.User{}, fmt.Errorf("database operation failed: %w", err)
	}
	return model.ToUserFromRow(foundUser), nil
}
