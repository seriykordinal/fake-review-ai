package repositories

import (
	"context"

	"github.com/seriykordinal/fake-review-ai/internal/models"
)

func (r *reviewRepository) Create(ctx context.Context, review *models.Review) (int64, error) {
	query := `
		INSERT INTO reviews (text)
		VALUES ($1)
		RETURNING id
	`
	var id int64
	err := r.db.QueryRowContext(ctx, query, review.Text).Scan(&id)
	return id, err
}

func (r *reviewRepository) UpdateResult(ctx context.Context, id int64, score float64, isFake bool) error {
	query := `
		UPDATE reviews
		SET score=$1, is_fake=$2
		WHERE id=$3
	`

	_, err := r.db.ExecContext(ctx, query, score, isFake, id)
	return err
}
