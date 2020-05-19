module github.com/proemergotech/chrome-webrtc-packet-log-parser

go 1.14

require (
	github.com/golang/protobuf v1.4.1
	github.com/pion/rtcp v1.2.2-0.20200509080250-ba3bfe4ae87c
)

replace github.com/pion/rtcp => github.com/proemergotech/rtcp v1.2.2-0.20200518131901-832aad80eb6a
