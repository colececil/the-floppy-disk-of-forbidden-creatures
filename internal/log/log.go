package log

import (
	"log"
	"os"
)

var Logger *log.Logger
var file *os.File

// Init initializes the logger.
func init() {
	var err error
	if file, err = os.Create("summon.log"); err != nil {
		panic(err)
	}
	Logger = log.New(file, "", log.LstdFlags)
}
