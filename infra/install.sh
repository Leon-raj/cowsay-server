#!/bin/bash

set -e

USER_NAME="cowserver"
BINARY_URL="https://github.com/Leon-raj/cowsay-server/releases/download/v0.1/cowserve"
SERVICE_URL="https://github.com/Leon-raj/cowsay-server/releases/download/v0.1/cows.service"

echo "==> Creating system user..."

if ! id -u "$USER_NAME" >/dev/null 2>&1; then
    sudo useradd \
        --system \
        --no-create-home \
        --shell /usr/sbin/nologin \
        "$USER_NAME"
fi

echo "==> Downloading cowserver..."

sudo curl -fsSL "$BINARY_URL" \
    -o /usr/local/bin/cowserver

sudo chmod 755 /usr/local/bin/cowserver
sudo chown root:root /usr/local/bin/cowserver

echo "==> Downloading systemd service..."

sudo curl -fsSL "$SERVICE_URL" \
    -o /etc/systemd/system/cows.service

sudo chmod 644 /etc/systemd/system/cows.service
sudo chown root:root /etc/systemd/system/cows.service

echo "==> Reloading systemd..."

sudo systemctl daemon-reload

echo "==> Enabling service..."

sudo systemctl enable cows

echo "==> Starting service..."

sudo systemctl restart cows

echo
echo "Installation complete!"
echo
systemctl --no-pager --full status cows