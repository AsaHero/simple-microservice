proto_gen:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto

clear:
	rm pb/*.go

run_server:
	go run cmd/server/main.go
run_clinet:
	go run cmd/client/main.go

migrate-up:
	migrate -source file://migrations -database postgres://asad:Tashkent2001@localhost:5432/laptop_service up

migrate-down:
	migrate -source file://migrations -database postgres://asad:Tashkent2001@localhost:5432/laptop_service down

.PHONY: proto_gen clear run_server run_client migrate-up migrate-down
