IOGW_VER?=$(shell git rev-parse --abbrev-ref HEAD)
WSK=https://github.com/apache/openwhisk-cli/releases/download/1.0.0/OpenWhisk_CLI-1.0.0-windows-amd64.zip
HERE:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

.PHONY: all
all: iogw.exe wsk.exe
	docker run -v $(HERE):/mnt -e VER=$(IOGW_VER) pagopa/builder-fpm-nsis bash /mnt/build.sh
	mv iogw-setup.exe iogw-$(IOGW_VER).exe

iogw.exe:
	GOOS=windows GOARCH=amd64 go build \
	-o iogw.exe -ldflags "-X main.Version=$(IOGW_VER) -X main.DockerHubUser=$(DOCKER_USER)" ../.. 

wsk.exe:
	curl -sL $(WSK) >wsk.zip
	unzip wsk.zip wsk.exe

clean:
	-rm *.zip *.exe


