package main

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"net"

	"github.com/hashicorp/terraform/helper/schema"
)

// https://docs.microsoft.com/en-us/powershell/module/dhcpserver/Add-DhcpServerv4Filter?view=win10-ps

func resourceMacAllow() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
		},
	}
}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	var mac net.HardwareAddr

	if d.Get("mac") == "" {
		log.Printf("[DEBUG] generate a mac address")
		mac = GenerateMac()
	}

	var err error

	if d.Get("mac") != nil {
		log.Printf("[DEBUG] validate a mac address")
		mac, err = net.ParseMAC(d.Get("mac").(string))
		if err != nil {
			log.Printf("[DEBUG] invalid mac address")
		}

	}

	c := m.(*Communicator)
	c.Connect()
	c.Run()
	d.SetId(mac.String())
	return nil
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceServerRead(d, m)
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func GenerateMac() net.HardwareAddr {
	mac, _ := net.ParseMAC(subGenerate("00", 5))
	return mac
}

func subGenerate(before string, count int) string {
	if count == 0 {
		return before
	}
	sub, _ := randomHex(1)
	new := before + ":" + sub
	return subGenerate(new, count-1)
}
