package main

import (
	"os/exec"
	"time"

	util "github.com/AVENTER-UG/util/util"
	"github.com/sirupsen/logrus"
)

// BuildVersion - build version of go-cron
var BuildVersion string

// GitVersion - git commit number during build
var GitVersion string

func execute(command string) {
	logrus.WithField("func", "main.execute").Info("Execute: ", command)

	out, err := exec.Command("/bin/sh", "-c", command).Output()
	if err != nil {
		logrus.WithField("func", "main.execute").Error(err.Error())
		return
	}

	logrus.WithField("func", "main.execute").Debug("Execute Output: ", string(out))
}

func main() {
	var cronTime time.Duration
	var command string
	var logLevel string

	cronTime, _ = time.ParseDuration(util.Getenv("CRON_TIME", "15m"))
	command = util.Getenv("COMMAND", "")
	logLevel = util.Getenv("LOG_LEVEL", "info")

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
