# Run acceptance tests
.PHONY: testacc

TEST?=./...
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
PKG_NAME=cloudtruth
WEBSITE_REPO=github.com/hashicorp/terraform-website

# todo: revert this
# VERSION=$(shell [ ! -z `git tag -l --contains HEAD` ] && git tag -l --contains HEAD || git rev-parse --short HEAD)

VERSION=0.0.1
GOPATH=$(shell go env GOPATH)

default: build

build: fmtcheck
	go install

# Install for local testing on linux, todo: parameterize
install: fmtcheck
	mkdir -p ./dist; \
	go build -o ./dist/terraform-provider-cloudtruth_$(VERSION); \
	INSTALL_DIR="$(HOME)/.terraform.d/plugins/registry.terraform.io/cloudtruth/cloudtruth/$(VERSION)/linux_amd64"; \
	mkdir -p $$INSTALL_DIR; \
	cp ./dist/terraform-provider-cloudtruth_$(VERSION) $$INSTALL_DIR

sweep:
	@echo "WARNING: This will destroy infrastructure. Use only in development accounts."
	go test $(TEST) -v -sweep=$(SWEEP) $(SWEEPARGS)

test: fmtcheck
	go test $(TEST) -timeout=30s -parallel=4

testacc: fmtcheck
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m -race

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

lint:
	golangci-lint run --timeout 1h ./...

errcheck:
	@sh -c "'$(CURDIR)/scripts/errcheck.sh'"

vendor-status:
	@govendor status

test-compile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./$(PKG_NAME)"; \
		exit 1; \
	fi
	go test -c $(TEST) $(TESTARGS)

release: fmtcheck
	for kernel in linux windows darwin; do \
		for dist in $$(go tool dist list | grep $$kernel); do  \
			GOOS=$$kernel; \
			GOARCH=$$(echo $$dist | cut -d/ -f2); \
			GOOS=$$GOOS GOARCH=$$GOARCH go build -o terraform-provider-cloudtruth_$(VERSION); \
			tar -czf terraform-provider-cloudtruth-$$GOOS-$$GOARCH.tar.gz terraform-provider-cloudtruth_$(VERSION) --remove-files; \
		done \
	done
