package main

import (
	"github.com/reaperhero/elasticsearch-alarm/pkg/cmd"
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
}

func main() {
	cmd.Run()
}
