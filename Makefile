.PHONY: build

build: bin/hknpm bin/hknode gonpm

bin/hknpm: tmp/node
	cp tmp/node-v0.10.32-darwin-x64/bin/npm bin/hknpm

tmp/node:
	mkdir -p tmp
	cd tmp; \
	curl http://nodejs.org/dist/v0.10.32/node-v0.10.32-darwin-x64.tar.gz -o node.tar.gz; \
	tar xzf node.tar.gz

gonpm:
	go build
