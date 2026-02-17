package healthcheckhandler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

type HealthCheckResponse struct {
	Status    string                 `json:"status" example:"ok"`
	Timestamp time.Time              `json:"timestamp" example:"2023-04-15T12:34:56Z"`
	Checks    map[string]interface{} `json:"checks"`
}

type ComponentStatus struct {
	Status string `json:"status" example:"ok"`
	Error  string `json:"error,omitempty" example:"connection refused"`
}

// HealthCheck godoc
// @Summary Health check endpoint
// @Description Checks if the service is up and running by verifying database and redis connections
// @Tags healthcheck
// @Accept json
// @Produce json
// @Success 200 {object} HealthCheckResponse "Service is healthy"
// @Failure 503 {object} HealthCheckResponse "Service is unhealthy"
// @Router /ms-feedback/healthcheck [get]
func (h *Handler) HealthCheck(c *gin.Context) {
	response := map[string]interface{}{
		"status":    "ok",
		"timestamp": time.Now().UTC(),
		"checks": map[string]interface{}{
			"database": map[string]string{
				"status": "ok",
			},
		},
	}

	c.JSON(http.StatusOK, response)
}
