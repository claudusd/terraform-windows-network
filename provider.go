package main

import (
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/masterzen/winrm"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Winrm Host",
			},
			"port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     5986,
				Description: "Winrm Port",
			},
			"endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "/",
				Description: "Winrm endpoint",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Winrm username",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Winrm password",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"windowsnetwork_dhcp_mac_allowed": resourceMacAllow(),
			"windowsnetwork_dhcp_reservation": resourceDHCPReservation(),
			"windowsnetwork_dns_record_a":     resourceRecordA(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	endpoint := &winrm.Endpoint{
		Host:     d.Get("host").(string),
		Port:     d.Get("port").(int),
		HTTPS:    true,
		Insecure: true,
	}

	communicator := &Communicator{
		username: d.Get("username").(string),
		password: d.Get("password").(string),
		endpoint: endpoint,
	}

	return communicator, nil
}

func NormalizeMacWindows(mac string) string {
	return strings.ToUpper(strings.Replace(mac, ":", "-", -1))
}
