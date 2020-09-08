package mqtt

import (
	"os"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"

	"github.com/samuelebistoletti/tcp-mqtt-forwarder/pkg/config"
)

var (
	appConf    *config.Config
	mqttClient MQTT.Client
)

func InitClient(conf *config.Config) {
	appConf = conf

	clientOpts := MQTT.NewClientOptions()

	clientOpts.AddBroker(appConf.MqttBroker)
	clientOpts.SetUsername(appConf.MqttUser)
	clientOpts.SetPassword(appConf.MqttPassword)
	clientOpts.SetClientID(appConf.MqttClientID)
	clientOpts.SetCleanSession(appConf.MqttCleanSession)
	clientOpts.SetAutoReconnect(appConf.MqttAutoreconnect)

	mqttClient = MQTT.NewClient(clientOpts)

	if mqttToken := mqttClient.Connect(); mqttToken.Wait() && mqttToken.Error() != nil {
		log.WithFields(log.Fields{"mqttBroker": appConf.MqttBroker, "mqttClientID": appConf.MqttClientID, "mqttError": mqttToken.Error()}).Error("MQTT client error, cannot connect")
		os.Exit(1)
	}

	log.WithFields(log.Fields{"mqttBroker": appConf.MqttBroker, "mqttClientID": appConf.MqttClientID}).Info("MQTT client connected")
}
