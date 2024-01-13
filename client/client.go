package main

import (
	"github.com/codecat/go-enet"
	"github.com/codecat/go-libs/log"
)

func client() {
	// Initialize enet
	enet.Initialize()

	// Create a client host
	client, err := enet.NewHost(nil, 1, 1, 0, 0)
	if err != nil {
		log.Error("Couldn't create host: %s", err.Error())
		return
	}

	// Connect the client host to the server
	peer, err := client.Connect(enet.NewAddress("127.0.0.1", 8095), 1, 0)
	if err != nil {
		log.Error("Couldn't connect: %s", err.Error())
		return
	}

	// The event loop
	for {
		// Wait until the next event
		ev := client.Service(1000)

		// Send a ping if we didn't get any event
		if ev.GetType() == enet.EventNone {
			peer.SendString("ping", 0, enet.PacketFlagReliable)
			continue
		}

		switch ev.GetType() {
		case enet.EventConnect: // We connected to the server
			log.Info("Connected to the server!")

		case enet.EventDisconnect: // We disconnected from the server
			log.Info("Lost connection to the server!")

		case enet.EventReceive: // The server sent us data
			packet := ev.GetPacket()
			log.Info("Received %d bytes from server", len(packet.GetData()))
			packet.Destroy()
		}
	}

	// Destroy the host when we're done with it
	client.Destroy()

	// Uninitialize enet
	enet.Deinitialize()
}