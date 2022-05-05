package ayo

// Listener channel receiver
type Listener = func(channel chan<- string) error

// Sender channel emitter
type Sender = func(channel <-chan string) error

// Service denotes a listen/receive combination
type Service struct {
	Listen Listener
	Send   []Sender
}

// Start the service
func (s *Service) Start() error {
	channel := make(chan string, 5)
	err := make(chan error, 1)

	// listen via given channel
	go func() {
		err <- s.Listen(channel)
	}()

	// send via given channel
	for _, send := range s.Send {
		send := send
		go func() {
			err <- send(channel)
		}()
	}

	// wait for errors and return
	return <-err
}
