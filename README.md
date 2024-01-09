# FullCycle ratelimiter

Rate Limiter challenge

# 🏔️ Env Variables

## IP configuration

- We have this env file called `./env.json.example`:

```json
{
  "app": {
    "port": "8080"
  },
  "redis": {
    "db": 0,
    "host": "redis",
    "port": "6379"
  },
  "rate_limiter": {
    "by_ip": {
      "time_window": 1,
      "max_requests": 10,
      "blocked_duration": 60
    }
  }
}
```

where:

- `time_window`: is the value of **SECONDS** that we can allow the maximum amount of request.
- `max_requests`: is the maximum amount of request that we can allow each `time_window` seconds.
- `blocked_duration`: is the number of **SECONDS** that the IP is blocked, so we do not allow requests from this IP.

## API Token configuration

In the case of API token, we can create a new token by using this http request:

```http request
POST http://127.0.0.1:8080/api-key
Content-Type: application/json

{
  "time_window": 1,
  "max_requests": 100,
  "blocked_duration": 600
}
```

Each token has different configuration.

# 🚀 Starting the application!

- Execute the command `make prepare`, this will copy the default configuration.
- Before you execute the docker compose command, you can edit the `env.json` configuration file with your onw values.
- Finally, you can execute using the command `make run`

# 🧪 How can I test?

To learn more about the CLI using for test, you can run this:

```shell
docker compose run --rm go-cli-test -h
```

## Testing with IP only

Using docker you can run this command:

```shell
docker compose run --rm go-cli-test -url http://go-app:8080/hello-world -m GET -t 1 -r 10
```

## Testing with API key

- First, you need to execute this http request:
   ```http request
    POST http://127.0.0.1:8080/api-key
    Content-Type: application/json
    
    {
      "time_window": 1,
      "max_requests": 10,
      "blocked_duration": 60
    }
   ```
- You will receive the response something like this (the api-key will be different)
   ```json
    {
      "api-key": "c6f7363326f62f2483756447a963f2369a0dd5e90b7e8a36c32bc1a62ed38f51"
    }
   ```
- Copy the value of `api-key` and execute the following command, you should put your own token in the `-k` flag.
   ```shell
   docker compose run --rm go-cli-test -url http://go-app:8080/hello-world-key -m GET -t 1 -r 10 -k c6f7363326f62f2483756447a963f2369a0dd5e90b7e8a36c32bc1a62ed38f51
   ```

- If you need to execute at the same time, please execute this shell script file: `./run.sh`

# 💿 Redis DB

Here, we'll explain about the keys stored in redisDB. We have these keys:

- **<api-key>**: we have this to store the configuration of each API, we have this value as example
   key: `880d207159a7ac1a5a800eabbb310cf851e3c00cb5a2ff6e1ab9f38ce21bcc99`

  ```json
  {
    "max_requests": 10,
    "time_window": 1,
    "blocked_duration": 60
  }
  ```

- **rate:api-key_<api-key>**:
  
  ```json
  {
    "max_requests": 10,
    "time_window_sec": 1,
    "requests": [
      1704787618,
      1704787618,
      1704787618,
      1704787618,
      1704787618,
      1704787618,
      1704787618,
      1704787618,
      1704787618,
      1704787618,
      1704787618
    ]
  }
  ```
