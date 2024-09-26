#!/usr/bin/make -f

test:
	go fmt ./...
	go mod tidy
	go test -cover -timeout=1s -race -count=10 ./...

clean:
	rm -rf ./docs ./generated

dev:
	go run tobloggan/main/tobloggan \
    	-source /Users/mike/src/tobloggan/content \
    	-target /Users/mike/src/tobloggan/generated \
    	-base-url "http://localhost:8000" && \
    open "http://localhost:8000/" && \
    python3 -m http.server --directory generated

publish: clean
	go run tobloggan/main/tobloggan \
		-source /Users/mike/src/tobloggan/content \
		-target /Users/mike/src/tobloggan/docs \
		-base-url "/tobloggan" && \
	git add ./docs && \
	git commit -m "auto-publish" && \
	git push origin master

.PHONY: test clean dev publish
