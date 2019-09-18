# Windows DHCP/DNS Terraform

This provider use winrm to create dhcp bail, mac filter and dns record for windows.

## Usage

 ``` json
provider "win" {
    host: "192.168.100.155",
    port: "5986",
    endpoint: "wsman",
    username: "",
    password: ""
}
```

To create a bail for a specific ip
``` json
resource "win_dhcp" "bail_1" {
    mac: "2A-F8-AF-19-FD-B2",
    ip: "192.168.168.5/24"
}
```

To create a bail for subnet and generate ip
``` json
resource "win_dhcp" "bail_2" {
    mac: "2A-F8-AF-19-FD-B3",
    subnet: "192.168.168.0/24"
}
```

To create a dns reservation
``` json
resource "win_dns_a" "dns" {
    record: "test.local.foo.lan",
    ip: "192.168.168.0/24"
}
```

Allow a mac address
``` json
resource "win_mac_allow" "mac_1" {
    mac: "2A-F8-AF-19-FD-B3"
}
```