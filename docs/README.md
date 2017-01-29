## Wstats HTTP Docs (Version 1)

* Wstats http api documentation (v1)

## Mandatory Headers

|Header|Value|
|:--|:--|
|Content-Type|application/json|

## Stats Index
- Request
```bash
curl -X GET \
-H "Content-Type: application/json" \
localhost:8080/v1/stats
```

- Response
```json
{
  "count": 6,
  "total": 30,
  "top_words": ["lorem","foo","porem","bar"],
  "top_letters": ["o","r","e","m","l"]
}
```

## Healthz
- Request
```bash
curl -X GET \
-H "Content-Type: application/json" \
localhost:8080/v1/healthz
```

- Response
```json
{
  "status": "up"
}
```
