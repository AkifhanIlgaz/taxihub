gen-driver-service:
	@protoc \
		--proto_path=proto "proto/driver/driver.proto" \
		--go_out=common/proto/ --go_opt=paths=source_relative \
  	--go-grpc_out=common/proto/ --go-grpc_opt=paths=source_relative


gen-swagger:
	swag init -g api-gateway/cmd/main.go -o api-gateway/docs

up:
	docker-compose up -d

down:
	docker-compose down
