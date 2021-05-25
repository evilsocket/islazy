package main

import (
	"github.com/evilsocket/islazy/log"
)

func main() {
	// equivalent to log.Output = ""
	log.Output = "/dev/stdout"
	log.Level = log.DEBUG
	log.OnFatal = log.NoneOnFatal
	log.DateFormat = "06-Jan-02"
	log.TimeFormat = "15:04:05"
	log.DateTimeFormat = "2006-01-02 15:04:05"
	log.Format = "{datetime} {level:color}{level:name}{reset} {message}"

	if err := log.Open(); err != nil {
		panic(err)
	}
	defer log.Close()

	log.Raw("hello world")
	log.Debug("hello world")
	log.Info("hello world")
	log.Important("hello world")
	log.Warning("hello world")
	log.Error("hello world")
	log.Fatal("hello world")

	log.OnFatal = log.ExitOnFatal
	log.NoEffects = true

	log.Raw("hello world")
	log.Debug("hello world")
	log.Info("hello world")
	log.Important("hello world")
	log.Warning("hello world")
	log.Error("hello world")
	log.Fatal("hello world")
}
