package config

import (
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var (
	confFile *os.File
	err      error

	conf *Config
)

type Config struct {
	Debug             bool   `json:"debug"`
	LogToFile         bool   `json:"log_to_file"`
	LogFile           string `json:"log_file"`
	LogFormat         string `json:"log_format"`
	TCPHost           string `json:"tcp_host"`
	MqttBroker        string `json:"mqtt_broker"`
	MqttUser          string `json:"mqtt_user"`
	MqttPassword      string `json:"mqtt_password"`
	MqttTopic         string `json:"mqtt_topic"`
	MqttClientID      string `json:"mqtt_client_id"`
	MqttCleanSession  bool   `json:"mqtt_clean_session"`
	MqttAutoreconnect bool   `json:"mqtt_autoreconnect"`
	MqttQos           int    `json:"mqtt_qos"`
}

func Load() (c *Config) {
	if _, err := os.Stat("/etc/tcp-mqtt-forwarder/config.json"); os.IsNotExist(err) {
		log.Warn("Fallback on default configuration")

		conf = &Config{
			Debug:             true,
			LogToFile:         false,
			LogFile:           "/var/log/tcp-mqtt-forwarder/app.log",
			LogFormat:         "text",
			TCPHost:           "0.0.0.0:3333",
			MqttBroker:        "tcp://127.0.0.1:1883",
			MqttUser:          "username",
			MqttPassword:      "password",
			MqttTopic:         "tcpmqtt/forwarder",
			MqttClientID:      "tcp-mqtt-forwarder",
			MqttCleanSession:  true,
			MqttAutoreconnect: true,
			MqttQos:           2,
		}
	} else {
		confFile, err = os.Open("/etc/tcp-mqtt-forwarder/config.json")
		if err != nil {
			log.WithField("path", "/etc/tcp-mqtt-forwarder/config.json").Errorf("Cannot open file: %v", err)
			os.Exit(2)
		}

		decoder := json.NewDecoder(confFile)
		conf = &Config{}

		err = decoder.Decode(&conf)
		if err != nil {
			log.WithField("path", "/etc/tcp-mqtt-forwarder/config.json").Errorf("Cannot decode file: %v", err)
			os.Exit(2)
		}
	}

	if conf.LogToFile {
		log.SetOutput(&lumberjack.Logger{
			Filename:   conf.LogFile,
			MaxSize:    50, // megabytes
			MaxBackups: 7,
			MaxAge:     7, //days
		})
	}

	if conf.Debug {
		log.SetLevel(log.DebugLevel)
		log.Warn("Running in debug mode")
	}

	switch conf.LogFormat {
	case "text":
		log.SetFormatter(&log.TextFormatter{DisableColors: false})
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	}

	return conf
}
