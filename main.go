package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"

	"github.com/golang/protobuf/proto"
	pionRtcp "github.com/pion/rtcp"
	pionRtp "github.com/pion/rtp"

	"github.com/proemergotech/chrome-webrtc-packet-log-parser/webrtc_rtclog"
)

type ExtendedPacket struct {
	pionRtcp.Packet
	Type string
}
type ExtendedRtpPacket struct {
	pionRtp.Packet
	Type  string `json:"type"`
	Error string `json:"error,omitempty"`
}

type ExtendedRtcpPacket struct {
	*webrtc_rtclog.RtcpPacket
	Packets []ExtendedPacket `json:"packets"`
	Error   string           `json:"error,omitempty"`
}

type Event_ExtendedRtcpPacket struct {
	RtcpPacket *ExtendedRtcpPacket
}

func main() {
	args := os.Args[1:]
	path := args[0]

	file, err := os.Open(path)
	if err != nil {
		log.Fatal("error while opening file", err)
	}
	defer file.Close()
	fmt.Printf("%s opened\n", path)

	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("error while opening file", err)
	}

	eventStream := &webrtc_rtclog.EventStream{}
	target := reflect.New(reflect.TypeOf(eventStream).Elem()).Interface().(proto.Message)
	err = proto.Unmarshal(content, target)
	if err != nil {
		log.Fatal("error decoding response", err)
	}

	eventStream = target.(*webrtc_rtclog.EventStream)

	for i, event := range eventStream.Stream {
		if event != nil && event.Type != nil {
			switch *event.Type {
			case webrtc_rtclog.Event_RTCP_EVENT:
				rtcpPacket := event.GetRtcpPacket()
				if rtcpPacket == nil {
					continue
				}

				var unmarshalError string
				packets, err := pionRtcp.Unmarshal(rtcpPacket.PacketData)
				if err != nil {
					unmarshalError = err.Error()
				}

				extendedPackets := make([]ExtendedPacket, 0, len(packets))
				for _, packet := range packets {
					extendedPackets = append(extendedPackets, ExtendedPacket{
						Packet: packet,
						Type:   reflect.ValueOf(packet).Type().String(),
					})
				}

				extendedEventSubtype := Event_ExtendedRtcpPacket{
					RtcpPacket: &ExtendedRtcpPacket{
						RtcpPacket: rtcpPacket,
						Packets:    extendedPackets,
						Error:      unmarshalError,
					},
				}

				b, err := json.Marshal(extendedEventSubtype)
				if err != nil {
					log.Fatal("error while marshaling event", err)
					return
				}
				fmt.Printf("%d:  %s\n", i, string(b))
			case webrtc_rtclog.Event_RTP_EVENT:
				rtpPacket := event.GetRtpPacket()
				if rtpPacket == nil {
					continue
				}

				extendedRtpPacket := ExtendedRtpPacket{
					Type: webrtc_rtclog.MediaType_name[int32(rtpPacket.GetType())],
					Packet: pionRtp.Packet{
						Header:  pionRtp.Header{},
						Payload: rtpPacket.XXX_unrecognized,
					},
				}
				err := extendedRtpPacket.Header.Unmarshal(rtpPacket.Header)
				if err != nil {
					extendedRtpPacket.Error = err.Error()
				}

				b, err := json.Marshal(extendedRtpPacket)
				if err != nil {
					log.Fatal("error while marshaling event", err)
				}
				fmt.Printf("%d:  %s\n", i, string(b))
			default:
				b, err := json.Marshal(event.Subtype)
				if err != nil {
					log.Fatal("error while marshaling event", err)
				}
				fmt.Printf("%d:  %s\n", i, string(b))
			}
		}
	}
}
