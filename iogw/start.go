package main

import "fmt"

// Start openwhisk-ide
func Start() (string, error) {
	if err := ConfigLoad(); err != nil {
		fmt.Println("You need to run 'iosdk init ', first.")
		return "", err
	}
	info, err := Preflight(Config.AppDir)
	if err != nil {
		return "", err
	}
	if *useDefaultAPIKey {
		Config.WhiskAPIKey = DefaultWhiskAPIKey
		fmt.Println("WARNING: using default OpenWhisk key")
	}
	err = RedisDeploy()
	if err != nil {
		return "", err
	}
	err = WhiskDeploy()
	if err != nil {
		return "", err
	}

	err = SchedulerDeploy()
	if err != nil {
		return "", err
	}

	if !*skipIde {
		err = IdeDeploy(Config.AppDir, info)
		if err != nil {
			return "", err
		}
	}
	return info, nil
}

// Stop openwhisk-ide
func Stop() error {
	IdeDestroy()
	SchedulerDestroy()
	WhiskDestroy()
	RedisDestroy()
	return nil
}
