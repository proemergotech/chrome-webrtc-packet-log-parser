![Linters](https://github.com/proemergotech/chrome-webrtc-packet-log-parser/workflows/Linters/badge.svg)

# chrome-webrtc-packet-log-parser

### Project status

At this point the project is primerly intended to be used by checking it out and modifying main.go, however you can also use it by downloading the binaries from releases.

### Overview
Go command-line application to make chrome generated webrtc `diagnostic packet and event recording` human readable.    
For more information see [here](https://tokbox.com/blog/how-to-get-a-webrtc-diagnostic-recording-from-chrome-49/).


To regenerate proto files, call `docker-compose up` in this directory.

### How to execute 
`go run main.go <log file>` 
  
Example output:
```bash
$ go run main.go webrtc_event_log_20200514_1817_191_3.log 
webrtc_event_log_20200514_1817_191_3.log opened
0:  null
1:  {"AudioReceiverConfig":{"remote_ssrc":1063824803,"local_ssrc":4195875351}}
2:  {"VideoReceiverConfig":{"remote_ssrc":1721891724,"local_ssrc":1,"rtcp_mode":2,"remb":false,"rtx_map":[{"payload_type":96,"config":{"rtx_ssrc":574331814,"rtx_payload_type":97}},{"payload_type":98,"config":{"rtx_ssrc":574331814,"rtx_payload_type":99}},{"payload_type":100,"config":{"rtx_ssrc":574331814,"rtx_payload_type":101}},{"payload_type":102,"config":{"rtx_ssrc":574331814,"rtx_payload_type":103}},{"payload_type":104,"config":{"rtx_ssrc":574331814,"rtx_payload_type":105}},{"payload_type":106,"config":{"rtx_ssrc":574331814,"rtx_payload_type":107}},{"payload_type":108,"config":{"rtx_ssrc":574331814,"rtx_payload_type":109}}],"decoders":[{"name":"VP8","payload_type":96},{"name":"VP9","payload_type":98},{"name":"VP9","payload_type":100},{"name":"H264","payload_type":102},{"name":"H264","payload_type":104},{"name":"H264","payload_type":106},{"name":"H264","payload_type":108}]}}
3:  {"AudioReceiverConfig":{"remote_ssrc":1063824803,"local_ssrc":4195875351,"header_extensions":[{"name":"http://www.ietf.org/id/draft-holmer-rmcat-transport-wide-cc-extensions-01","id":3},{"name":"http://www.webrtc.org/experiments/rtp-hdrext/abs-send-time","id":2},{"name":"urn:ietf:params:rtp-hdrext:sdes:mid","id":4},{"name":"urn:ietf:params:rtp-hdrext:sdes:repaired-rtp-stream-id","id":6},{"name":"urn:ietf:params:rtp-hdrext:sdes:rtp-stream-id","id":5},{"name":"urn:ietf:params:rtp-hdrext:ssrc-audio-level","id":1}]}}
4:  {"AudioSenderConfig":{"ssrc":2997831384,"header_extensions":[{"name":"http://www.ietf.org/id/draft-holmer-rmcat-transport-wide-cc-extensions-01","id":3},{"name":"http://www.webrtc.org/experiments/rtp-hdrext/abs-send-time","id":2},{"name":"urn:ietf:params:rtp-hdrext:sdes:mid","id":4},{"name":"urn:ietf:params:rtp-hdrext:sdes:repaired-rtp-stream-id","id":6},{"name":"urn:ietf:params:rtp-hdrext:sdes:rtp-stream-id","id":5},{"name":"urn:ietf:params:rtp-hdrext:ssrc-audio-level","id":1}]}}
5:  {"AudioReceiverConfig":{"remote_ssrc":1063824803,"local_ssrc":2997831384,"header_extensions":[{"name":"http://www.ietf.org/id/draft-holmer-rmcat-transport-wide-cc-extensions-01","id":3},{"name":"http://www.webrtc.org/experiments/rtp-hdrext/abs-send-time","id":2},{"name":"urn:ietf:params:rtp-hdrext:sdes:mid","id":4},{"name":"urn:ietf:params:rtp-hdrext:sdes:repaired-rtp-stream-id","id":6},{"name":"urn:ietf:params:rtp-hdrext:sdes:rtp-stream-id","id":5},{"name":"urn:ietf:params:rtp-hdrext:ssrc-audio-level","id":1}]}}
6:  {"VideoReceiverConfig":{"remote_ssrc":1721891724,"local_ssrc":1,"rtcp_mode":2,"remb":false,"rtx_map":[{"payload_type":98,"config":{"rtx_ssrc":574331814,"rtx_payload_type":99}},{"payload_type":100,"config":{"rtx_ssrc":574331814,"rtx_payload_type":101}},{"payload_type":96,"config":{"rtx_ssrc":574331814,"rtx_payload_type":97}},{"payload_type":102,"config":{"rtx_ssrc":574331814,"rtx_payload_type":122}},{"payload_type":127,"config":{"rtx_ssrc":574331814,"rtx_payload_type":121}},{"payload_type":125,"config":{"rtx_ssrc":574331814,"rtx_payload_type":107}},{"payload_type":108,"config":{"rtx_ssrc":574331814,"rtx_payload_type":109}}],"header_extensions":[{"name":"http://tools.ietf.org/html/draft-ietf-avtext-framemarking-07","id":8},{"name":"http://www.ietf.org/id/draft-holmer-rmcat-transport-wide-cc-extensions-01","id":3},{"name":"http://www.webrtc.org/experiments/rtp-hdrext/abs-send-time","id":2},{"name":"http://www.webrtc.org/experiments/rtp-hdrext/color-space","id":9},{"name":"http://www.webrtc.org/experiments/rtp-hdrext/playout-delay","id":12},{"name":"http://www.webrtc.org/experiments/rtp-hdrext/video-content-type","id":11},{"name":"http://www.webrtc.org/experiments/rtp-hdrext/video-timing","id":7},{"name":"urn:3gpp:video-orientation","id":13},{"name":"urn:ietf:params:rtp-hdrext:sdes:mid","id":4},{"name":"urn:ietf:params:rtp-hdrext:sdes:repaired-rtp-stream-id","id":6},{"name":"urn:ietf:params:rtp-hdrext:sdes:rtp-stream-id","id":5},{"name":"urn:ietf:params:rtp-hdrext:toffset","id":14}],"decoders":[{"name":"VP9","payload_type":98},{"name":"VP9","payload_type":100},{"name":"VP8","payload_type":96},{"name":"H264","payload_type":102},{"name":"H264","payload_type":127},{"name":"H264","payload_type":125},{"name":"H264","payload_type":108}]}}
7:  {"VideoSenderConfig":{"ssrcs":[3527308836],"header_extensions":[{"name":"http://tools.ietf.org/html/draft-ietf-avtext-framemarking-07","id":8},{"name":"http://www.ietf.org/id/draft-holmer-rmcat-transport-wide-cc-extensions-01","id":3},{"name":"http://www.webrtc.org/experiments/rtp-hdrext/abs-send-time","id":2},{"name":"http://www.webrtc.org/experiments/rtp-hdrext/color-space","id":9},{"name":"http://www.webrtc.org/experiments/rtp-hdrext/playout-delay","id":12},{"name":"http://www.webrtc.org/experiments/rtp-hdrext/video-content-type","id":11},{"name":"http://www.webrtc.org/experiments/rtp-hdrext/video-timing","id":7},{"name":"urn:3gpp:video-orientation","id":13},{"name":"urn:ietf:params:rtp-hdrext:sdes:mid","id":4},{"name":"urn:ietf:params:rtp-hdrext:sdes:repaired-rtp-stream-id","id":6},{"name":"urn:ietf:params:rtp-hdrext:sdes:rtp-stream-id","id":5}],"rtx_ssrcs":[657777563],"rtx_payload_type":99,"encoder":{"name":"VP9","payload_type":98}}}
8:  {"VideoReceiverConfig":{"remote_ssrc":1721891724,"local_ssrc":3527308836,"rtcp_mode":2,"remb":false,"rtx_map":[{"payload_type":98,"config":{"rtx_ssrc":574331814,"rtx_payload_type":99}},{"payload_type":100,"config":{"rtx_ssrc":574331814,"rtx_payload_type":101}},{"payload_type":96,"config":{"rtx_ssrc":574331814,"rtx_payload_type":97}},{"payload_type":102,"config":{"rtx_ssrc":574331814,"rtx_payload_type":122}},{"payload_type":127,"config":{"rtx_ssrc":574331814,"rtx_payload_type":121}},{"payload_type":125,"config":{"rtx_ssrc":574331814,"rtx_payload_type":107}},{"payload_type":108,"config":{"rtx_ssrc":574331814,"rtx_payload_type":109}}],"header_extensions":[{"name":"http://tools.ietf.org/html/draft-ietf-avtext-framemarking-07","id":8},{"name":"http://www.ietf.org/id/draft-holmer-rmcat-transport-wide-cc-extensions-01","id":3},{"name":"http://www.webrtc.org/experiments/rtp-hdrext/abs-send-time","id":2},{"name":"http://www.webrtc.org/experiments/rtp-hdrext/color-space","id":9},{"name":"http://www.webrtc.org/experiments/rtp-hdrext/playout-delay","id":12},{"name":"http://www.webrtc.org/experiments/rtp-hdrext/video-content-type","id":11},{"name":"http://www.webrtc.org/experiments/rtp-hdrext/video-timing","id":7},{"name":"urn:3gpp:video-orientation","id":13},{"name":"urn:ietf:params:rtp-hdrext:sdes:mid","id":4},{"name":"urn:ietf:params:rtp-hdrext:sdes:repaired-rtp-stream-id","id":6},{"name":"urn:ietf:params:rtp-hdrext:sdes:rtp-stream-id","id":5},{"name":"urn:ietf:params:rtp-hdrext:toffset","id":14}],"decoders":[{"name":"VP9","payload_type":98},{"name":"VP9","payload_type":100},{"name":"VP8","payload_type":96},{"name":"H264","payload_type":102},{"name":"H264","payload_type":127},{"name":"H264","payload_type":125},{"name":"H264","payload_type":108}]}}
9:  {"AudioPlayoutEvent":{"local_ssrc":1063824803}}
10:  {"IceCandidatePairConfig":{"config_type":0,"candidate_pair_id":110362309,"local_candidate_type":0,"local_relay_protocol":4,"local_network_type":0,"local_address_family":0,"remote_candidate_type":0,"remote_address_family":0,"candidate_pair_protocol":0}}
11:  {"IceCandidatePairConfig":{"config_type":0,"candidate_pair_id":3617290008,"local_candidate_type":0,"local_relay_protocol":4,"local_network_type":0,"local_address_family":0,"remote_candidate_type":0,"remote_address_family":0,"candidate_pair_protocol":0}}
12:  {"IceCandidatePairConfig":{"config_type":0,"candidate_pair_id":527797281,"local_candidate_type":0,"local_relay_protocol":4,"local_network_type":0,"local_address_family":0,"remote_candidate_type":0,"remote_address_family":0,"candidate_pair_protocol":0}}

...

38765:  {"AudioPlayoutEvent":{"local_ssrc":1063824803}}
38766:  {"RtpPacket":{"incoming":false,"packet_length":79,"header":"kG9iZR5IWCmyr0bYvt4AAiLAm6ZAMBDR"}}
38767:  {"AudioPlayoutEvent":{"local_ssrc":1063824803}}
38768:  {"LossBasedBweUpdate":{"bitrate_bps":5687408,"fraction_loss":0,"total_packets":0}}
38769:  {"RtpPacket":{"incoming":false,"packet_length":1081,"header":"kHw0xU4p977SPnYkvt4AAiLApeMxNGUA"}}
38770:  {"RtpPacket":{"incoming":false,"packet_length":1081,"header":"kHw0xk4p977SPnYkvt4AAiLApeMxNGYA"}}
38771:  {"RtpPacket":{"incoming":false,"packet_length":1081,"header":"kHw0x04p977SPnYkvt4AAiLApeMxNGcA"}}
38772:  {"RtpPacket":{"incoming":false,"packet_length":1081,"header":"kHw0yE4p977SPnYkvt4AAiLApeMxNGgA"}}
38773:  {"AudioPlayoutEvent":{"local_ssrc":1063824803}}
38774:  {"RtpPacket":{"incoming":false,"packet_length":1081,"header":"kHw0yU4p977SPnYkvt4AAiLAqwIxNGkA"}}
38775:  {"RtpPacket":{"incoming":false,"packet_length":1081,"header":"kHw0yk4p977SPnYkvt4AAiLAqwIxNGoA"}}
38776:  {"RtpPacket":{"incoming":false,"packet_length":1081,"header":"kHw0y04p977SPnYkvt4AAiLAqwIxNGsA"}}
38777:  {"RtpPacket":{"incoming":false,"packet_length":1082,"header":"kHw0zE4p977SPnYkvt4AAiLAqwIxNGwA"}}
```

### License

[MIT](LICENSE)