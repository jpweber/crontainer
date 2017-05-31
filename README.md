
# Crontainer - a container based cron system



## API

The API supports  
* getting a list of jobs in the system.  
    Done via a `HTTP GET` request
* Creating New jobs  
    Done via `HTTP POST` request with json body



### Example GET Request
```
curl "http://localhost:8675"
```

### Example POST Request
```
curl -X "POST" "http://localhost:8675" \
     -H "Content-Type: application/json; charset=utf-8" \
     -d $'{
  "CronPattern": "*/1 * * * *",
  "RunCommand": [
    "echo",
    "hello",
    "world"
  ],
  "State": 1,
  "ImageName": "alpine"
}'
```