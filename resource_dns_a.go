package main

import (
	"errors"
	"log"
	"net"
	"strings"

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
			"ptr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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

	if err != nil {
		return err
	}

	d.SetId("A_z:" + zone + "_n:" + name + "_ip:" + ip.String())	

	ptr := d.Get("ptr").(string)

	if ptr != "" {
		ptrArr, lastByteArr := computePtrAndLastByte(ptr, ip)	
		err = c.AddDNSRecordPTR(zone, ip, name, ptrArr, lastByteArr)
	}	

	return err
}

func computePtrAndLastByte(ptr string, ip net.IP) ([]string, []string) {
	ptrArr := strings.Split(ptr, ".")
	ipArr := strings.Split(ip.String(), ".")
	lastByteArr := make([]string, 0)
	if len(ptrArr) == 3 {
		lastByteArr[0] = ipArr[3]
	}
	if len(ptrArr) == 2 {
		lastByteArr = ipArr[2:4]
	}
	return ptrArr, lastByteArr
}

func deleteRecordA(d *schema.ResourceData, m interface{}) error {
	c := m.(*Communicator)
	c.Connect()

	ip := net.ParseIP(d.Get("ip").(string))

	ptr := d.Get("ptr").(string)

	if ptr != "" {
		ptrArr, lastByteArr := computePtrAndLastByte(ptr, ip)	
		err := c.RemoveDNSRecordPTR(ptrArr, lastByteArr)
		if err != nil {
			return err
		}
	}	

	return c.RemoveDNSRecordA(d.Get("zone").(string), ip, d.Get("name").(string))
}

func readRecordA(d *schema.ResourceData, m interface{}) error {
	return nil
}

func updateRecordA(d *schema.ResourceData, m interface{}) error {
	return nil
}
