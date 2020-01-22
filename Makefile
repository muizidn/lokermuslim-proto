generated_dir := generated

all: swift go

swift:
	rm -rf ${generated_dir}/swift
	mkdir -p ${generated_dir}/swift
	# Swift requires absolute path
	protoc *.proto \
		--proto_path=. \
		--plugin=$(which protoc-gen-swift) \
		--swift_opt=Visibility=Public \
		--swift_out=${generated_dir}/swift
	protoc *.proto \
		--proto_path=. \
		--plugin=$(which protoc-gen-grpc-swift) \
		--grpc-swift_opt=Visibility=Public \
		--grpc-swift_out=${generated_dir}/swift
go:
	rm -rf ${generated_dir}/go
	mkdir -p ${generated_dir}/go
	protoc --go_out=plugins=grpc:${generated_dir}/go *.proto
	
.PHONY: all
