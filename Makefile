gen-driver-service:
	@protoc \
		--proto_path=proto "proto/driver/driver.proto" \
		--go_out=common/proto/ --go_opt=paths=source_relative \
  	--go-grpc_out=common/proto/ --go-grpc_opt=paths=source_relative
