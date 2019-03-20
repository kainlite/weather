.PHONY: build clean deploy

build:
build:
	go get .

	env GOOS=linux go build -ldflags="-s -w" -o main endpoints/locations/create/main.go
	mkdir -p bin/locations
	zip bin/locations/create.zip main
	mv main bin/locations/create

	env GOOS=linux go build -ldflags="-s -w" -o main endpoints/locations/delete/main.go
	mkdir -p bin/locations
	zip bin/locations/create.zip main
	mv main bin/locations/create

	env GOOS=linux go build -ldflags="-s -w" -o main endpoints/locations/search/main.go
	mkdir -p bin/locations
	zip bin/locations/delete.zip main
	mv main bin/locations/delete

clean:
	rm -rf ./bin ./vendor

deploy: clean build
	sls deploy --verbose
