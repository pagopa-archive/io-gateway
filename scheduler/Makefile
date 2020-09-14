IOGW_VER?=$(shell git rev-parse --abbrev-ref HEAD)
IMG=pagopa/iogw-scheduler:$(IOGW_VER)

.PHONY: test clean

build:
	npm install
	npm run build
	docker build -t $(IMG) .

push:
	docker push $(IMG)

clean:
	docker rmi -f $(IMG)

start:
	docker run -ti -d \
	-p 3100:3100 \
	--rm --name iogw-scheduler --hostname iogwscheduler \
	-e IO_GW_SCHEDULER_CONFIG='/scheduler/config/io-gw-scheduler-container-config.json' \
    -v ${HOME}:/scheduler/config \
	$(IMG)

test:
	npm install
	npm run test
