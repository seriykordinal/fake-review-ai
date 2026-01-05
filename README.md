# fake-review-ai
My service

## How to run?

First you need to run ml-service, then go-api

### Run of ml-service

From root of project (fake-review-ai) go to ml-service

``` 
cd backend/ml-service
```

Run .ps1 file to start service 
``` 
.\run.ps1
```


### Run of go-api

From root of project (fake-review-ai) go to go-api service

``` 
cd backend/go-api
```

Run .ps1 file to start service
``` 
.\run.ps1
```


## Test 

After run ml-service to be shure
```
 Invoke-RestMethod `
   -Uri http://localhost:8000/predict `
   -Method Post `
   -ContentType "application/json" `
   -Body '{"text":"Отличный товар!!!"}'
```

After run go-api and ml-service to be shure
```
Invoke-RestMethod `
   -Uri http://localhost:8080/reviews `
   -Method Post `
   -ContentType "application/json" `
   -Body '{"text": "Отличный товар, всем рекомендую!"}'
```