# go_sample_api_server 

# About
This is REST API server sample with golang.  
Using gorilla/mux for Routing. Not using ORM.  
Structure is a Layer Architecture like Clean Architecture.


# Environment
- go v1.12
- mysql v5.7

# Sample Server
[In Heroku](https://go-sample-api-server.herokuapp.com/).  
Use Container deploy and Github-Actions(`.github/workflows/production-deploy.yml`)


## Installation & Run
```
# Download this project
go get github.com/taka011002/go_sample_api_server
```

### docker-compose
```
$ docker-compose up
``` 

### local
```
$ go run main.go
```

### Initialize
```
$ go run cmd/migrate/main.go && \
$ go run cmd/import_character_rarities/main.go other/csv/character_raryties.csv && \
$ go run cmd/import_characters/main.go other/csv/characters.csv
```
## API
Please read `doc/api-doc.yaml`

## cmd

### docker
shell
```
$ docker exec -it go_sample_api_server_app_1 /bin/ash 
```

migrate
```
$ docker exec -it go_sample_api_server_app_1 go run cmd/migrate/main.go
```

...

### local
migrate
```
$ go run cmd/migrate/main.go
```

import character_rarities
```
$ go run cmd/import_character_rarities/main.go other/csv/character_raryties.csv 
```

import characters
```
$ go run cmd/import_characters/main.go other/csv/characters.csv 
```

migrate up
```
$ migrate -source file://app/infra/migrations/ -database 'mysql://username:secret@tcp(127.0.0.1:3306)/go_sample_api_server' up
```
migrate down
```
$ migrate -source file://app/infra/migrations/ -database 'mysql://username:secret@tcp(127.0.0.1:3306)/go_sample_api_server' down
```

## database tables
Please read `other/uml/table.puml`


