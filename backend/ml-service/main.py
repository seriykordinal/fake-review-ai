from fastapi import FastAPI
from pydantic import BaseModel
import random

app = FastAPI(title="Fake Review ML Service")

class PredictRequest(BaseModel):
    text: str


class PredictResponse(BaseModel):
    score: float
    is_fake: bool


@app.post("/predict", response_model=PredictResponse)
def predict(request: PredictRequest):
   
    #тут будет вызов нейронки


    score = round(random.uniform(0, 1), 2)


    is_fake = score > 0.7

    return PredictResponse(
        score=score,
        is_fake=is_fake
    )


@app.get("/health")
def health():
    return {"status": "ok"}
