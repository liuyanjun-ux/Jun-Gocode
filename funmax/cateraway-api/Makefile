.PHONY: build clean deploy

functions := $(shell find handlers -name \*main.go | awk -F'/' '{print $$2}')

build:
	@for function in $(functions) ; do \
		env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/$$function handlers/$$function/main.go ; \
	done

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose
