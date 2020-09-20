IOGW_VER?=$(shell git rev-parse --abbrev-ref HEAD)
WSK_URL=https://github.com/apache/openwhisk-cli/releases/download/1.0.0/OpenWhisk_CLI-1.0.0-mac-amd64.zip

BIN=application/bin
WSK=$(BIN)/wsk
SDK=$(BIN)/iogw

.PHONY: pkg
pkg: $(SDK) $(WSK)
	bash build-macos-x64.sh iogw $(IOGW_VER)
	mv target/pkg/iogw-macos-installer-x64-$(IOGW_VER).pkg iogw-$(IOGW_VER).pkg 

$(SDK):
	mkdir -p $(BIN)
	GOOS=darwin GOARCH=amd64 go build \
	-o $(SDK) -ldflags "-X main.Version=$(IOGW_VER) -X main.DockerHubUser=$(DOCKER_USER)" ../.. 

$(WSK):
	mkdir -p $(BIN)
	curl -sL $(WSK_URL) >wsk.zip
	unzip wsk.zip wsk -d $(BIN)

clean:
	-rm -Rvf target application
	-rm *.pkg *.bin

