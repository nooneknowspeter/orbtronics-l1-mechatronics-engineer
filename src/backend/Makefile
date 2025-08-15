.PHONY: all clean test 

run:
	go run .
watch:
	air
compile:
	mkdir -p build
	go build -o build
api:
	openapi-tui -i api-client/api.yaml