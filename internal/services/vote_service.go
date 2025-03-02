package services

import (
	"lion_parcel/internal/models"
	"lion_parcel/internal/repositories"
)

type VoteService struct {
	VoteRepo *repositories.VoteRepository
}

func NewVoteService(voteRepo *repositories.VoteRepository) *VoteService {
	return &VoteService{VoteRepo: voteRepo}
}

func (s *VoteService) Vote(userID, movieID uint) error {
	if s.VoteRepo.Exists(userID, movieID) {
		return nil // Already voted
	}
	vote := models.Vote{UserID: userID, MovieID: movieID}
	return s.VoteRepo.Create(&vote)
}

func (s *VoteService) Unvote(userID, movieID uint) error {
	return s.VoteRepo.Delete(userID, movieID)
}

func (s *VoteService) GetUserVotes(userID uint) ([]models.Vote, error) {
	return s.VoteRepo.GetUserVotes(userID)
}
