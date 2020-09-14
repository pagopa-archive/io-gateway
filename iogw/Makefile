IOGW_VER?=$(shell git rev-parse --abbrev-ref HEAD)
DOCKER_USER?=pagopa

iogw: $(shell ls *.go)
	go build -ldflags "-X main.Version=$(IOGW_VER) -X main.DockerHubUser=$(DOCKER_USER)" -o iogw

.PHONY: test clean
test:
	go test -v | tee difftest.out
	python3 difftest.py

clean:
	-rm iogw
	$(MAKE) -C setup/windows clean
	$(MAKE) -C setup/linux clean
	$(MAKE) -C setup/mac clean
