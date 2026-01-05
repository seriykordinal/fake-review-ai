$ErrorActionPreference = "Stop"

Write-Host "Starting ML service..."


Write-Host "Installing dependencies..."
pip install -r requirements.txt


Write-Host "Running server on http://localhost:8000"
uvicorn main:app --host 0.0.0.0 --port 8000
