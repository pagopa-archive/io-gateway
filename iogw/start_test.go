package main

import (
	"fmt"
	"io/ioutil"
)

func ExampleStart() {
	//*DryRunFlag = true
	fmt.Println("=== Init ===")
	_, err := Start()
	fmt.Println(err)
	DryRunPush("/tmp/iogw-test/javascript", "javascript", "123456")
	dir, err := Init("", "", ioutil.Discard)
	fmt.Println(dir, err)
	fmt.Println(Configure(dir))
	fmt.Print(run("ls -a /tmp/iogw-test/.io*"))
	fmt.Println("=== Start ===")
	DryRunPush("\nTotal Memory: 11GiB\n", MinDockerVersion, "", "123", "", "1.2.3.4", "", "", "", "172.17.0.2")
	_, err = Start()
	fmt.Println(err)
	// Output:
	// === Init ===
	// You need to run 'iosdk init ', first.
	// stat /tmp/iogw-test/.iosdk.v3: no such file or directory
	// Select one of the available templates for importers, or provide your own.
	// The javascript template is for Excel import.
	// The php template is for SQL import.
	// The python template is for GraphQL import.
	// The github template requires a github repo (user/path).
	// Preparing work directory /tmp/iogw-test/javascript for https://github.com/pagopa/io-sdk-javascript
	// Done.
	// /tmp/iogw-test/javascript <nil>
	// Wrote /tmp/iogw-test/.iosdk.v3
	// <nil>
	// /tmp/iogw-test/.iosdk.v3
	// === Start ===
	// docker info
	// docker version --format {{.Server.Version}}
	// WARNING: using default OpenWhisk key
	// Deploying Redis...
	// docker pull library/redis:5
	// docker run -d -p 6379:6379 --rm --name iogw-redis --hostname redis library/redis:5
	//
	// Deploying Whisk...
	// docker pull pagopa/iogw-openwhisk:test
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} iogw-redis
	// docker run -d -p 3280:3280 --rm --name iogw-openwhisk --hostname openwhisk -e CONTAINER_EXTRA_ENV=__OW_REDIS=1.2.3.4 -e CONFIG_FORCE_whisk_users_guest=23bc46b1-71f6-4ed5-8c54-816aa4f8c502:123zO3xZCLrMN6v2BKK1dXYFpXlPkccOFqm12CdAsMgRU4VrNZ9lyGVCGuMDGIwP -v //var/run/docker.sock:/var/run/docker.sock pagopa/iogw-openwhisk:test
	// docker exec iogw-openwhisk waitready
	// Deploying io-sdk scheduler engine
	// docker pull pagopa/iogw-scheduler:master
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} iogw-openwhisk
	// docker run -ti -d -p 3100:3100 --rm --name iogw-scheduler --hostname scheduler -e IO_GW_SCHEDULER_CONFIG=/scheduler/config/io-sdk-scheduler-container-config.json -v /tmp/iosdk-test/javascript:/scheduler/config --add-host=openwhisk:172.17.0.2 pagopa/iogw-scheduler:master
	// Deploying IDE...
	// docker pull pagopa/iogw-theia:test
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} iogw-openwhisk
	// docker run -d -p 3000:3000 --rm --name iogw-theia -e HOME=/home/project -v /tmp/iogw-test/javascript:/home/project --add-host=openwhisk:172.17.0.2 pagopa/iogw-theia:test
	// <nil>
}

func ExampleStop() {
	*DryRunFlag = true
	DryRunPush()
	fmt.Println(Stop())
	// Output:
	// Destroying IDE...
	// docker kill iogw-theia
	//
	// Destroying io-sdk scheduler engine
	// docker stop iosdk-scheduler
	//
	// Destroying Whisk...
	// docker exec iogw-openwhisk stop
	//
	// Destroying Redis...
	// docker stop iogw-redis
	//
	// <nil>
}
