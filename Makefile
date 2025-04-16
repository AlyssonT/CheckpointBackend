run-api:
	swag init -g cmd/api/main.go
	swag fmt
	go run cmd/api/main.go

run-dbsync-igdb:
	go run ./cmd/db_sync_igdb

run-dbsync-steam:
	go run ./cmd/db_sync_steam