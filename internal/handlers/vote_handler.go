package handlers

import (
	"net/http"
	"strconv"

	"lion_parcel/internal/services"

	"github.com/labstack/echo/v4"
)

type VoteHandler struct {
	VoteService *services.VoteService
}

func NewVoteHandler(voteService *services.VoteService) *VoteHandler {
	return &VoteHandler{VoteService: voteService}
}

func (h *VoteHandler) Vote(c echo.Context) error {
	userID := c.Get("user_id").(uint)
	movieID, err := strconv.Atoi(c.Param("movie_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid movie ID"})
	}

	if err := h.VoteService.Vote(userID, uint(movieID)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Voted successfully"})
}

func (h *VoteHandler) Unvote(c echo.Context) error {
	userID := c.Get("user_id").(uint)
	movieID, err := strconv.Atoi(c.Param("movie_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid movie ID"})
	}

	if err := h.VoteService.Unvote(userID, uint(movieID)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Unvoted successfully"})
}

func (h *VoteHandler) GetUserVotes(c echo.Context) error {
	userID := c.Get("user_id").(uint)

	votes, err := h.VoteService.GetUserVotes(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, votes)
}
