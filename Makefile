.PHONY: build clean deploy

build:
build:
	go get .
	mkdir -p bin/locations

	env GOOS=linux go build -ldflags="-s -w" -o main endpoints/locations/create/main.go
	zip bin/locations/create.zip main
	mv main bin/locations/create

	env GOOS=linux go build -ldflags="-s -w" -o main endpoints/locations/search/main.go
	zip bin/locations/delete.zip main
	mv main bin/locations/delete

	env GOOS=linux go build -ldflags="-s -w" -o main endpoints/locations/delete/main.go
	zip bin/locations/delete.zip main
	mv main bin/locations/delete

clean:
	rm -rf ./bin ./vendor

deploy: clean build
	sls deploy --verbose
