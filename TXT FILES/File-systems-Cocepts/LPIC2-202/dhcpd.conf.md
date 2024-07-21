# dhcpd.conf

The `dhcpd.conf` file is the configuration file for `dhcpd`, the ISC DHCP server. This file contains all the configuration settings for the DHCP server, including the IP address ranges (scopes) it will manage, options it will provide to clients, and other operational parameters.

### Structure and Syntax

The `dhcpd.conf` file uses a straightforward, declarative syntax. Here’s an overview of the key components and directives:

1. **Global Parameters**: These settings apply to the entire DHCP server.
2. **Subnet Declarations**: These define IP address ranges and options specific to a subnet.
3. **Host Declarations**: These provide fixed address assignments for specific hosts.

### Example Configuration

Here is a basic example of a `dhcpd.conf` file:

```plaintext
# Global Parameters
default-lease-time 600;
max-lease-time 7200;
option domain-name "example.org";
option domain-name-servers ns1.example.org, ns2.example.org;

# Subnet Declaration
subnet 192.168.1.0 netmask 255.255.255.0 {
    range 192.168.1.10 192.168.1.100;
    option routers 192.168.1.1;
    option broadcast-address 192.168.1.255;
    option subnet-mask 255.255.255.0;
}

# Another Subnet Declaration
subnet 192.168.2.0 netmask 255.255.255.0 {
    range 192.168.2.10 192.168.2.100;
    option routers 192.168.2.1;
    option broadcast-address 192.168.2.255;
    option subnet-mask 255.255.255.0;
}

# Host Declaration
host specialclient {
    hardware ethernet 00:11:22:33:44:55;
    fixed-address 192.168.1.50;
}
```

### Key Components Explained

#### Global Parameters

- **default-lease-time**: The default lease time in seconds for IP addresses if the client does not request a specific time.
- **max-lease-time**: The maximum lease time in seconds for IP addresses.
- **option domain-name**: The domain name to be assigned to clients.
- **option domain-name-servers**: The DNS servers to be assigned to clients.

#### Subnet Declarations

- **subnet**: Defines a subnet and its netmask.
- **range**: Specifies the range of IP addresses that the DHCP server can assign to clients within the subnet.
- **option routers**: The default gateway for clients on this subnet.
- **option broadcast-address**: The broadcast address for this subnet.
- **option subnet-mask**: The subnet mask for clients on this subnet.

#### Host Declarations

- **host**: Defines a specific client by its MAC address.
- **hardware ethernet**: The MAC address of the client.
- **fixed-address**: A fixed IP address to be assigned to this client.

### Advanced Configuration Options

#### Dynamic DNS Updates

The DHCP server can be configured to update DNS records dynamically. This requires additional configuration in both `dhcpd.conf` and `named.conf`.

```plaintext
ddns-update-style interim;
ignore client-updates;

zone example.org. {
    primary 192.168.1.1;
    key rndc-key;
}

zone 1.168.192.in-addr.arpa. {
    primary 192.168.1.1;
    key rndc-key;
}
```

#### Failover Configuration

For high availability, two DHCP servers can be configured to share load and provide redundancy.

```plaintext
failover peer "dhcp-failover" {
    primary;
    address 192.168.1.1;
    port 519;
    peer address 192.168.1.2;
    peer port 520;
    max-response-delay 30;
    max-unacked-updates 10;
    load balance max seconds 3;
    mclt 3600;
    split 128;
    hba 00:00:00:01:00:00:00:00:00:00:00:00:00:00:00:00;
}
```

### Practical Use Cases

#### Adding a New Subnet

1. **Edit `dhcpd.conf`**:
   ```bash
   sudo nano /etc/dhcp/dhcpd.conf
   ```
   Add a new subnet declaration:
   ```plaintext
   subnet 192.168.3.0 netmask 255.255.255.0 {
       range 192.168.3.10 192.168.3.100;
       option routers 192.168.3.1;
       option broadcast-address 192.168.3.255;
       option subnet-mask 255.255.255.0;
   }
   ```

2. **Restart the DHCP Server**:
   ```bash
   sudo systemctl restart isc-dhcp-server
   ```

#### Assigning a Fixed IP to a Client

1. **Edit `dhcpd.conf`**:
   ```bash
   sudo nano /etc/dhcp/dhcpd.conf
   ```
   Add a host declaration:
   ```plaintext
   host printer {
       hardware ethernet 11:22:33:44:55:66;
       fixed-address 192.168.1.60;
   }
   ```

2. **Restart the DHCP Server**:
   ```bash
   sudo systemctl restart isc-dhcp-server
   ```

### Security Considerations

- **Restrict Access**: Limit the range of IP addresses to only what is necessary.
- **Logging**: Enable logging to monitor DHCP server activities.
- **Backup**: Regularly back up the `dhcpd.conf` file to prevent data loss.

### Conclusion

The `dhcpd.conf` file is the heart of the ISC DHCP server configuration, allowing administrators to define how IP addresses are allocated, what options are provided to clients, and how the server interacts with other network services. By understanding its structure and options, you can effectively manage and customize your DHCP server to fit your network needs.
