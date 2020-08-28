IOSDK_VER?=$(shell git rev-parse --abbrev-ref HEAD)

iosdk: $(shell ls *.go)
	go build -ldflags "-X main.Version=$(IOSDK_VER) -X main.DockerHubUser=$(DOCKER_USER)" -o iosdk

.PHONY: test clean
test:
	go test -v | tee difftest.out
	python3 difftest.py

clean:
	-rm iosdk
	-rm setup/windows/iosdk.exe
