NODE_VERSION=v0.10.32
NODE_BASE=node-$(NODE_VERSION)-darwin-x64

.PHONY: build

build: node
	go build

node:
	mkdir -p tmp
	curl http://nodejs.org/dist/$(NODE_VERSION)/$(NODE_BASE).tar.gz -o tmp/$(NODE_BASE).tar.gz
	cd tmp; tar xzf $(NODE_BASE).tar.gz
	mv tmp/$(NODE_BASE) node
