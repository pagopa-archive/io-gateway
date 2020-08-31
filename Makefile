VER?=$(shell git tag --points-at HEAD | head -1)
DOCKER_USER?=pagopa

.PHONY: preflight branch build test release release_mac snapshot all clean
branch:
	$(MAKE) build 
	$(MAKE) test

release:
	test -n "$(VER)"
	$(MAKE) IOGW_VER=$(VER) DOCKER_USER=$(DOCKER_USER) build
	$(MAKE) test
	$(MAKE) IOGW_VER=$(VER) DOCKER_USER=$(DOCKER_USER) push
	$(MAKE) IOGW_VER=$(VER) DOCKER_USER=$(DOCKER_USER) -C iogw/setup/linux
	$(MAKE) IOGW_VER=$(VER) DOCKER_USER=$(DOCKER_USER) -C iogw/setup/windows

release_mac:
	test -n "$(VER)"
	$(MAKE) IOGW_VER=$(VER) DOCKER_USER=$(DOCKER_USER) -C iogw/setup/mac

clean:
	-$(MAKE) -C admin clean
	-$(MAKE) -C ide clean
	-$(MAKE) -C iogw clean

build: preflight
	$(MAKE) DOCKER_USER=$(DOCKER_USER) -C admin
	$(MAKE) DOCKER_USER=$(DOCKER_USER) -C ide
	$(MAKE) DOCKER_USER=$(DOCKER_USER) -C iogw

push:
	docker login
	$(MAKE) DOCKER_USER=$(DOCKER_USER) -C admin push
	$(MAKE) DOCKER_USER=$(DOCKER_USER) -C ide push

test:
	# test cli
	make -C iogw test
	# test execution
	bash test.sh
	# test actions
	make -C admin/actions test

snapshot:
	date "+%Y.%m%d.%H%M-snapshot" | tee .snapshot
	cat .snapshot | xargs git tag 
	git push origin --tags
	cat .snapshot | xargs git tag -d 

preflight:
	echo "checking required versions"
	node -v | grep v12
	python3 -V | grep 3.7
	go version | grep go1.15
