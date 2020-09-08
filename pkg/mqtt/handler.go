package mqtt

import log "github.com/sirupsen/logrus"

func SendPayload(payload string) {
	if mqttClient.IsConnected() {
		if mqttToken := mqttClient.Publish(appConf.MqttTopic, byte(*&appConf.MqttQos), false, payload); mqttToken.Wait() && mqttToken.Error() != nil {
			log.WithFields(log.Fields{"mqttBroker": appConf.MqttBroker, "mqttClientID": appConf.MqttClientID, "mqttError": mqttToken.Error(), "mqttPayload": payload, "mqttQos": appConf.MqttQos, "mqttTopic": appConf.MqttTopic}).Error("MQTT client error, cannot send payload")
		} else {
			log.WithFields(log.Fields{"mqttBroker": appConf.MqttBroker, "mqttClientID": appConf.MqttClientID, "mqttPayload": payload, "mqttQos": appConf.MqttQos, "mqttTopic": appConf.MqttTopic}).Debug("MQTT client payload sent")
		}
	} else {
		log.WithFields(log.Fields{"mqttBroker": appConf.MqttBroker, "mqttClientID": appConf.MqttClientID}).Error("MQTT client is not connected, cannot send payload")
	}
}
