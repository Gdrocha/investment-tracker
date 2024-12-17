#!/bin/bash

# Build application first
utils/build_application.sh

APP_NAME="investment-tracker"
EXEC_PATH="$INVESTMENT_TRACKER_BUILD_LOCATION/investment-tracker"
WORKING_DIR=$(pwd)
USER=$(whoami)
GROUP=$(id -gn)
SERVICE_PATH="/etc/systemd/system/$APP_NAME.service"
GOPATH=$(go env GOPATH)

cat <<EOF | sudo tee $SERVICE_PATH > /dev/null
[Unit]
Description=Investment tracker go application
After=network.target

[Service]
ExecStart=$EXEC_PATH
WorkingDirectory=$WORKING_DIR
Restart=always
User=$USER
Group=$GROUP
Environment=GOPATH=$GOPATH

[Install]
WantedBy=multi-user.target
EOF

sudo systemctl daemon-reload
sudo systemctl enable $APP_NAME.service
sudo systemctl start $APP_NAME.service
sudo systemctl status $APP_NAME.service