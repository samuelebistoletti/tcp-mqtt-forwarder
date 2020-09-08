#!/bin/bash

set -e

if [ -f /usr/bin/DEB_PROJECT ] ; then
	systemctl stop DEB_PROJECT.service
  systemctl disable DEB_PROJECT.service
fi

if [ -d /etc/DEB_PROJECT ]; then
	rm -rf /etc/DEB_PROJECT
fi
