package server

import (
	database "ms-feedback/database"
	config "ms-feedback/internal/config"
	"ms-feedback/internal/handler/auth"
	"ms-feedback/internal/handler/feedback"
	healthcheckhandler "ms-feedback/internal/handler/healthcheck"
	"ms-feedback/internal/handler/ratelimit"
	userhandler "ms-feedback/internal/handler/user"
	feedbackrepo "ms-feedback/internal/repos/feedback"
	userrepo "ms-feedback/internal/repos/user"
	"ms-feedback/internal/service/email"
	feedbackservice "ms-feedback/internal/service/feedback"
	slackservice "ms-feedback/internal/service/slack"
	userservice "ms-feedback/internal/service/user"
	"ms-feedback/pkg/middleware"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Start() error {
	// Load auth service

	// Load configuration
	cfg := config.MustLoadConfig()

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		return err
	}
	defer db.Close()

	// Initialize Email Service

	emailService := email.NewEmailService(
		cfg.EmailHost,
		cfg.EmailPort,
		cfg.EmailUsername,
		cfg.EmailPassword,
		cfg.EmailFrom,
	)
	healthCheckHandler := healthcheckhandler.NewHandler()
	// Rate limiting
	window := config.NewSlidingWindow(int64(100), time.Minute, config.NewSlidingWindowInMemory(), clock.New(), 1e-9)
	rateLimitHandler := ratelimit.NewRateLimitHandler(cfg, *window)

	// mock slack service
	mock := &slackservice.MockSlackClient{}
	slackService := slackservice.SlackService{Slack: mock}

	feedbackRepo := feedbackrepo.NewFeedbackRepository(db)
	feedbackService := feedbackservice.NewFeedbackService(feedbackRepo, slackService, cfg)
	feedbackHandler := feedback.NewFeedbackHandler(feedbackService)

	// auth handler validate jwt access token
	authHandler := auth.NewAuthHandler(cfg)
	userRepo := userrepo.NewUserRepository(db)
	userService := userservice.NewUserService(userRepo, cfg, feedbackService, emailService, authHandler)
	userHandler := userhandler.NewUserHandler(userService)
	// Set up Gin router
	router := gin.Default()

	// Apply CORS middleware
	router.Use(middleware.CORS())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, func(c *ginSwagger.Config) {
		c.Title = "BasicGo API"
	}))

	// Define root route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ms-basic is running",
		})
	})
	router.Use(rateLimitHandler.ValidateRateLimit)
	router.POST("/ms-feedback/user-signup", userHandler.CreateUser)

	router.GET("/ms-feedback/users/:userId", userHandler.GetUserById)
	router.GET("/ms-feedback/user-signup/:email", userHandler.GetUserByEmail)

	router.Use(authHandler.ValidateBearer)
	router.POST("/ms-feedback/feedbacks", feedbackHandler.CreateFeedback)
	router.GET("/ms-feedback/healthcheck", healthCheckHandler.HealthCheck)
	return router.Run(":" + "8001")
}
