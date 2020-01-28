This is go_sample_api_server

migrate

```bash
$ migrate -source file://app/infra/migrate/ -database 'mysql://username:secret@tcp(127.0.0.1:3306)/go_sample_api_server' up
```