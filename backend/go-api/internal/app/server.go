package app

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"path/filepath"

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

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	projectRoot := filepath.Join(wd, "..", "..")

	log.Println(projectRoot + "/frontend/public")
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

	mux.Handle("/", http.FileServer(http.Dir(projectRoot+"/frontend/public")))
	mux.Handle("/src/", http.StripPrefix("/src/", http.FileServer(http.Dir(projectRoot+"/frontend/src"))))

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
