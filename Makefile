generate:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/ping.proto

run_server:
	@echo "---- Running Server ----"
	@go run server/*

run_client:
	@echo "---- Running Client ----"
	@go run client/*


# define GenerateProtoFiles
# 	@protoc --go_out=$(1) $(GO_OPT) --go-grpc_out=$(1) $(GRPC_OPT) $(PROTO_FLAG)
# 	@echo '---> Generating $(2) Protofiles'
# endef
