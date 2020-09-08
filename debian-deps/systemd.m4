[Unit]
Description=DEB_PROJECT
After=multi-user.target network.target local-fs.target

[Service]
Type=simple

User=tcp-mqtt-forwarder
Group=tcp-mqtt-forwarder

ExecStart=/usr/bin/DEB_PROJECT

Restart=on-failure

SyslogIdentifier=DEB_PROJECT

[Install]
WantedBy=multi-user.target
