.PHONY: fmt test deps-up

PKG_PATH = github.com/rez-go/langmap
GOLANG_IMAGE ?= golang:1.14

fmt:
	@echo "Formatting files..."
	@docker run --rm \
		-v $(CURDIR):/go \
		--entrypoint gofmt \
		$(GOLANG_IMAGE) -w -l -s .

deps-up:
	@echo "Updating all dependencies..."
	@docker run --rm \
		-v $(CURDIR):/$(PKG_PATH) \
		--workdir /$(PKG_PATH) \
		$(GOLANG_IMAGE) /bin/sh -c "go get -u all && go mod tidy"
