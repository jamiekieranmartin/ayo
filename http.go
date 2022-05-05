package ayo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// HTTP interface
type HTTP struct {
	Port string
	Uri  string
}

// Listen to HTTP requests
func (h *HTTP) Listen() Listener {
	return func(channel chan<- string) error {
		// listen on /
		http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
			// read body
			body, err := ioutil.ReadAll(req.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}

			// pass msg to channel
			channel <- string(body)
			fmt.Printf("[http]<-: %s | %s\n", h.Port, body)

			w.WriteHeader(http.StatusOK)
		})

		// start HTTP server
		port := fmt.Sprintf(":%s", h.Port)
		fmt.Printf("[http]: Listening on %s\n", h.Port)
		return http.ListenAndServe(port, nil)
	}
}

// Send HTTP requests
func (h *HTTP) Send() Sender {
	fmt.Printf("[http]: Sending to %s\n", h.Uri)

	return func(channel <-chan string) error {
		// iterate over channel messages
		for msg := range channel {
			// prepare buffer and send payload via POST
			buf := new(bytes.Buffer)
			json.NewEncoder(buf).Encode(msg)

			_, err := http.Post(h.Uri, "application/json", buf)
			if err != nil {
				return err
			}

			fmt.Printf("[http]->: %s | %s\n", h.Uri, msg)
		}

		return nil
	}
}
