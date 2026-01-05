$ErrorActionPreference = "Stop"

$env:SERVER_PORT = "8080"
$env:POSTGRES_DSN = "postgres://postgres:Shuter13@localhost:5432/fake-review-ai?sslmode=disable"
$env:ML_SERVICE_URL = "http://localhost:8000"

Write-Host "Starting Go API..."

go run ./cmd/main.go
