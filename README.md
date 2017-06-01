
# Crontainer - a container based cron system



## API

The API supports  
* Getting a list of jobs in the system.  
    Done via a `HTTP GET` request
* Get info about a single job
    Done via a `HTTP GET` request
* Creating New jobs  
    Done via `HTTP POST` request with json body
* Deleting Jobs
    Done via `HTTP DELETE` request with resource in the url


## API Examples
### GET Jobs Request
```
curl "http://localhost:8675/jobs"
```

### GET Schedule Request
```
curl "http://localhost:8675/schedule"
```
__Response__
```
{
  "0": [],
  "1496284800": [
    {
      "CronPattern": "*/5 * * * *",
      "ImageName": "alpine",
      "RunCommand": [
        "echo",
        "foo",
        "bar"
      ],
      "NextRun": 1496284800,
      "State": 1,
      "Hash": "19861539a1c0dcf0c72fa80e37a86904ccee2e60"
    },
    {
      "CronPattern": "*/10 * * * *",
      "ImageName": "alpine",
      "RunCommand": [
        "echo",
        "foo",
        "baz"
      ],
      "NextRun": 1496284800,
      "State": 1,
      "Hash": "0c76f5dfd7c1b440715b23016e3ca277c9540fd3"
    }
  ]
}
```

### GET Job Request
```
curl "http://localhost:8675/job"
```

__Response__
```
{
  "CronPattern": "*/1 * * * *",
  "ImageName": "alpine",
  "RunCommand": [
    "echo",
    "hello",
    "world"
  ],
  "NextRun": 1496281500,
  "State": 1,
  "Hash": "7c251e80adab3e1606c65b781be68ba9c05e0b18"
}
```

### POST Request
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

### DELETE Job
```
curl -X "DELETE" "http://localhost:8675/job/7c251e80adab3e1606c65b781be68ba9c05e0b18"
```