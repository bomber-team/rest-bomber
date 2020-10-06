package nats_listener

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

/*
CreateNewConnectionToNats - initialize new connection to nats
*/
func CreateNewConnectionToNats(preference *NatsConnectionConfiguration) (*nats.Conn, error) {
	totalWait := time.Minute * time.Duration(preference.MaxWait)
	connectionOpts := []nats.Option{}
	connectionOpts = append(connectionOpts, nats.Name(preference.NameClient))
	connectionOpts = append(connectionOpts, nats.ReconnectWait(time.Duration(preference.ReconnectDelay)))
	connectionOpts = append(connectionOpts, nats.MaxReconnects(int(totalWait/time.Duration(preference.ReconnectDelay))))
	connectionOpts = append(connectionOpts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		log.Println("Disconnected dut to: ", err, " will attempt reconnects for: ", totalWait.Minutes())
	}))
	connectionOpts = append(connectionOpts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Println("Reconnected: ", nc.ConnectedUrl())
	}))
	connectionOpts = append(connectionOpts, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Panic("Exiting ", nc.LastError())
		// TODO: impletemend channel for gracefull shutdown
	}))
	connet, errConnect := nats.Connect(preference.URL, connectionOpts...)
	if errConnect != nil {
		return nil, errConnect
	}
	return connet, nil
}
