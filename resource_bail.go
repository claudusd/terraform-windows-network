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
				Optional: true,
				Default:  nil,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"scope_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
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

	c := m.(*Communicator)
	c.Connect()

	if d.Get("ip").(string) == "" {
		ip, err := c.getFreeIp(d.Get("scope_id").(string))
		if err != nil {
			return err
		}
		d.Set("ip", ip.String())
	}

	ipv4 := net.ParseIP(d.Get("ip").(string))
	if ipv4 == nil {
		return errors.New("Invalid ip Address")
	}

	c.AddBail(NormalizeMacWindows(mac.String()), ipv4, d.Get("scope_id").(string), d.Get("description").(string), d.Get("name").(string))
	d.SetId(mac.String() + "_" + ipv4.String())
	return nil
}

func deleteBail(d *schema.ResourceData, m interface{}) error {
	c := m.(*Communicator)
	c.Connect()
	return c.RemoveBail(NormalizeMacWindows(d.Get("mac").(string)), d.Get("scope_id").(string))
}

func readBail(d *schema.ResourceData, m interface{}) error {
	return nil
}

func updateBail(d *schema.ResourceData, m interface{}) error {
	return nil
}
