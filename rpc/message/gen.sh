clang-format -i *.proto
protoc --go_out=plugins=grpc:. *.proto