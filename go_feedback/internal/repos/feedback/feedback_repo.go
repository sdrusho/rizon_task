package feedbackrepo

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"ms-feedback/internal/db/generated/feedback"
	"ms-feedback/internal/model"
	"ms-feedback/pkg/utils"

	"github.com/google/uuid"
)

type FeedbackRepository interface {
	CreateFeedback(ctx context.Context, userId uuid.UUID, isEnjoying, isLeaveReview bool, comments string) (feedback.RizonDbFeedback, error)
	GetFeedbackByUserId(ctx context.Context, userID string) (model.Feedback, error)
}

type feedbackRepositoryImpl struct {
	queries *feedback.Queries
	db      *sql.DB
}

func NewFeedbackRepository(db *sql.DB) FeedbackRepository {
	return &feedbackRepositoryImpl{
		queries: feedback.New(db),
		db:      db,
	}
}

func (r *feedbackRepositoryImpl) CreateFeedback(ctx context.Context, userId uuid.UUID, isEnjoying, isLeaveReview bool, comments string) (feedback.RizonDbFeedback, error) {

	params := feedback.CreateFeedbackParams{
		Userid:        userId,
		Isenjoying:    isEnjoying,
		Isleavereview: isLeaveReview,
		Comments:      comments,
	}

	createdFeedback, err := r.queries.CreateFeedback(ctx, params)
	if err != nil {
		return feedback.RizonDbFeedback{}, fmt.Errorf("failed to create feedback: %w", err)
	}

	return createdFeedback, nil
}

func (r *feedbackRepositoryImpl) GetFeedbackByUserId(ctx context.Context, userID string) (model.Feedback, error) {
	uuidUserID, err := utils.ParseAndConvertUUID(userID)
	if err != nil {
		return model.Feedback{}, fmt.Errorf("failed to parse userId: %w", err)
	}

	foundFeedback, err := r.queries.GetFeedbackByUserID(ctx, uuidUserID.UUID)
	if err != nil {
		log.Printf("Get user by ID %s, returned %v\n", userID, err.Error())
		return model.Feedback{}, fmt.Errorf("database operation failed: %w", err)
	}
	return model.ToFeedbackFromRow(foundFeedback), nil
}
