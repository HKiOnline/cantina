
build:
	go build -o bin/cantina .
	go install .

generate_proto:
	protoc --go_out=pkg/lib --proto_path=proto character.proto ship.proto
