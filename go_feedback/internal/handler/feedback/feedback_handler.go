package feedback

import (
	"ms-feedback/internal/service/feedback"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FeedbackHandler struct {
	feedbackService feedbackservice.FeedbackService
}

func NewFeedbackHandler(feedbackService feedbackservice.FeedbackService) *FeedbackHandler {
	return &FeedbackHandler{
		feedbackService: feedbackService,
	}
}

// CreateFeedback post a new feedback user
// @Summary Post a new signup feedback
// @Tags feedback
// @Accept json
// @Produce json
// @Param user body model.Feedback true "User data"
// @Success 201 {object} model.Feedback
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /ms-feedback/feedbacks [post]
func (h *FeedbackHandler) CreateFeedback(c *gin.Context) {
	var req struct {
		IsEnjoying    bool   `json:"isEnjoying"`
		IsLeaveReview bool   `json:"isLeaveReview"`
		Comments      string `json:"comments"`
		UserId        string `json:"userId"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdFeedback, err := h.feedbackService.CreateFeedback(c, req.UserId, req.IsEnjoying, req.IsLeaveReview, req.Comments)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdFeedback)

}
