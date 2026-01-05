package services

import (
	"context"
	"log"

	"github.com/seriykordinal/fake-review-ai/internal/dto"
	"github.com/seriykordinal/fake-review-ai/internal/models"
	"github.com/seriykordinal/fake-review-ai/internal/repositories"
)

type ReviewService struct {
	repo repositories.ReviewRepository
	ml   MLClient
}

func NewReviewService(r repositories.ReviewRepository, ml MLClient) *ReviewService {
	return &ReviewService{repo: r, ml: ml}
}

func (s *ReviewService) ProcessReview(ctx context.Context, req dto.CreateReviewRequest) (*dto.ReviewResponse, error) {
	review := &models.Review{
		Text: req.Text,
	}

	id, err := s.repo.Create(ctx, review)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	score, isFake, err := s.ml.Analyze(ctx, req.Text)
	if err != nil {
		log.Print(err)

		return nil, err
	}

	err = s.repo.UpdateResult(ctx, id, score, isFake)
	if err != nil {
		log.Print(err)

		return nil, err
	}

	return &dto.ReviewResponse{
		ID:     id,
		Score:  score,
		IsFake: isFake,
	}, nil
}
