# protoファイルからgoコードの自動生成
.PHONY: gogen
gogen:
	protoc -Iproto --go_out=plugins=grpc:. proto/*.proto

# protoファイルからrubyコードの自動生成
.PHONY: rubygen
rubygen:
	cd ./client && bundle exec grpc_tools_ruby_protoc -I ../proto --ruby_out=app/gen/api/pancake/baker --grpc_out=app/gen/api/pancake/baker ../proto/pancake.proto


# gofmtの実行
.PHONY: gofmt
gofmt:
	gofmt -l -w ./.