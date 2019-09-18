package main

import (
	"log"

	"github.com/masterzen/winrm"
)

type Communicator struct {
	username string
	password string
	client   *winrm.Client
	endpoint *winrm.Endpoint
}

func (c *Communicator) Connect() error {
	params := winrm.DefaultParameters
	client, err := winrm.NewClientWithParameters(c.endpoint, c.username, c.password, params)
	if err != nil {
		return err
	}
	shell, err := client.CreateShell()
	if err != nil {
		// error here if cannot connect
		return err
	}
	shell.Close()
	c.client = client
	return nil
}

func (c *Communicator) Run() error {
	stdout, stderr, _, _ := c.client.RunWithString("ipconfig /all", "")
	log.Printf(stdout)
	log.Printf(stderr)
	return nil
}
