.PHONY: lint test vendor clean

export GO111MODULE=on

default: lint test

lint:
	golangci-lint run

test:
	go test -v -cover ./...

yaegi_test:
	yaegi test -v .

vendor:
	go mod vendor

clean:
	rm -rf ./vendor

start_headers_reader:
	python3 testconfig/printheaders.py

testcontainer:
	docker build -t traefiktest .
	docker run\
		--rm \
		--name traefiktest \
		--network host \
		-it traefiktest
