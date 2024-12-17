#!/bin/bash

echo "Installing application"
go install src/main.go

#####

echo "Building application"

INVESTMENT_TRACKER_BUILD_LOCATION='/bin/investment-tracker'

sudo go build -o $INVESTMENT_TRACKER_BUILD_LOCATION/investment-tracker src/main.go

echo "All done! Application stored at: $INVESTMENT_TRACKER_BUILD_LOCATION/investment-tracker"

export INVESTMENT_TRACKER_BUILD_LOCATION