package main

import (
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"

	"github.com/samuelebistoletti/tcp-mqtt-forwarder/pkg/config"
	"github.com/samuelebistoletti/tcp-mqtt-forwarder/pkg/mqtt"
	"github.com/samuelebistoletti/tcp-mqtt-forwarder/pkg/tcp"
)

var (
	Version   = "development"
	BuildTime = "1970-01-01"

	ssmlOutput string
)

func main() {
	log.WithFields(log.Fields{"version": Version, "build_time": BuildTime}).Info("TCP-MQTT forwarder started")
	conf := config.Load()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGPIPE, syscall.SIGKILL, syscall.SIGHUP)

	mqtt.InitClient(conf)
	go tcp.StartServer(conf)

	sig := <-c

	log.WithField("signal", sig).Info("exiting the program. Bye!")
}
