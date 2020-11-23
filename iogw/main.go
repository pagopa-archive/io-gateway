package main

import (
	"fmt"
	"os"
	"time"

	"github.com/pkg/browser"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
)

// VerboseFlag is flag for verbose

var (

	// TestModeFlag enable behaviours useful for testing

	// global flags
	verboseFlag = kingpin.Flag("verbose", "Verbose Output").Short('v').Default("false").Bool()

	// hidden global flags
	useDefaultAPIKey  = kingpin.Flag("use-default-api-key", "Use Default Whisk Api Key").Hidden().Default("false").Bool()
	skipDockerVersion = kingpin.Flag("skip-docker-version", "Skip check of docker version").Hidden().Default("false").Bool()
	skipPullImages    = kingpin.Flag("skip-pull-images", "skip pull images").Hidden().Default("false").Bool()
	skipOpenBrowser   = kingpin.Flag("skip-open-browser", "skip pull images").Hidden().Default("false").Bool()
	skipIde           = kingpin.Flag("skip-ide", "skip starting ide").Hidden().Default("false").Bool()

	// hidden debug commands
	debugCmd        = kingpin.Command("debug", "debug")
	wskProps        = debugCmd.Command("wskprops", "Create WskProps file")
	ideDeployCmd    = debugCmd.Command("ide-deploy", "Create IDE deployment")
	ideDestroyCmd   = debugCmd.Command("ide-destroy", "Destroy IDE deployment")
	whiskDeployCmd  = debugCmd.Command("whisk-deploy", "Create Whisk deployment")
	whiskDestroyCmd = debugCmd.Command("whisk-destroy", "Destroy Whisk deployment")
	redisDeployCmd  = debugCmd.Command("redis-deploy", "Create Redis deployment")
	redisDestroyCmd = debugCmd.Command("redis-destroy", "Destroy Redis deployment")
	inputCmd        = debugCmd.Command("input", "Input test")
	inputArgCmd     = inputCmd.Arg("input arg", "input arg").Default("").String()
	inputSelectFlag = inputCmd.Flag("select", "select").Bool()
	// start, stop, init and status
	startCmd = kingpin.Command("start", "Start Development Enviroment")
	// init
	initCmd          = kingpin.Command("init", "Initialise SDK Repository and related informations")
	initDirArg       = initCmd.Arg("directory", "work directory").Default("").String()
	initRepoArg      = initCmd.Arg("repo", "Repository").Default("").String()
	initWhiskKeyFlag = initCmd.Flag("whisk-apikey", "Whisk API Key").Default("").String()
	initIOKeyFlag    = initCmd.Flag("io-apikey", "IO API Key").Default("").String()
	initWskPropsFlag = initCmd.Flag("wskprops", "Write .wskprops").Default("false").Bool()

	// stop
	stopCmd = kingpin.Command("stop", "Stop Development Environment")
	// restart
	restartCmd = kingpin.Command("restart", "Restart Development Environment")
	// status
	statusCmd = kingpin.Command("status", "Check Containers Status")
)

func parseDebug(cmd string) bool {
	switch cmd {
	case wskProps.FullCommand():
		ConfigLoad()
		WskPropsSave()
	case ideDeployCmd.FullCommand():
		FatalIf(ConfigLoad())
		info, _ := Preflight(Config.AppDir)
		IdeDeploy(Config.AppDir, info)
		configureIde(info)
	case ideDestroyCmd.FullCommand():
		IdeDestroy()
	case whiskDeployCmd.FullCommand():
		FatalIf(ConfigLoad())
		WhiskDeploy()
		WhiskUpdatePackageParameters("iosdk", ConfigMap())
	case whiskDestroyCmd.FullCommand():
		WhiskDestroy()
	case redisDeployCmd.FullCommand():
		RedisDeploy()
	case redisDestroyCmd.FullCommand():
		RedisDestroy()
	case inputCmd.FullCommand():
		if !*inputSelectFlag {
			fmt.Printf("result: '%s'\n", Input("Input Test", *inputArgCmd))
		} else {
			fmt.Printf("select: '%s'\n", Select("Select Test", *inputArgCmd))
		}
	default:
		return false
	}
	return true
}

func parse(cmd string) {
	// debugging (hidden) commands
	if parseDebug(cmd) {
		return
	}
	// user visible commands
	switch cmd {
	// Start
	case startCmd.FullCommand():
		info, err := Start()
		ShowError(err)
		if err == nil {
			PropagateConfig(info)
			if !*skipOpenBrowser {
				time.Sleep(2 * time.Second)
				browser.OpenURL(BrowserURL)
			}
		}
	// Stop
	case stopCmd.FullCommand():
		Stop()
	case restartCmd.FullCommand():
		Stop()
		info, err := Start()
		ShowError(err)
		if err == nil {
			PropagateConfig(info)
		}
		fmt.Println("\nRestarted, please reload the browser.")
	// Init
	case initCmd.FullCommand():
		dir, err := Init(*initDirArg, *initRepoArg, os.Stderr)
		if err == nil {
			err = Configure(dir)
		}
		ShowError(err)
	// Status
	case statusCmd.FullCommand():
		dockerStatus("iogw-openwhisk")
		dockerStatus("iogw-redis")
		dockerStatus("iogw-theia")
		dockerStatus("iogw-scheduler")
	}
}

// Main entrypoint for wskide
func Main(version string) {
	kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version(version).Author(Author)
	kingpin.HelpFlag.Short('h')
	kingpin.CommandLine.Help = Description
	cmd := kingpin.Parse()
	if *verboseFlag {
		log.SetLevel(log.TraceLevel)
	}
	parse(cmd)
}
