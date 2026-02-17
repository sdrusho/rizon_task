package userservice

import (
	"context"
	"fmt"
	"log"
	"ms-feedback/internal/config"
	"ms-feedback/internal/handler/auth"
	"ms-feedback/internal/model"
	userrepo "ms-feedback/internal/repos/user"
	"ms-feedback/internal/service/email"
	feedbackservice "ms-feedback/internal/service/feedback"
)

type UserService interface {
	CreateUser(ctx context.Context, email string, name, deviceId *string) (model.User, error)
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	GetUserById(ctx context.Context, userID string) (model.User, error)
}

const (
	REQUESTED string = "REQUESTED"
	SUBMITTED string = "SUBMITTED"
)

type userServiceImpl struct {
	userRepo        userrepo.UserRepository
	config          config.Config
	feedbackService feedbackservice.FeedbackService
	emailService    *email.EmailService
	authHandler     *auth.AuthHandler
}

func NewUserService(repo userrepo.UserRepository, cfg config.Config, feedbackService feedbackservice.FeedbackService, emailService *email.EmailService, authHandler *auth.AuthHandler) UserService {

	userService := &userServiceImpl{
		userRepo:        repo,
		config:          cfg,
		feedbackService: feedbackService,
		emailService:    emailService,
		authHandler:     authHandler,
	}
	return userService
}

func (s *userServiceImpl) CreateUser(ctx context.Context, email string, name, deviceId *string) (model.User, error) {
	nameString := ""
	if name != nil {
		nameString = *name
	}
	createdUser, err := s.userRepo.CreateUser(ctx, email, nameString, REQUESTED, deviceId)
	if err != nil {
		return model.User{}, fmt.Errorf("error creating user: %w", err)
	}
	s.SendSignupEmail(ctx, email)
	return model.ToUser(createdUser), nil
}

func (s *userServiceImpl) SendSignupEmail(ctx context.Context, email string) error {
	go func(email string) {
		if err := s.emailService.SendSignupEmail(email); err != nil {
			log.Printf("Failed to send contact us email to %s:", email)
		}
	}(email)

	return nil
}

func (s *userServiceImpl) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	currentUser, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return model.User{}, fmt.Errorf("no user found with this email")
	}
	return currentUser, nil
}

func (s *userServiceImpl) GetUserById(ctx context.Context, userId string) (model.User, error) {

	existingFeedback, _ := s.feedbackService.GetFeedbackByUserId(ctx, userId)
	if (model.Feedback{} != existingFeedback) {
		return model.User{}, fmt.Errorf("user feedback submitted")
	}

	currentUser, err := s.userRepo.GetUserById(ctx, userId)
	if err != nil {
		return model.User{}, fmt.Errorf("no user found with this userId")
	}
	token, refresh, err := s.authHandler.CreateToken(currentUser.UserID)
	if err != nil {
		return model.User{}, fmt.Errorf("no user found with this email")
	}
	return model.SetAuthToUser(token, refresh, currentUser), nil
}
