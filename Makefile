SHELL:=/bin/bash
GIT_SHA ?= $(shell git rev-parse HEAD)

.PHONY: docker.build
docker.build:
	docker build -t $(PROJECT_NAME):$(GIT_SHA) .

.PHONY: docker
docker: docker.build
	docker login -u $(DOCKER_USER) -p $(DOCKER_PASSWORD)
	docker tag $(PROJECT_NAME):$(GIT_SHA) $(REPOSITORY_NAME):latest
	docker tag $(PROJECT_NAME):$(GIT_SHA) $(REPOSITORY_NAME):build-$(GIT_SHA)
	docker push $(REPOSITORY_NAME):latest
	docker push $(REPOSITORY_NAME):build-$(GIT_SHA)

.PHONY: default
default: docker