run:
	go run ./cmd/main.go

swagger:
	swag init -g cmd/main.go -o ./api

mock:
	mockery --dir=internal/domain/dashboard --name=Repository --filename=dashboard_repository_interface.go --output=internal/domain/dashboard/mocks/repomocks --outpkg=dashboard_repomocks