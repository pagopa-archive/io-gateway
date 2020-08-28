package main

import (
	"fmt"
)

func ExampleConfigLoad() {
	run("rm -Rvf /tmp/iogw-test ; mkdir /tmp/iogw-test")
	fmt.Println(ConfigLoad())
	DryRunPush("123456")
	Configure("/tmp/iogw-test/javascript")
	fmt.Println(ConfigLoad())
	fmt.Print(run("ls -a /tmp/iogw-test/.io*"))
	fmt.Println(Config.IoAPIKey)
	fmt.Println(len(Config.WhiskAPIKey))
	// Output:
	// stat /tmp/iogw-test/.iosdk.v3: no such file or directory
	// Wrote /tmp/iogw-test/.iosdk.v3
	// <nil>
	// /tmp/iogw-test/.iosdk.v3
	// 123456
	// 101
}
