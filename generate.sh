#!/bin/bash

set -e

rm -f /tmp/proto/rtc_event_log*

wget -q -O - https://chromium.googlesource.com/external/webrtc/+/refs/heads/master/logging/rtc_event_log/rtc_event_log.proto?format=TEXT | base64 -d > /tmp/proto/rtc_event_log.proto

cd /tmp/proto && protoc --proto_path=/tmp/proto --go_out=plugins=unmarshaler_all:. /tmp/proto/rtc_event_log.proto
