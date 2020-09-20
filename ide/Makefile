DOCKER_USER?=pagopa
TAG=$(shell git tag --points-at HEAD)
ifeq ($(TAG),)
VER=$(shell git rev-parse --abbrev-ref HEAD)
else
VER=$(TAG)
endif
ifeq ($(VER),)
VER=noversion
endif
IMG=$(DOCKER_USER)/iogw-theia:$(VER)

.PHONY: build
build:
	docker build -t $(IMG) .

.PHONY: push
push: 
	docker push $(IMG)

.PHONY: clean
clean:
	-docker rmi -f $(IMG)

.PHONY: start
start:
	docker run -ti \
	--rm --name iogw-theia \
	--add-host openwhisk:$(shell docker inspect iogw-openwhisk | jq -r '.[0].NetworkSettings.IPAddress') \
	-p 3000:3000 \
	-v /var/run/docker.sock:/var/run/docker.sock \
	$(IMG)

.PHONY: stop
stop:
	-docker kill iogw-theia
