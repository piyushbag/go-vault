PROTO_DIR := protobuf
OUT_DIR := services/common/genproto/orders

gen:
	@mkdir -p $(OUT_DIR)
	@protoc \
		--proto_path=$(PROTO_DIR) $(PROTO_DIR)/orders.proto \
		--go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative
