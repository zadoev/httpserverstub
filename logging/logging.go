package logging

import (
	"log"
	"os"
)

const FMT = log.Ldate | log.Lmicroseconds | log.Lshortfile

var Trace = log.New(os.Stdout, "Trace: ", FMT)
var Info = log.New(os.Stdout, "Info: ", FMT)
//var Warning = log.New(os.Stdout, "Warning: ", FMT)
var Error = log.New(os.Stdout, "Error: ", FMT)
