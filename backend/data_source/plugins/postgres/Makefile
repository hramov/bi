gen:
	protoc --go_out=./ --go_opt=paths=source_relative \
	--go-grpc_out=./ --go-grpc_opt=paths=source_relative ./proto/postgres.proto

test:
	go build -o bin/postgres ./plugin/main.go
	go run ./main.go