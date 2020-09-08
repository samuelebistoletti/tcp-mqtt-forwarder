#!/bin/bash

set -e

chmod 644 /lib/systemd/system/DEB_PROJECT.service

systemctl enable DEB_PROJECT.service
systemctl start DEB_PROJECT.service
