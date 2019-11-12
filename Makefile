IMAGE_NAME = socket-ad
VERSION = 1.0.0

CONTAINER_NAME = socket-ad
CONTAINER_INSTANCE = default
PORTS = -p 9090:9090

.PHONY: build run

build: Dockerfile
	docker build -t $(IMAGE_NAME):$(VERSION) -f Dockerfile .

run:
	docker run --rm --name $(CONTAINER_NAME)-$(CONTAINER_INSTANCE) $(PORTS) $(IMAGE_NAME):$(VERSION)

default: build
