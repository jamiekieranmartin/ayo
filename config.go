package ayo

import (
	"io/ioutil"

	"github.com/pelletier/go-toml"
)

// Interface denotes a generalised IO
type Interface interface {
	Listen() Listener
	Send() Sender
}

// IO is an input/output configuration
type IO struct {
	TCP  []TCP
	UDP  []UDP
	HTTP []HTTP
}

// Interfaces from IO
func (i *IO) Interfaces() []Interface {
	interfaces := []Interface{}

	for _, tcp := range i.TCP {
		interfaces = append(interfaces, &tcp)
	}

	for _, udp := range i.UDP {
		interfaces = append(interfaces, &udp)
	}

	for _, http := range i.HTTP {
		interfaces = append(interfaces, &http)
	}

	return interfaces
}

// Config denotes the main configuration structure
type Config struct {
	Input  IO
	Output IO
}

// ReadConfigFile from given path
func ReadConfigFile(path string) (*Config, error) {
	// read file
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// translate TOML to Config
	config := &Config{}
	err = toml.Unmarshal(f, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
