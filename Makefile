startdb:
	docker compose up -d

stopdb:
	docker compose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

migrateup:
	migrate -path schema/migration -database "postgresql://main:main@localhost:4000/main?sslmode=disable" -verbose up

migratedown:
	migrate -path schema/migration -database "postgresql://main:main@localhost:4000/main?sslmode=disable" -verbose down

.PHONY: startdb stopdb sqlc migrateup migratedown test