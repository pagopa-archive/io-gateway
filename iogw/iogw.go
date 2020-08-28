package main

// Version is the current version - it will be set when built
var Version = "master"

//DockerHubUser changed by ldflags
var	DockerHubUser = "pagopa"

//main
func main() {
	Main(DockerHubUser+"/"+Version)
}
