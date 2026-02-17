package ratelimit

import (
	"errors"
	"log"
	"ms-feedback/internal/config"
	"ms-feedback/internal/handler/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RateLimitHandler struct {
	config config.Config
	sl     config.SlidingWindow
}

func NewRateLimitHandler(config config.Config, sl config.SlidingWindow) *RateLimitHandler {

	return &RateLimitHandler{
		config: config,
		sl:     sl,
	}
}

func (h *RateLimitHandler) ValidateRateLimit(c *gin.Context) {

	// TODO: This should be removed once the front end supports full authentication
	if h.config.DisableAuth {
		c.Next()
		return
	}
	/*capacity := int64(3)
	rate := time.Minute
	clock := clock.New()
	epsilon := 1e-9
	window := config.NewSlidingWindow(capacity, rate, config.NewSlidingWindowInMemory(), clock, epsilon)*/
	time, err := h.sl.Limit(c)
	if err != nil {
		errors.New("error limit exhausted")
		c.AbortWithStatusJSON(http.StatusBadRequest, auth.UnsignedResponse{Message: "error limit exhausted"})
		return
	}
	log.Printf("time for rate limit: %v", time)
	c.Next()
}
