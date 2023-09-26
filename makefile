build:
	go build -o bin/server main.go
run-api:
	go run tests/multi_requester/api.go
server:
	./bin/server
	
run: build server

# make run
# make run-api
