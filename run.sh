#!/bin/bash

echo "Generating API KEY..."

# you can modify --data with your own info
API_KEY_VAL=$(curl -s \
-X POST \
--header 'Content-Type: application/json' \
--data '{
  "time_window": 1,
  "max_requests": 10,
  "blocked_duration": 60
}' http://127.0.0.1:8080/api-key | jq '.api_key')

echo "$API_KEY_VAL"

echo "Running docker compose..."

docker compose run \
  --rm go-cli-test \
  -url http://go-app:8080/hello-world-key \
  -m GET \
  -t 1 \
  -r 100 \
  -k "${API_KEY_VAL}"

echo "Execution ended"
