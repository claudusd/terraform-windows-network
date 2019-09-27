package main

import (
	"errors"
	"log"
	"net"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceRecordA() *schema.Resource {
	return &schema.Resource{
		Create: createRecordA,
		Delete: deleteRecordA,
		Read:   readRecordA,
		Update: updateRecordA,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"zone": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func createRecordA(d *schema.ResourceData, m interface{}) error {

	log.Printf("[DEBUG] validate a ip address")
	ip := net.ParseIP(d.Get("ip").(string))

	if ip == nil {
		return errors.New("Invalid IP Address")
	}

	c := m.(*Communicator)
	c.Connect()

	zone := d.Get("zone").(string)
	name := d.Get("name").(string)

	err := c.AddDNSRecordA(zone, ip, name)

	d.SetId("A_z:" + zone + "_n:" + name + "_ip:" + ip.String())

	return err
}

func deleteRecordA(d *schema.ResourceData, m interface{}) error {
	c := m.(*Communicator)
	c.Connect()

	ip := net.ParseIP(d.Get("ip").(string))

	return c.RemoveDNSRecordA(d.Get("zone").(string), ip, d.Get("name").(string))
}

func readRecordA(d *schema.ResourceData, m interface{}) error {
	return nil
}

func updateRecordA(d *schema.ResourceData, m interface{}) error {
	return nil
}
