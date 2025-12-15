# Run locally
```shell
go run main.go

# expect 401
curl -X POST http://localhost:3000

# expect 200
curl -X POST http://localhost:3000 \
    -H "Authorization: Bearer token"
```
# Building code into image and run as container
```shell
docker build -t auth:latest .

docker run auth:latest -p 3000:3000

# expect 401
curl -X POST http://localhost:3000

# expect 200
curl -X POST http://localhost:3000 \
    -H "Authorization: Bearer token"
```