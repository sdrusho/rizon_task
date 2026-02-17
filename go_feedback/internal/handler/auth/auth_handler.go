package auth

import (
	"errors"
	"ms-feedback/internal/config"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	config config.Config
}

func NewAuthHandler(config config.Config) *AuthHandler {
	return &AuthHandler{
		config: config,
	}
}

type UnsignedResponse struct {
	Message interface{} `json:"message"`
}

// ValidateBearer - Extract the bearer token from the HTTP incoming headers and validate signing and expiry
// Returns -
// 400 if Authorization is missing with system error message
// 400 if JWT token is un-parsable message :'bad jwt token'
// 500 if there are no valid claims: 'unable to parse claims'
// 412 if the token is not an access token (e.g. refresh token)
// context.Next if all ok
func (h *AuthHandler) ValidateBearer(c *gin.Context) {

	// TODO: This should be removed once the front end supports full authentication
	if h.config.DisableAuth {
		c.Next()
		return
	}

	expectedType := c.Param("expectedType")
	if len(expectedType) == 0 {
		expectedType = "access"
	}

	jwtToken, err := ExtractAuthorizationHeader(c.GetHeader("Authorization"), "Bearer")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{Message: err.Error()})
		return
	}

	token, err := config.ParseToken(h.config.BearerSignerKey, jwtToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{Message: "unable to parse"})
		return
	}

	if token.Type != expectedType {
		c.AbortWithStatusJSON(http.StatusPreconditionFailed, UnsignedResponse{Message: "Invalid token"})
		return
	}

	// ** Store all the claims in the `context` in case any downstream actions require them **
	c.Set("claimUserId", token.User)
	c.Next()
}

func ExtractAuthorizationHeader(header string, authType string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}

	authHeader := strings.Split(header, " ")
	if len(authHeader) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	if authHeader[0] != authType {
		return "", errors.New("wrong type of authorization header")
	}

	return authHeader[1], nil
}

func (h *AuthHandler) CreateToken(userId string) (string, string, error) {

	tokenExp := time.Now()
	tokenExp = tokenExp.AddDate(0, 0, h.config.AccessExpiryDays)

	tokenStr, err := config.CreateToken(h.config.BearerSignerKey, "access", userId, tokenExp)
	if err != nil {
		return "", "", err
	}

	tokenExp = tokenExp.AddDate(0, 0, h.config.RefreshExpiryDays)

	refreshTokenStr, err := config.CreateToken(h.config.BearerSignerKey, "refresh", userId, tokenExp)
	if err != nil {
		return "", "", err
	}

	return tokenStr, refreshTokenStr, nil
}
