package ayo

import (
	"fmt"
	"net"
	"strings"
)

// UDP interface
type UDP struct {
	Interface
	Hostname string
	Port     string
}

// Addr from hostname and port
func (u *UDP) Addr() string {
	return fmt.Sprintf("%s:%s", u.Hostname, u.Port)
}

// Listen to UDP server
func (u *UDP) Listen() Listener {
	return func(channel chan<- string) error {
		port := fmt.Sprintf(":%s", u.Port)

		// start UDP server on configured port
		l, err := net.ListenPacket("udp", port)
		if err != nil {
			return err
		}
		defer l.Close()

		fmt.Printf("[udp]: Listening on %s\n", u.Port)

		for {
			// read from the buffer
			buf := make([]byte, 1024)
			n, _, err := l.ReadFrom(buf[:])
			if err != nil {
				return err
			}

			// trim delimiter and unneccessary spaces for channel send
			msg := strings.TrimSpace(string(buf[:n]))
			channel <- msg

			fmt.Printf("[udp]<-: %s | %s\n", u.Port, msg)
		}
	}
}

// Send UDP requests
func (u *UDP) Send() Sender {
	fmt.Printf("[udp]: Sending to %s\n", u.Addr())

	return func(channel <-chan string) error {
		// generate address and dial connection
		addr := u.Addr()
		conn, err := net.Dial("udp", addr)
		if err != nil {
			return err
		}

		// iterate over messages and send
		for msg := range channel {
			_, err = conn.Write([]byte(msg))
			if err != nil {
				return err
			}

			fmt.Printf("[udp]->: %s | %s\n", u.Port, msg)
		}

		return conn.Close()
	}
}
