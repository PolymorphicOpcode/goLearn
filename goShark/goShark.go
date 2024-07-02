package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	snaplen = int32(1600)
	promisc = false
	timeout = pcap.BlockForever
	filter  = "tcp and port 80"
)

func devInfo() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Panicln(err)
	}
	for _, device := range devices {
		fmt.Println(device.Name)
		for _, address := range device.Addresses {
			fmt.Printf("\tIP:%s\n", address.IP)
		}
	}
}

func liveCapture(intf string) {
	handle, err := pcap.OpenLive(intf, snaplen, promisc, timeout)
	if err != nil {
		log.Panicln(err)
	}
	defer handle.Close()

	if err := handle.SetBPFFilter(filter); err != nil {
		log.Panicln(err)
	}

	source := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range source.Packets() {
		fmt.Println(packet)
	}
}

func main() {
	devInfo()
	var intf string
	fmt.Print("Interface to use: ")
	fmt.Scanf("%s", &intf)
	liveCapture(intf)

}
