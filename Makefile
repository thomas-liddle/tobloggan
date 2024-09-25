#!/usr/bin/make -f

test:
	go fmt ./...
	go mod tidy
	go test -cover -timeout=1s -race -count=10 ./...

clean:
	rm -rf ./generated

generate: clean
	go run github.com/mdwhatcott/tobloggan/main/tobloggan \
	-source /Users/mike/src/github.com/mdwhatcott/tobloggan/content \
	-target /Users/mike/src/github.com/mdwhatcott/tobloggan/generated \
	-base-url "/tobloggan"