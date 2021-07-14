# protoファイルからコードの自動生成
.PHONY: protogen
protogen:
	protoc -Iproto --go_out=plugins=grpc:. proto/*.proto