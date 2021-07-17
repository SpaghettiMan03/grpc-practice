# protoファイルからgoコードの自動生成
.PHONY: gogen
gogen:
	protoc -Iproto --go_out=plugins=grpc:. proto/*.proto

# protoファイルからrubyコードの自動生成
.PHONY: rubygen
rubygen:
	cd ./client && bundle exec grpc_tools_ruby_protoc -I ../proto --ruby_out=../gen/api/pancake/maker --grpc_out=../gen/api/pancake/maker ../proto/pancake.proto


# gofmtの実行
.PHONY: gofmt
gofmt:
	gofmt -l -w ./.