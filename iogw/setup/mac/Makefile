IOGW_VER?=0.0-unknown
WSK_URL=https://github.com/apache/openwhisk-cli/releases/download/1.0.0/OpenWhisk_CLI-1.0.0-mac-amd64.zip

BIN=application/bin
WSK=$(BIN)/wsk
SDK=$(BIN)/iogw

.PHONY: pkg
pkg: $(SDK) $(WSK)
	bash build-macos-x64.sh iogw $(IOGW_VER)

$(SDK):
	mkdir -p $(BIN)
	GOOS=darwin GOARCH=amd64 go build \
	-o $(SDK) -ldflags "-X main.Version=$(IOGW_VER) -X main.DockerHubUser=$(DOCKER_USER)" ../.. 

$(WSK):
	mkdir -p $(BIN)
	curl -sL $(WSK_URL) >wsk.zip
	unzip wsk.zip wsk -d $(BIN)

