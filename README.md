# go_sample_api_server 

# Environment
- go v1.12

## Installation & Run

### docker-compose
```
$ docker-compose up
$ docker exec -it go_sample_api_server_app_1 go run cmd/migrate/main.go #migrate
$ docker exec -it go_sample_api_server_app_1 /bin/ash 
``` 

### local
```
$ go run cmd/api/main.go
```

## cmd
migrate up
```
$ migrate -source file://app/infra/migrate/ -database 'mysql://username:secret@tcp(127.0.0.1:3306)/go_sample_api_server' up
```
migrate down
```
$ migrate -source file://app/infra/migrate/ -database 'mysql://username:secret@tcp(127.0.0.1:3306)/go_sample_api_server' down
```