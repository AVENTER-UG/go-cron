package main

import (
	"os/exec"
	"time"

	util "github.com/AVENTER-UG/util/util"
	"github.com/sirupsen/logrus"
)

// BuildVersion of go-cron
var BuildVersion string

// GitVersion is the revision and commit number
var GitVersion string

func execute(command string) {
	logrus.WithField("func", "main.execute").Info("Execute: ", command)
	out, err := exec.Command("/bin/sh", "-c", command).Output()
	if err != nil {
		logrus.WithField("func", "main.execute").Error(err.Error())
	}
	logrus.WithField("func", "main.execute").Debug("Execute Output: ", string(out))
}

func main() {
	var logLevel string
	var cronTime time.Duration
	var command string

	cronTime, _ = time.ParseDuration(util.Getenv("CRON_TIME", "15m"))
	logLevel = util.Getenv("LOGLEVEL", "info")
	command = util.Getenv("COMMAND", "")

	enableSyslog := false
	if util.Getenv("ENABLE_SYSLOG", "false") == "true" {
		enableSyslog = true
	}

	util.SetLogging(logLevel, enableSyslog, "go-cron")
	logrus.Println("go-cron" + " build " + BuildVersion + " git " + GitVersion)

	ticker := time.NewTicker(cronTime)
	defer ticker.Stop()
	//nolint:gosimple
	for {
		select {
		case <-ticker.C:
			if command != "" {
				execute(command)
			}
		}
	}
}
