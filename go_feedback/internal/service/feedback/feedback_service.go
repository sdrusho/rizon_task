package feedbackservice

import (
	"context"
	"fmt"
	"ms-feedback/internal/config"
	"ms-feedback/internal/model"
	feedbackrepo "ms-feedback/internal/repos/feedback"
	slackservice "ms-feedback/internal/service/slack"
	"ms-feedback/pkg/utils"
)

type FeedbackService interface {
	CreateFeedback(ctx context.Context, userId string, isEnjoying, isLeaveReview bool, comments string) (model.Feedback, error)
	GetFeedbackByUserId(ctx context.Context, userId string) (model.Feedback, error)
}
type feedbackServiceImpl struct {
	feedbackRepo feedbackrepo.FeedbackRepository
	slackService slackservice.SlackService
	config       config.Config
}

func NewFeedbackService(repo feedbackrepo.FeedbackRepository, slackService slackservice.SlackService, cfg config.Config) FeedbackService {

	feedbackService := &feedbackServiceImpl{
		feedbackRepo: repo,
		slackService: slackService,
		config:       cfg,
	}
	return feedbackService
}

func (s *feedbackServiceImpl) CreateFeedback(ctx context.Context, userId string, isEnjoying, isLeaveReview bool, comments string) (model.Feedback, error) {
	userIdUUID, err := utils.ParseAndConvertUUID(userId)
	if err != nil {
		return model.Feedback{}, fmt.Errorf("invalid user id")
	}
	createdFeedback, err := s.feedbackRepo.CreateFeedback(ctx, userIdUUID.UUID, isEnjoying, isLeaveReview, comments)
	if err != nil {
		return model.Feedback{}, fmt.Errorf("error creating user: %w", err)
	}
	s.slackService.NotifyUser("U0AFJB2254L", createdFeedback.Comments)
	return model.ToFeedback(createdFeedback), nil
}

func (s *feedbackServiceImpl) GetFeedbackByUserId(ctx context.Context, userId string) (model.Feedback, error) {

	currentFeedback, err := s.feedbackRepo.GetFeedbackByUserId(ctx, userId)
	if err != nil {
		return model.Feedback{}, fmt.Errorf("no feedback found with this userId")
	}

	return currentFeedback, nil
}
