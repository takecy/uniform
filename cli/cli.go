package cli

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Client struct {
	client *http.Client
	path   string
	prop   string
	port   int
}

func NewCli(path, prop string, port int) *Client {
	return &Client{
		client: &http.Client{Timeout: 5 * time.Second},
		path:   path,
		prop:   prop,
		port:   port,
	}
}

// Alive is request to instance of binary version
func (c *Client) Alive(addresses []string) (incorrectAddresses []string, err error) {
	if len(addresses) == 0 {
		return incorrectAddresses, errors.New("not hosts")
	}

	incorrectAddresses = make([]string, 0, len(addresses))

	for _, address := range addresses {
		req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%d%s", address, c.port, c.path), nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "req.error %s", err.Error())
			incorrectAddresses = append(incorrectAddresses, address)
			continue
		}

		res, err := c.client.Do(req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "req.error %s", err.Error())
			incorrectAddresses = append(incorrectAddresses, address)
			continue
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			fmt.Fprintf(os.Stderr, "req.error %s", err.Error())
			incorrectAddresses = append(incorrectAddresses, address)
			continue
		}
	}

	return
}
