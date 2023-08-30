go_build:
	docker-compose up
go_run:
	go clean -testcache && go test ./tests/
go_clean:
	docker compose down -v
go_swagger:
	swag init -g cmd/server.go