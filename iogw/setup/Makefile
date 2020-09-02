.PHONY: all linux mac windows
IOGW_VER?=0.1-unknown

all: linux mac windows

linux:
	$(MAKE) IOGW_VER=$(IOGW_VER) -C linux

mac:
	if which pkgbuild ;\
	then $(MAKE) IOGW_VER=$(IOGW_VER) -C mac ;\
	else echo "Not a Mac, skipping" ; fi

windows:
	$(MAKE) IOGW_VER=$(IOGW_VER) -C windows

clean:
	$(MAKE) clean -C windows
	$(MAKE) clean -C linux
	$(MAKE) clean -C mac

builder-fpm-nsis:
	docker build -t pagopa/builder-fpm-nsis .
	docker push pagopa/builder-fpm-nsis
