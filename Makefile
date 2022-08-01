include .env

postgres_up:
	docker-compose up -d
postgres_down:
	docker-compose down
migrate_up:	
	docker run --rm -v ${PWD}/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database ${DB_DSN} up
migrate_down:
	docker run --rm -v ${PWD}/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database ${DB_DSN} down