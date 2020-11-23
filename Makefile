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

.PHONY: preflight branch build test release release_mac snapshot all clean
branch:
	$(MAKE) build 
	$(MAKE) test

release_images:
	test -n "$(VER)"
	$(MAKE) IOGW_VER=$(VER) DOCKER_USER=$(DOCKER_USER) build
	$(MAKE) test
	$(MAKE) IOGW_VER=$(VER) DOCKER_USER=$(DOCKER_USER) push

release_mac:
	test -n "$(VER)"
	$(MAKE) IOGW_VER=$(VER) DOCKER_USER=$(DOCKER_USER) -C iogw/setup/mac

release_win:
	test -n "$(VER)"
	$(MAKE) IOGW_VER=$(VER) DOCKER_USER=$(DOCKER_USER) -C iogw/setup/windows

release_lin:
	test -n "$(VER)"
	$(MAKE) IOGW_VER=$(VER) DOCKER_USER=$(DOCKER_USER) -C iogw/setup/linux

clean:
	-$(MAKE) -C admin clean
	-$(MAKE) -C ide clean
	-$(MAKE) -C iogw clean
	-$(MAKE) -C scheduler clean

build: preflight
	$(MAKE) DOCKER_USER=$(DOCKER_USER) IOGW_VER=$(VER) -C admin
	$(MAKE) DOCKER_USER=$(DOCKER_USER) IOGW_VER=$(VER) -C ide
	$(MAKE) DOCKER_USER=$(DOCKER_USER) IOGW_VER=$(VER) -C iogw
	$(MAKE) DOCKER_USER=$(DOCKER_USER) IOGW_VER=$(VER) -C scheduler


push:
	docker login
	$(MAKE) DOCKER_USER=$(DOCKER_USER) IOGW_VER=$(VER) -C admin push
	$(MAKE) DOCKER_USER=$(DOCKER_USER) IOGW_VER=$(VER) -C ide push
	$(MAKE) DOCKER_USER=$(DOCKER_USER) IOGW_VER=$(VER) -C scheduler push

test:
	# test cli
	make -C iogw test
	# test execution
	bash test.sh
	# test actions
	make -C admin/actions test
	# test scheduler
	make -C scheduler test

snapshot:
	date "+%Y.%m%d.%H%M-snapshot" | tee .snapshot
	cat .snapshot | xargs git tag 
	git push origin
	git push origin --tags
	cat .snapshot | xargs git tag -d 

preflight:
	echo "checking required versions"
	node -v | grep v12
	python3 -V | grep 3.7
	go version | grep go1.15
