GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=raspberry_exporter
all: aarch64 armhf
aarch64:
	GOOS=linux GOARCH=arm64 $(GOBUILD) -o out/aarch64/$(BINARY_NAME)
armhf:
	GOOS=linux GOARCH=arm $(GOBUILD) -o out/armhf/$(BINARY_NAME)