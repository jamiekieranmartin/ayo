package ayo

// App denotes the main application
type App struct {
	services []Service
}

// New App from config
func New(config Config) (*App, error) {
	services := []Service{}

	inputs := config.Input.Interfaces()
	outputs := config.Output.Interfaces()

	// generate services from configured inputs and outputs
	for _, input := range inputs {
		// generate sender services
		sends := []Sender{}
		for _, output := range outputs {
			sends = append(sends, output.Send())
		}

		// compile service of listener and senders
		service := Service{
			Listen: input.Listen(),
			Send:   sends,
		}
		services = append(services, service)
	}

	app := &App{
		services,
	}

	return app, nil
}

// ListenAndServe configured services
func (a *App) ListenAndServe() error {
	err := make(chan error, 1)

	// start services in separate goroutines
	for _, service := range a.services {
		service := service
		go func() {
			err <- service.Start()
		}()
	}

	return <-err
}
