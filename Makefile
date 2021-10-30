GO ?= $(shell which go)
GO_FLAGS = -tags netgo -ldflags '-w -s'
GO_SRC = $(shell find -name '*.go')
EMBED_SRC = $(shell find cmd/loona/static)

all: loona

loona: $(GO_SRC)
	CGO_ENABLED=0 $(GO) build -v $(GO_FLAGS) -o $@

.PHONY: clean
deploy: loona
	./deploy.sh

.PHONY: clean
clean:
	rm -rf loona