This is go_sample_api_server

migrate

```shell script
$ migrate -source file://app/infra/migrate/ -database 'mysql://username:secret@tcp(127.0.0.1:3306)/go_sample_api_server' up
```



migrate
```shell script
$ docker exec -it go_sample_api_server_app_1 go run cmd/migrate/main.go
```

docker exec -it go_sample_api_server_app_1 /bin/ash