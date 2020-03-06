# Windows DHCP/DNS Terraform

This provider use winrm to create dhcp reservation, mac filter and dns record for windows.

## Usage

 ``` json
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

|  Argument   | Required      | Definition |
|-------------|---------------|------------ |
| mac         | Yes           | Mac address to allowed on dhcp |
| description | Yes           | Description of reservation |
| mac_windows | No            | This is the mac address in a format accepted by windows dhcp |

``` json
resource "windowsnetwork_dhcp_mac_allowed" "a_mac" {
    mac = "2A-F8-AF-19-FD-B2"
    ip "192.168.165.5"
}
```

### windowsnetwork_dhcp_reservation

|  Argument   | Required | Definition |
|-------------|----------|------------|
| mac         | Yes      | Mac address for reservation |
| ip          | No       | Ip for reservation. If not set the provider request a free ip to the dhcp |
| description | Yes      | Description of reservation |
| scope_id    | Yes      | The dhcp's scope id |
| name        | Yes      | The reservation name |

``` json
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

When this resource is destroy the reservation and the lease are remove.

### windowsnetwork_dns_record_a

|  Argument | Required | Definition |
|-----------|----------|------------|
| name      | Yes      | The record's name |
| ip        | Yes      | The record's ip |
| zone      | Yes      | The record's zone |


``` json
resource "windowsnetwork_dns_record_a" "www" {
    name = "www"
    zone = "example.com"
    ip = "192.168.168.5"
}
```