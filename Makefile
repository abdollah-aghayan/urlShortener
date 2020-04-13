# Make is verbose in Linux. Make it silent.
# Also we can silent a command with putting @ at the bigining of a command
MAKEFLAGS += --silent

build:
	echo "  >  Building binary..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-s' -o ./bin/urlShortner
	echo "Done!"

build-containers:
	echo "  >  Building containers..."
	docker-compose up
	echo "Done"

run: build build-containers

test: 
	go test -v ./...

create-test-coverage: 
	go test -coverprofile cover.out ./...

show-html-coverage: create-test-coverage
	@go tool cover -html=cover.out