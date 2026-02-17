package feedbackhandler

import (
	userservice "ms-feedback/internal/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService userservice.UserService
}

func NewUserHandler(userService userservice.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// CreateUser post a new signup user
// @Summary Post a new signup user
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.User true "User data"
// @Success 201 {object} model.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /ms-feedback/user-signup [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req struct {
		Email    string  `json:"email"`
		Name     *string `json:"name"`
		DeviceId *string `json:"deviceId"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := h.userService.CreateUser(c, req.Email, req.Name, req.DeviceId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)

}

// GetUserById retrieves a user by ID
// @Summary Get signup user by ID
// @Tags users
// @Param userId path string true "User ID"
// @Produce json
// @Success 200 {object} model.User
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /ms-feedback/users/{userId} [get]
func (h *UserHandler) GetUserById(c *gin.Context) {
	userID := c.Param("userId")

	foundUser, err := h.userService.GetUserById(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch user"})
		return
	}

	c.JSON(http.StatusOK, foundUser)
}

// GetUserByEmail retrieves a user by email
// @Summary Get a signup user by email
// @Tags users
// @Produce json
// @Param email path string true "User Email"
// @Success 200 {object} model.User
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /ms-feedback/users/{email} [get]
func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")

	foundUser, err := h.userService.GetUserByEmail(c, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, foundUser)
}
