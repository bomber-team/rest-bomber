package discovery

import (
	"net"
	"time"
)

/*Speaker - producer packets for discovered by main backend service on multicast listen*/
type Speaker struct {
	Timeout         int
	Packet          []byte
	Address         string
	MaxDatagramSize int32
}

func (speak *Speaker) RunEcho() error {
	addr, err := net.ResolveUDPAddr("udp", speak.Address)
	if err != nil {
		return err
	}
	connection, err2 := net.DialUDP("udp", nil, addr)
	if err2 != nil {
		return err2
	}

	for {
		connection.Write(speak.Packet)
		time.Sleep(time.Second * time.Duration(speak.Timeout))
	}
}
