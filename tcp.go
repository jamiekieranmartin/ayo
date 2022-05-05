package ayo

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// TCP interface
type TCP struct {
	Interface
	Hostname string
	Port     string
}

// Addr from hostname and port
func (t *TCP) Addr() string {
	return fmt.Sprintf("%s:%s", t.Hostname, t.Port)
}

// Listen to TCP requests
func (t *TCP) Listen() Listener {
	return func(channel chan<- string) error {
		port := fmt.Sprintf(":%s", t.Port)

		// start TCP server on configured port
		l, err := net.Listen("tcp4", port)
		if err != nil {
			return err
		}
		defer l.Close()

		fmt.Printf("[tcp]: Listening on %s\n", t.Port)

		for {
			// accept connections
			c, err := l.Accept()
			if err != nil {
				return err
			}

			// read buffer
			data, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
				return err
			}

			// trim delimiter and unneccessary spaces for channel send
			msg := strings.TrimSpace(data)
			channel <- msg

			fmt.Printf("[tcp]<-: %s | %s\n", t.Port, msg)
		}
	}
}

// Send TCP requests
func (t *TCP) Send() Sender {
	fmt.Printf("[tcp]: Sending to %s\n", t.Addr())

	return func(channel <-chan string) error {
		// generate address and dial connection
		addr := t.Addr()
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			return err
		}

		// iterate over messages and send
		for msg := range channel {
			_, err = conn.Write([]byte(msg))
			if err != nil {
				return err
			}

			fmt.Printf("[tcp]->: %s | %s\n", t.Port, msg)
		}

		return conn.Close()
	}
}
