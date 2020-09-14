package main

import (
	"fmt"
)

// SchedulerDeploy deploys scheduler docker with config parameters
func SchedulerDeploy() error {
	fmt.Println("Deploying io-sdk scheduler engine")
	fmt.Println(schedulerDockerRun(DefaultSchedulerConfigFile,Config.AppDir,DefaultSchedulerConfigPath, SchedulerImage))
	return nil
}

// DestroyDestroy stop the scheduler docker
func SchedulerDestroy() error {
	fmt.Println("Destroying io-sdk scheduler engine")
	fmt.Println(Sys("docker stop iosdk-scheduler"))
	return nil
}

// return empty string if ok, otherwise the error
func schedulerDockerRun(configFile string, appDir string, configPath string, image string) string {
	if err := dockerPull(image); err != nil {
		return err.Error()
	}

	openWhiskIP := dockerIP("iosdk-openwhisk")
        if openWhiskIP == nil {
            		return "cannot find openwhisk";
        }

	cmd := fmt.Sprintf(`docker run -ti -d -p 3100:3100
    --rm --name iosdk-scheduler --hostname scheduler
    -e IO_SDK_SCHEDULER_CONFIG=%s
    -v %s:%s
    --add-host=openwhisk:%s
    %s`,configFile, appDir, configPath, *openWhiskIP, image)

	_, err := SysErr(cmd)
	if err != nil {
		return "cannot start iosdk-scheduler: " + err.Error()
	}
	return ""
}
