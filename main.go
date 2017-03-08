package main

import (
	"flag"

	"./utils/config"
	"./utils/css"
	"./utils/logging"
	"./web"

	log "github.com/Sirupsen/logrus"
)

// func main for starting server. If you call the main with some parameters
// you do some other jobs, e.g. precompiling gcss files
//

func main() {
	// Read commandling arguments
	preCompilePtr := flag.Bool("precompile", false, "Precompiles GCSS Files and starts server (only for development)")
	preCompileOnlyPtr := flag.Bool("precompileonly", false, "Precompiles GCSS Files and quits applicaion (for build)")
	flag.Parse()

	// Precompile GCSS only? Without server start?
	if *preCompileOnlyPtr {
		rc, errCSS := css.CompileCSS()
		if !rc {
			log.Errorf("CSS Precompile failed! %v\n", errCSS)
			panic("Application terminated...")
		}

		// Exit application
		return
	}

	// ---- NORMAL START ----

	// Initialize Config
	rc, errConfig := config.ReadConfig()
	if !rc {
		log.Errorf("Config-File-Error: %v", errConfig)
		panic("Config-File-Error!")
	}

	// Init Logging
	logging.InitLog()
	defer logging.CloseLog()

	// PreComile CSS if started with flag -precompile
	if *preCompilePtr {
		rc, errCSS := css.CompileCSS()
		if !rc {
			log.Errorf("CSS Precompile failed! %v", errCSS)
		}
	}

	// Initialize and start webserver
	web.StartServer(config.GetStringValue("WebserverPort"), config.GetBoolValue("WebserverDebug"))
}
