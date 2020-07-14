build:
	docker-compose build --no-cache

up:
	docker-compose up -d

down:
	docker-compose down

restart:
	make down
	make up

destroy:
	docker-compose down --rmi all --volumes

destroy-volumes:
	docker-compose down --volumes

ps:
	docker-compose ps

bash:
	docker-compose exec app /bin/ash

migrate:
	docker-compose exec app go run cmd/migrate/main.go

init:
	docker-compose exec app go run cmd/migrate/main.go && \
	docker-compose exec app go run cmd/import_character_rarities/main.go other/csv/character_raryties.csv && \
	docker-compose exec app go run cmd/import_characters/main.go other/csv/characters.csv