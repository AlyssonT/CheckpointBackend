run-api:
	swag init -g cmd/api/main.go
	swag fmt
	go run cmd/api/main.go

run-dbsync:
	go run cmd/db_sync/main.go