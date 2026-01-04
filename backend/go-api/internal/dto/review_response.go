package dto

type ReviewResponse struct {
	ID     int64   `json:"id"`
	IsFake bool    `json:"is_fake"`
	Score  float64 `json:"score"`
}
