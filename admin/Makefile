DOCKER_USER?=pagopa
TAG=$(shell git tag --points-at HEAD)
ifeq ($(TAG),)
VER=$(shell git branch --show-current)
else
VER=$(TAG)
endif
IMG=$(DOCKER_USER)/iogw-openwhisk:$(VER)

.PHONE: build
build: public/build/bundle.js
	$(MAKE) -C actions
	docker build -t $(IMG) .

.PHONY: push
push:
	docker push $(IMG)

.PHONY: clean
clean:
	-$(MAKE) -C actions clean
	-docker rmi -f $(IMG)

.PHONY: deploy
deploy:
	$(MAKE) -C actions deploy

.PHONY: devel
devel:
	npm install
	npm run dev

.PHONY: start
start: ../iogw/iogw
	../iogw/iogw debug wskprops
	../iogw/iogw start --skip-ide --skip-pull-images

.PHONY: stop
stop: ../iogw/iogw
	-../iogw/iogw stop

../iogw/iogw: build
	$(MAKE) -C ../iogw

public/build/bundle.js: $(wildcard src/*.svelte)
	node -v | grep v12
	npm install
	npm run build


.PHONY: undeploy
undeploy:
	$(MAKE) -C actions undeploy
