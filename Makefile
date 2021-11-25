run:
	go run cmd/server/server.go

client:
	go run cmd/client/client.go

.DEFAULT_GOAL := run
