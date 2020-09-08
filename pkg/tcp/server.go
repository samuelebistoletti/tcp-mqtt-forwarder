package tcp

import (
	"net"
	"os"

	"github.com/samuelebistoletti/tcp-mqtt-forwarder/pkg/config"

	log "github.com/sirupsen/logrus"
)

var (
	appConf *config.Config
)

func StartServer(conf *config.Config) {
	appConf = conf

	log.WithFields(log.Fields{"tcpHost": appConf.TCPHost}).Info("Launching TCP server")
	server, err := net.Listen("tcp", appConf.TCPHost)

	if err != nil {
		log.WithFields(log.Fields{"tcpHost": appConf.TCPHost, "error": err.Error}).Error("Error during TCP server launch")
		os.Exit(1)
	}
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.WithFields(log.Fields{"tcpHost": appConf.TCPHost, "error": err.Error}).Error("Error during incoming TCP connection")
			os.Exit(1)
		}

		go handleRequest(conn)
	}
}
