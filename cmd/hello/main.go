package main

import (
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	//"github.com/hn8/zerologr"
	"github.com/rs/zerolog"
)

func main() {
	zerologr.NameFieldName = "logger"
	zerologr.NameSeparator = "/"

	zl := zerolog.New(os.Stderr)
	var log logr.Logger = zerologr.New(&zl)

	log.Info("Test merge : msg 1b, take 2")
}
