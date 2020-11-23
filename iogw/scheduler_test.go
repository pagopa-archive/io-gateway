package main

import "fmt"

func ExampleSchedulerRunOk() {
	*DryRunFlag = true
	DryRunPush("", "172.17.0.2")
	fmt.Println(schedulerDockerRun("/scheduler/config/io-gw-scheduler-container-config.json", "/home/users/test", "/scheduler/config", "pagopa/iosdk-scheduler:test"))
	// Output:
	// docker pull pagopa/iosdk-scheduler:test
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} iogw-openwhisk
	// docker run -ti -d -p 3100:3100 --rm --name iosdk-scheduler --hostname scheduler -e IO_GW_SCHEDULER_CONFIG=/scheduler/config/io-gw-scheduler-container-config.json -v /home/users/test:/scheduler/config --add-host=openwhisk:172.17.0.2 pagopa/iosdk-scheduler:test
}

func ExampleSchedulerDockerDestroy() {
	// *DryRunFlag = false
	fmt.Println(SchedulerDestroy())
	// Output:
	// Destroying io-sdk scheduler engine
	// docker stop iosdk-scheduler
	//
	// <nil>
}
