package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the info severity or above.
	log.SetLevel(log.InfoLevel)
}

func foo() {
	log.WithFields(log.Fields{
		"prefix":      "sensor",
		"temperature": -4,
	}).Info("Temperature changes")
}

func main() {
	foo()
	log.Info("Some info. Earth is not flat.")
	log.Warning("This is a warning")
	log.Error("Not fatal. An error. Won't stop execution")
	log.Fatal("MAYDAY MAYDAY MAYDAY. Execution will be stopped here")
	log.Panic("Do not panic")

}
