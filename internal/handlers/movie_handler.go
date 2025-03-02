package handlers

import (
	"net/http"
	"strconv"

	"lion_parcel/internal/models"
	"lion_parcel/internal/services"

	"github.com/labstack/echo/v4"
)

type MovieHandler struct {
	MovieService *services.MovieService
}

func NewMovieHandler(movieService *services.MovieService) *MovieHandler {
	return &MovieHandler{MovieService: movieService}
}

func (h *MovieHandler) ListMovies(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	movies, err := h.MovieService.ListMovies(page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, movies)
}

func (h *MovieHandler) SearchMovies(c echo.Context) error {
	query := c.QueryParam("q")

	movies, err := h.MovieService.SearchMovies(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, movies)
}

func (h *MovieHandler) TrackView(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.MovieService.TrackView(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "View tracked"})
}

func (h *MovieHandler) CreateMovie(c echo.Context) error {
	var movie models.Movie
	if err := c.Bind(&movie); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	if err := h.MovieService.CreateMovie(&movie); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Movie created successfully"})
}

func (h *MovieHandler) UpdateMovie(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var movie models.Movie
	if err := c.Bind(&movie); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	movie.ID = uint(id)
	if err := h.MovieService.UpdateMovie(&movie); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Movie updated successfully"})
}

func (h *MovieHandler) GetMostViewedMovies(c echo.Context) error {
	movies, err := h.MovieService.GetMostViewedMovies()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, movies)
}

func (h *MovieHandler) GetMostViewedGenres(c echo.Context) error {
	genres, err := h.MovieService.GetMostViewedGenres()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, genres)
}
