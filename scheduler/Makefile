IMG=ftt/iosdk-scheduler:latest

.PHONE: build
build:
	npm install
	npm run build
	docker build -t $(IMG) .

start:
	docker run -ti -d \
	-p 3000:3000 \
	--rm --name iosdk-scheduler --hostname iosdkscheduler \
	-e IO_SDK_SCHEDULER_CONFIG='/scheduler/config/io-sdk-scheduler-container-config.json' \
    -v ${HOME}:/scheduler/config \
	$(IMG)
