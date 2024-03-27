gen_grpc_protoc:
	protoc \
	--go_out=grpc \
	--go_opt=paths=source_relative \
	--go-grpc_out=grpc \
	--go-grpc_opt=paths=source_relative \
	proto/*.proto
export_path:
	export PATH=$PATH:$(go env GOPATH)/bin