# GoURL
A fast URL shortener written in Go using Redis and Gin. 

## Installation
$ go get github.com/solarcreature/GoURL

## How to Run

### Required
-Redis

### Run
___________________________________________________________________________________________________________________________________________________________
- $ redis-server
___________________________________________________________________________________________________________________________________________________________
- $ go run main.go
___________________________________________________________________________________________________________________________________________________________

### Requests

curl --request POST \
--data '{
    "long_url": "example-url",
    "user_id" : "example-userID"
}' \
  http://localhost:9808/create-short-url

### Response

{
    "message": "short url created successfully",
    "short_url": "http://localhost:9808/example"
}
