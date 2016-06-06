// Package css consisits of funcs for precompiling gcss files
//
package css

import (
	"os"

	log "github.com/Sirupsen/logrus"

	"github.com/yosssi/gcss"
)

// Filpath Constants
//
const (
	gcssPath = "public/gcss/" // Path to GCSS Files (Raw Input)
	cssPath  = "public/css/"  // Path to CSS Files (Compiled Output)
)

// CompileCSS precompiles a gcss file.
//
func CompileCSS() (bool, error) {
	// APP.CSS
	log.Info("Precompile app.gcss")
	err := compileCSSFile("app")
	if err != nil {
		return false, err
	}

	return true, nil
}

// compileCSSFile compiles a single GCSS file. GCSS Files are savedin
// public/gcss an will be compiled to public/css
//
// PARAMETERS:
//   filename - Filename of gcss file saved in public/gcss (e.g. app.gcss)
//
func compileCSSFile(filename string) error {
	// Open Input-File in GCSS Format
	inputfile, err := os.Open(gcssPath + filename + ".gcss")
	if err != nil {
		return err
	}

	defer func() {
		if err := inputfile.Close(); err != nil {
			panic(err)
		}
	}()

	// Open Output-File for compiled CSS
	outputfile, err := os.Create(cssPath + filename + ".css")
	if err != nil {
		return err
	}

	// Compile GCSS to CSS
	_, err = gcss.Compile(outputfile, inputfile)
	if err != nil {
		return err
	}

	return nil
}
