package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/masterzen/winrm"
)

type WinrmError struct {
	codeReturn int
	message    string
	stderr     string
}

func (e *WinrmError) Error() string {
	return fmt.Sprintf("Winrm (%d) - %s", e.codeReturn, e.message)
}

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

func (c *Communicator) AddFilterAllowAddress(mac string, description string) error {
	command := fmt.Sprintf(
		"Add-DhcpServerv4Filter -List Allow -macAddress \"%s\" -Description \"%s\"",
		mac,
		description,
	)

	_, stderr, returnCode := c.Execute(command)

	if returnCode != 0 {
		return &WinrmError{returnCode, "Cannot allow mac address in dhcp, maybe already allowed.", stderr}
	}

	return nil
}

func (c *Communicator) RemoveFilterAllowAddress(mac string) error {
	command := fmt.Sprintf(
		"Remove-DhcpServerv4Filter \"%s\"", mac,
	)

	c.Execute(command)

	return nil
}

func (c *Communicator) GetAllAllowedMacAddress() []string {
	stdout, _, _ := c.Execute("Get-DhcpServerv4Filter -List Allow")
	lines := strings.Split(stdout, "\n")

	var macs []string

	re := regexp.MustCompile(`(([0-9ABCDEF]{2})-?){6,8}`)

	for _, element := range lines {
		matched, _ := regexp.MatchString(`^(([0-9ABCDEF]{2})-?){6,8}`, element)
		if matched {
			mac := string(re.Find([]byte(element)))
			macs = append(macs, mac)
		}
	}
	return macs
}

func (c *Communicator) Execute(command string) (string, string, int) {
	stdout, stderr, returnCode, _ := c.client.RunWithString(winrm.Powershell(command), "")
	log.Printf(stdout)
	log.Printf(stderr)
	return stdout, stderr, returnCode
}
