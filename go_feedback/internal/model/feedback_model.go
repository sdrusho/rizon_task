package model

import (
	"ms-feedback/internal/db/generated/feedback"
	"ms-feedback/pkg/utils"
)

type Feedback struct {
	ID            int64  `json:"id"`
	UserID        string `json:"userId"`
	IsEnjoying    bool   `json:"isEnjoying"`
	IsLeaveReview bool   `json:"isLeaveReview"`
	Comments      string `json:"comments"`
	CreatedAt     string `json:"createdAt"`
}

func ToFeedback(f feedback.RizonDbFeedback) Feedback {
	return Feedback{
		ID:            f.ID,
		UserID:        f.Userid.String(),
		IsLeaveReview: f.Isleavereview,
		IsEnjoying:    f.Isenjoying,
		Comments:      f.Comments,
		CreatedAt:     utils.FormatTimestamp(f.Createdat),
	}
}

func ToFeedbackFromRow(f feedback.RizonDbFeedback) Feedback {
	return Feedback{
		ID:            f.ID,
		UserID:        f.Userid.String(),
		IsLeaveReview: f.Isleavereview,
		IsEnjoying:    f.Isenjoying,
		Comments:      f.Comments,
		CreatedAt:     utils.FormatTimestamp(f.Createdat),
	}
}
