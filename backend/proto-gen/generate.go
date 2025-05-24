package main

//go:generate protoc --proto_path=../protos/inventory --go_out=./pb/inventory --go-grpc_out=./pb/inventory ../protos/inventory/inventory.proto
//go:generate protoc --proto_path=../protos/order --go_out=./pb/order --go-grpc_out=./pb/order ../protos/order/order.proto
//go:generate protoc --proto_path=../protos/statistics --go_out=./pb/statistics --go-grpc_out=./pb/statistics ../protos/statistics/statistics.proto
//go:generate protoc --proto_path=../protos/order --go_out=./pb/order --go-grpc_out=./pb/order ../protos/order/statistics.proto

func main() {}
