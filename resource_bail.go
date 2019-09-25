package main

import (
	"errors"
	"net"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceBail() *schema.Resource {
	return &schema.Resource{
		Create: createBail,
		Delete: deleteBail,
		Read:   readBail,
		Update: updateBail,

		Schema: map[string]*schema.Schema{
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func createBail(d *schema.ResourceData, m interface{}) error {

	mac, err := net.ParseMAC(d.Get("mac").(string))
	if err != nil {
		return errors.New("Invalid mac Address")
	}

	ipv4, ipv4Net, err := net.ParseCIDR(d.Get("ip").(string))

	if err != nil {
		return errors.New("Invalid ip address, should be in CIDR address")
	}

	c := m.(*Communicator)
	c.Connect()

	c.AddBail(NormalizeMacWindows(mac.String()), ipv4, ipv4Net.IP, d.Get("description").(string))
	d.SetId(mac.String() + "_" + ipv4.String())
	return nil
}

func deleteBail(d *schema.ResourceData, m interface{}) error {
	_, ipv4Net, err := net.ParseCIDR(d.Get("ip").(string))

	if err != nil {
		return errors.New("Invalid ip address, should be in CIDR address")
	}

	c := m.(*Communicator)
	c.Connect()
	return c.RemoveBail(NormalizeMacWindows(d.Get("mac").(string)), ipv4Net.IP)
}

func readBail(d *schema.ResourceData, m interface{}) error {
	return nil
}

func updateBail(d *schema.ResourceData, m interface{}) error {
	return nil
}
