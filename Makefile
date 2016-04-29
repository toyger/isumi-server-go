TARGETS=$(shell GO15VENDOREXPERIMENT=1 go list ./... | awk '$$0 !~ /$(IGNORE)/{print $0}')
ARCH=$(shell uname | tr '[:upper:]' '[:lower:]')
VERSION=$(shell git rev-parse --verify HEAD)

install-go:
	sh install-go.sh 1.6.2

install-glide:
	sh install-glide.sh 0.8.3

deps: install-glide
	GO15VENDOREXPERIMENT=1 glide install

deps-update:
	glide update

vet:
	GO15VENDOREXPERIMENT=1 go vet $(TARGETS)

test: vet
	GO15VENDOREXPERIMENT=1 go test -cover $(TARGETS)

build:
	GO15VENDOREXPERIMENT=1 GOOS=$(ARCH) GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)" -o bin/isumi main.go
