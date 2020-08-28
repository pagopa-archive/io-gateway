IOGW_VER?=$(shell git rev-parse --abbrev-ref HEAD)
DOCKER_USER?=pagopa
IMG=$(DOCKER_USER)/iogw-theia:$(IOGW_VER)

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
