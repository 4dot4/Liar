package main
import (
	"github.com/codecat/go-enet"
	"github.com/codecat/go-libs/log"
)

func server() {
	// Initialize enet
	enet.Initialize()

	// Create a host listening on 0.0.0.0:8095
	host, err := enet.NewHost(enet.NewListenAddress(8095), 32, 1, 0, 0)
	if err != nil {
		log.Error("Couldn't create host: %s", err.Error())
		return
	}

	// The event loop
	for{
		// Wait until the next event
		ev := host.Service(1000)

		// Do nothing if we didn't get any event
		if ev.GetType() == enet.EventNone {
			continue
		}

		switch ev.GetType() {
		case enet.EventConnect: // A new peer has connected
			log.Info("New peer connected: %s", ev.GetPeer().GetAddress())

		case enet.EventDisconnect: // A connected peer has disconnected
			log.Info("Peer disconnected: %s", ev.GetPeer().GetAddress())

		case enet.EventReceive: // A peer sent us some data
			// Get the packet
			packet := ev.GetPacket()

			// We must destroy the packet when we're done with it
			defer packet.Destroy()

			// Get the bytes in the packet
			packetBytes := packet.GetData()

			// Respond "pong" to "ping"
			if string(packetBytes) == "ping" {
				ev.GetPeer().SendString("pong", ev.GetChannelID(), enet.PacketFlagReliable)
				continue
			}

			// Disconnect the peer if they say "bye"
			if string(packetBytes) == "bye" {
				log.Info("Bye!")
				ev.GetPeer().Disconnect(0)
				continue
			}
		}
	}

	// Destroy the host when we're done with it
	host.Destroy()

	// Uninitialize enet
	enet.Deinitialize()
}