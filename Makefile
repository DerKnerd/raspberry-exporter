GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=raspberry_exporter
all: aarch64 armhf
aarch64:
	$(GOBUILD) -o out/armhf/$(BINARY_NAME)
armhf:
	$(GOBUILD) -o out/aarch64/$(BINARY_NAME)