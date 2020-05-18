#!/bin/bash

set -e

rm -f /tmp/proto/rtc_event_log*

wget -O /tmp/proto/rtc_event_log.proto https://raw.githubusercontent.com/jitsi/webrtc/master/logging/rtc_event_log/rtc_event_log.proto

cd /tmp/proto && protoc --proto_path=/tmp/proto --go_out=plugins=unmarshaler_all:. /tmp/proto/rtc_event_log.proto
