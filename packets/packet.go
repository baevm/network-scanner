package packets

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func Scan() {
	deviceNum := 0
	const size = 1024
	const promiscmode = false

	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatalf("error retrieving devices - %v", err)
	}

	for i, device := range devices {
		fmt.Printf("%d. %s\n", i, device.Description)
	}

	fmt.Println("Select device to scan:")
	fmt.Scanln(&deviceNum)

	device := devices[deviceNum]

	handle, err := pcap.OpenLive(device.Name, size, promiscmode, time.Second*30)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Starting scanning on %s \n", device.Description)

	packetSrc := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSrc.Packets() {
		fmt.Println(packet.Dump())
	}
}
