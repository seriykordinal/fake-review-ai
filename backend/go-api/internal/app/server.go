package app

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/seriykordinal/fake-review-ai/internal/config"
	"github.com/seriykordinal/fake-review-ai/internal/handlers"
	"github.com/seriykordinal/fake-review-ai/internal/repositories"
	"github.com/seriykordinal/fake-review-ai/internal/services"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config) (*Server, error) {

	db, err := sql.Open("postgres", cfg.PostgresDSN)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	repo := repositories.NewReviewRepository(db)
	mlClient := services.NewMLCLient(cfg.MLServiceURL)

	reviewService := services.NewReviewService(repo, mlClient)
	reviewHandler := handlers.NewReviewHandler(reviewService)

	mux := http.NewServeMux()
	mux.Handle("/reviews", reviewHandler)

	return &Server{
		httpServer: &http.Server{
			Addr:    ":" + cfg.ServerPort,
			Handler: mux,
		},
	}, nil
}

func (s *Server) Run() error {
	log.Println("Server run on " + s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}
