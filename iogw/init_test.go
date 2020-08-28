package main

import (
	"fmt"
	"io/ioutil"
)

func ExampleInit2() {
	run("rm -Rf /tmp/iogw-test")
	DryRunPush("/tmp/iogw-test/javascript", "javascript")
	dir, err := Init("", "", ioutil.Discard)
	fmt.Println(1, dir, err)
	DryRunPush("")
	dir, err = Init("", "", ioutil.Discard)
	fmt.Println(2, err)
	DryRunPush("/tmp/iogw-test/javascript", "")
	dir, err = Init("", "", ioutil.Discard)
	fmt.Println(3, err, dir)
	// Output:
	// Select one of the available templates for importers, or provide your own.
	// The javascript template is for Excel import.
	// The php template is for SQL import.
	// The python template is for GraphQL import.
	// The github template requires a github repo (user/path).
	// Preparing work directory /tmp/iogw-test/javascript for https://github.com/pagopa/io-sdk-javascript
	// Done.
	// 1 /tmp/iogw-test/javascript <nil>
	// 2 Directory not specified
	// 3 <nil> /tmp/iogw-test/javascript
}

func ExampleInit() {
	//*DryRunFlag = false
	run("rm -Rvf /tmp/iogw-test ; mkdir /tmp/iogw-test")
	fmt.Println(0, last(Init("/var/tmp/demo", "", ioutil.Discard)))
	DryRunPush("javascript")
	fmt.Println(1, last(Init("/tmp/iogw-test/demo", "", ioutil.Discard)))
	fmt.Print(run("ls /tmp/iogw-test"))
	fmt.Println(2, last(Init("/tmp/iogw-test/demo", "", ioutil.Discard)))
	run("rm -Rf /tmp/iogw-test")
	fmt.Println(3, last(Init("/tmp/iogw-test/clone", "pagopa/io-sdk-javascript", ioutil.Discard)))
	fmt.Print(run("ls /tmp/iogw-test"))
	fmt.Println(4, last(Init("/tmp/iogw-test/donotexist", "pagopa/do-not-exist", ioutil.Discard)))
	// Output:
	// 0 work directory /var/tmp/demo should be below your home directory /tmp/iogw-test;
	// this is required to be accessible by Docker
	// Select one of the available templates for importers, or provide your own.
	// The javascript template is for Excel import.
	// The php template is for SQL import.
	// The python template is for GraphQL import.
	// The github template requires a github repo (user/path).
	// Preparing work directory /tmp/iogw-test/demo for https://github.com/pagopa/io-sdk-javascript
	// Done.
	// 1 <nil>
	// demo
	// 2 <nil>
	// Preparing work directory /tmp/iogw-test/clone for https://github.com/pagopa/io-sdk-javascript
	// Done.
	// 3 <nil>
	// clone
	// Preparing work directory /tmp/iogw-test/donotexist for https://github.com/pagopa/do-not-exist
	// 4 authentication required
}
