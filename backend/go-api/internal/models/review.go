package models

import "time"

type Review struct {
	ID        int64
	Text      string
	IsFake    bool
	Score     float64
	CreatedAt time.Time
}
