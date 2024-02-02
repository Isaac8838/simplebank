migrateup:
	migrate -path db/migration -database "postgresql://root:password@host.docker.internal:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:password@host.docker.internal:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: migrateup migretedown sqlc test