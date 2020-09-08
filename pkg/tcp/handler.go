package tcp

import (
	"net"

	"github.com/samuelebistoletti/tcp-mqtt-forwarder/pkg/mqtt"

	log "github.com/sirupsen/logrus"
)

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)

	bufLen, err := conn.Read(buf)
	if err != nil {
		log.WithFields(log.Fields{"tcpHost": appConf.TCPHost, "error": err.Error}).Error("Error during reading TCP message data")
	} else {
		message := string(buf[:bufLen])

		mqtt.SendPayload(message)

		log.WithFields(log.Fields{"tcpHost": appConf.TCPHost}).Debugf("Received TCP message: %s", message)
	}

	conn.Close()
}
