package main

import "fmt"

func ExampleWhiskDockerRunOk() {
	//*DryRunFlag = false

	DryRunPush("", "1.2.3.4", "1234566789", "")
	fmt.Println(whiskDockerRun())
	// Output:
	// docker pull pagopa/iogw-openwhisk:test
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} iogw-redis
	// docker run -d -p 3280:3280 --rm --name iogw-openwhisk --hostname openwhisk -e CONTAINER_EXTRA_ENV=__OW_REDIS=1.2.3.4 -e CONFIG_FORCE_whisk_users_guest=a14ccfb1-d2c8-4ddc-91e6-a6fc4617a53c:Ov9VnxiJjcDIxzXarElXjJizi481LWs8Y2WNp8zCkN7WbggyG5g8xv1n0HTYBkMn	 -v //var/run/docker.sock:/var/run/docker.sock pagopa/iogw-openwhisk:test
	// docker exec iogw-openwhisk waitready
}

func ExampleWhiskDockerRunKo() {
	//*DryRunFlag = false
	DryRunPush("cannot pull pagopa/iogw-openwhisk:test")
	fmt.Println(1, whiskDockerRun())
	DryRunPush("", "Error: cannot find ide")
	fmt.Println(2, whiskDockerRun())

	DryRunPush("", "1.2.3.4", "!cannot start")
	fmt.Println(3, whiskDockerRun())
	DryRunPush("", "1.2.3.4", "1234", "!no wait")
	fmt.Println(4, whiskDockerRun())
	// Output:
	// docker pull pagopa/iogw-openwhisk:test
	// 1 cannot pull pagopa/iogw-openwhisk:test
	// docker pull pagopa/iogw-openwhisk:test
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} iogw-redis
	// 2 cannot locate redis
	// docker pull pagopa/iogw-openwhisk:test
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} iogw-redis
	// docker run -d -p 3280:3280 --rm --name iogw-openwhisk --hostname openwhisk -e CONTAINER_EXTRA_ENV=__OW_REDIS=1.2.3.4 -e CONFIG_FORCE_whisk_users_guest=23bc46b1-71f6-4ed5-8c54-816aa4f8c502:123zO3xZCLrMN6v2BKK1dXYFpXlPkccOFqm12CdAsMgRU4VrNZ9lyGVCGuMDGIwP -v //var/run/docker.sock:/var/run/docker.sock pagopa/iogw-openwhisk:test
	// 3 cannot start server: cannot start
	// docker pull pagopa/iogw-openwhisk:test
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} iogw-redis
	// docker run -d -p 3280:3280 --rm --name iogw-openwhisk --hostname openwhisk -e CONTAINER_EXTRA_ENV=__OW_REDIS=1.2.3.4 -e CONFIG_FORCE_whisk_users_guest=23bc46b1-71f6-4ed5-8c54-816aa4f8c502:123zO3xZCLrMN6v2BKK1dXYFpXlPkccOFqm12CdAsMgRU4VrNZ9lyGVCGuMDGIwP -v //var/run/docker.sock:/var/run/docker.sock pagopa/iogw-openwhisk:test
	// docker exec iogw-openwhisk waitready
	// 4 server readyness error: !no wait
}

func ExampleWhiskDockerRm() {
	// *DryRunFlag = false
	fmt.Println(WhiskDestroy())
	// Output:
	// Destroying Whisk...
	// docker exec iogw-openwhisk stop
	//
	// <nil>
}
