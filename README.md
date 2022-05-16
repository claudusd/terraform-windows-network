# Windows DHCP/DNS Terraform

This provider use winrm to create dhcp reservations, mac filters and dns records on Windows Server.

## Usage

```hcl
provider "windowsnetwork" {
   host = "192.168.100.155"
   port = "5986"
   endpoint = "wsman"
   username = ""
   password = ""
}
```

## Resources

### windowsnetwork_dhcp_mac_allowed

| Argument    | Required | Definition                                                   |
| ----------- | -------- | ------------------------------------------------------------ |
| mac         | Yes      | Mac address to authorize on DHCP filters                     |
| description | Yes      | Description of the DHCP reservation                          |
| mac_windows | No       | This is the mac address in a format accepted by windows DHCP |

```hcl
resource "windowsnetwork_dhcp_mac_allowed" "a_mac" {
    mac = "2A-F8-AF-19-FD-B2"
    ip "192.168.165.5"
}
```

### windowsnetwork_dhcp_reservation

| Argument    | Required | Definition                                                                            |
| ----------- | -------- | ------------------------------------------------------------------------------------- |
| mac         | Yes      | Mac address of the reservation                                                        |
| ip          | No       | Ip of the reservation. If not set the provider request a free ip from the DHCP server |
| description | Yes      | Description of reservation                                                            |
| scope_id    | Yes      | The DHCP's scope id                                                                   |
| name        | Yes      | The reservation name                                                                  |

```hcl
// Create a reservation with an ip
resource "windowsnetwork_dhcp_reservation" "reservation_with_ip" {
    mac = "2A-F8-AF-19-FD-B2"
    ip = "192.168.168.5"
    description = "A reservation"
    scope_id = "192.168.168.0"
    name = "vm-1"
}

// Create a reservation without ip
resource "windowsnetwork_dhcp_reservation" "reservation_without_ip" {
    mac = "2A-F8-AF-19-FD-B2"
    description = "A reservation"
    scope_id = "192.168.168.0"
    name = "vm-1"
}
```

When this resource is destroyed, the reservation and the lease are removed.

### windowsnetwork_dns_record_a

| Argument | Required | Definition        |
| -------- | -------- | ----------------- |
| name     | Yes      | The record's name |
| ip       | Yes      | The record's ip   |
| zone     | Yes      | The record's zone |

```hcl
resource "windowsnetwork_dns_record_a" "www" {
    name = "www"
    zone = "example.com"
    ip = "192.168.168.5"
}
```
