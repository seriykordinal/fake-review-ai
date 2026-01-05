package repositories

import (
	"context"
	"database/sql"

	"github.com/seriykordinal/fake-review-ai/internal/models"
)

type ReviewRepository interface {
	Create(ctx context.Context, r *models.Review) (int64, error)
	UpdateResult(ctx context.Context, id int64, score float64, isFake bool) error
}

type reviewRepository struct {
	db *sql.DB
}

func NewReviewRepository(db *sql.DB) ReviewRepository {
	return &reviewRepository{db: db}
}
