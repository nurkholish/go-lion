package handlers

import (
	"lion_parcel/internal/config"
	"lion_parcel/internal/repositories"
	"lion_parcel/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"lion_parcel/internal/middleware"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("Failed to load configuration")
	}

	// Initialize repositories
	movieRepo := repositories.NewMovieRepository(db)
	userRepo := repositories.NewUserRepository(db)
	voteRepo := repositories.NewVoteRepository(db)
	blacklistedTokenRepo := repositories.NewBlacklistedTokenRepository(db)

	// Initialize services
	movieService := services.NewMovieService(movieRepo)
	userService := services.NewUserService(userRepo, blacklistedTokenRepo)
	voteService := services.NewVoteService(voteRepo)

	// Initialize handlers
	movieHandler := NewMovieHandler(movieService)
	userHandler := NewUserHandler(userService, cfg.JWTSecret)
	voteHandler := NewVoteHandler(voteService)

	// Public Routes
	publicGroup := e.Group("/api")
	{
		publicGroup.POST("/auth/register", userHandler.Register)
		publicGroup.POST("/auth/login", userHandler.Login)
		publicGroup.GET("/movies", movieHandler.ListMovies)
		publicGroup.GET("/movies/search", movieHandler.SearchMovies)
		publicGroup.POST("/movies/:id/view", movieHandler.TrackView)
	}

	// Authenticated Routes
	authGroup := e.Group("/api/auth")
	authGroup.Use(middleware.NewAuthMiddleware(cfg.JWTSecret).Authenticate)
	{
		authGroup.POST("/logout", userHandler.Logout)
		authGroup.POST("/votes/:movie_id", voteHandler.Vote)
		authGroup.DELETE("/votes/:movie_id", voteHandler.Unvote)
		authGroup.GET("/votes", voteHandler.GetUserVotes)
	}

	// Admin Routes
	adminGroup := e.Group("/api/admin")
	adminGroup.Use(middleware.NewAuthMiddleware(cfg.JWTSecret).Authenticate)
	adminGroup.Use(middleware.NewAdminMiddleware().Authorize)
	{
		adminGroup.POST("/movies", movieHandler.CreateMovie)
		adminGroup.PUT("/movies/:id", movieHandler.UpdateMovie)
		adminGroup.GET("/movies/popular", movieHandler.GetMostViewedMovies)
		adminGroup.GET("/genres/popular", movieHandler.GetMostViewedGenres)
	}

	// Health Check
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})
}
