# TCP-MQTT Forwarder

Simple utility that forwards TCP incoming payloads to a MQTT server over the specified topic.

## Quick start in development mode

You can launch the server in dev mode with Foreman utility here https://github.com/ddollar/foreman. After Foreman install simply issue `foreman start` command in the local folder to start the server.

## Quick compile binaries

You can also compile go source code thanks to the makefile here. Simply issue `make all` to compile binaries for amd64, 386 and arm platforms.

## Quick debian package build

You can also prepare a debian package that could be installed directly on your server and be production ready.
Simply issue `make deb` command here.

After debian package install you will find the json file at the following path:

```sh
/etc/tcp-mqtt-forwarder/config.json
```

The server is managed by a SystemD unit, you can control the service with:

```sh
sudo systemctl status tcp-mqtt-forwarder
```

```sh
sudo systemctl start tcp-mqtt-forwarder
```

```sh
sudo systemctl stop tcp-mqtt-forwarder
```

```sh
sudo systemctl restart tcp-mqtt-forwarder
```

## Sample config file

Here is a sample of the config file:

```sh
{
  "log_file": "/var/log/tcp-mqtt-forwarder/app.log",
  "log_format": "text",
  "debug": false,
  "tcp_host": "0.0.0.0:3333",
  "mqtt_broker": "tcp://127.0.0.1:1883",
  "mqtt_user": "username",
  "mqtt_password": "password",
  "mqtt_topic": "tcpmqtt/forwarder",
  "mqtt_client_id": "tcp-mqtt-forwarder",
  "mqtt_clean_session": true,
  "mqtt_autoreconnect": true,
  "mqtt_qos": 2
}
```
