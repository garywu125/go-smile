package main

import (
	"fmt"
	"os"

	"github.com/garywu125/ellipse"
	"github.com/garywu125/ellipse/hello"
	log "github.com/sirupsen/logrus"
	"rsc.io/quote"

	f1 "github.com/garywu125/go-smile/feature01"
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
	f1.PrintFeature01()

	fmt.Println("quote hello:", quote.Hello())

	//Initalise the Init function with value of A,B
	e := ellipse.Init{
		9, 2,
	}

	hello.Greeting()
	//this will give answer as 0.9749960430435691
	log.Info("GetEccentricity:", e.GetEccentricity())
	log.Fatal("MAYDAY MAYDAY MAYDAY. Execution will be stopped here")
	log.Panic("Do not panic")
}
