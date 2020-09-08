#!/bin/bash

set -e

d_user=tcp-mqtt-forwarder
user_exists=false

getent passwd $d_user >/dev/null 2>&1 && user_exists=true

if ! $user_exists; then
  useradd --system --no-create-home $d_user
fi

if [ -f /usr/bin/DEB_PROJECT ] && [ -f /lib/systemd/system/DEB_PROJECT.service ]; then
  systemctl stop DEB_PROJECT.service
fi

if [ ! -d /var/log/DEB_PROJECT ]; then
	mkdir -p /var/log/DEB_PROJECT
	chown -R $d_user:$d_user /var/log/DEB_PROJECT
else
  rm -rf /var/log/DEB_PROJECT
	mkdir -p /var/log/DEB_PROJECT
	chown -R $d_user:$d_user /var/log/DEB_PROJECT
fi

if [ -d /etc/DEB_PROJECT ]; then
	rm -rf /etc/DEB_PROJECT
fi

if [ -d /boot/18months/DEB_PROJECT ]; then
	rm -rf /boot/18months/DEB_PROJECT
fi
