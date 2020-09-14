IOSDK_VER?=$(shell git rev-parse --abbrev-ref HEAD)
IMG=pagopa/iosdk-scheduler:$(IOSDK_VER)

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
	--rm --name iosdk-scheduler --hostname iosdkscheduler \
	-e IO_SDK_SCHEDULER_CONFIG='/scheduler/config/io-sdk-scheduler-container-config.json' \
    -v ${HOME}:/scheduler/config \
	$(IMG)

test:
	npm install
	npm run test
