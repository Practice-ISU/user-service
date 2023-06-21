run: go_gen_registration go_gen_ping
	echo Generated

go_gen_registration:
	protoc --go_out=../../pkg/grpc/discovery/registration --go_opt=paths=source_relative discovery-registration.proto
	protoc --go-grpc_out=../../pkg/grpc/discovery/registration --go-grpc_opt=paths=source_relative discovery-registration.proto

go_gen_ping:
	protoc --go_out=../../pkg/grpc/discovery/ping --go_opt=paths=source_relative discovery-ping.proto
	protoc --go-grpc_out=../../pkg/grpc/discovery/ping --go-grpc_opt=paths=source_relative discovery-ping.proto