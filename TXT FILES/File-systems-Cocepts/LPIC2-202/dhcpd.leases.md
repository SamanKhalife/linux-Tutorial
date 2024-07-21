# dhcpd.leases

The `dhcpd.leases` file is a critical component of the ISC DHCP server, `dhcpd`. This file stores the current leases that the DHCP server has granted to clients. It keeps track of which IP addresses have been assigned, to which devices, and for how long.

### Structure and Contents

The `dhcpd.leases` file contains a series of lease declarations, each representing a single lease granted by the DHCP server. Each lease includes various pieces of information about the lease, such as the IP address, MAC address, lease duration, and other options provided to the client.

### Example Contents

Here is an example of what the `dhcpd.leases` file might look like:

```plaintext
lease 192.168.1.10 {
  starts 3 2023/07/19 14:37:06;
  ends 3 2023/07/19 15:37:06;
  tstp 3 2023/07/19 15:37:06;
  cltt 3 2023/07/19 14:37:06;
  binding state active;
  next binding state free;
  hardware ethernet 00:11:22:33:44:55;
  uid "\001\000\021\"\033DU";
  client-hostname "client1";
}
```

### Key Components Explained

- **lease <IP address>**: Indicates the start of a lease block for the specified IP address.
- **starts <day> <YYYY/MM/DD> <HH:MM:SS>**: The start time of the lease.
- **ends <day> <YYYY/MM/DD> <HH:MM:SS>**: The end time of the lease.
- **tstp <day> <YYYY/MM/DD> <HH:MM:SS>**: The time when the server will stop considering the lease valid.
- **cltt <day> <YYYY/MM/DD> <HH:MM:SS>**: The client last transaction time, indicating the last time the client and server communicated regarding this lease.
- **binding state <state>**: The current state of the lease, typically "active" or "free".
- **next binding state <state>**: The next state the lease will transition to after it expires or is released.
- **hardware ethernet <MAC address>**: The MAC address of the client to which the IP address is leased.
- **uid <UID>**: The unique identifier for the client.
- **client-hostname <hostname>**: The hostname of the client.

### Lease States

- **active**: The lease is currently in use by a client.
- **free**: The lease is available to be reassigned.
- **backup**: The lease is reserved for failover purposes.
- **expired**: The lease has expired and is no longer valid.

### Practical Use Cases

#### Viewing Active Leases

To see which IP addresses are currently leased, you can simply view the `dhcpd.leases` file:

```bash
cat /var/lib/dhcp/dhcpd.leases
```

#### Parsing the `dhcpd.leases` File

For scripting or automated monitoring purposes, you can parse the `dhcpd.leases` file to extract information about active leases. Here is an example using `awk`:

```bash
awk '/lease/ {ip=$2} /hardware ethernet/ {mac=$3} /client-hostname/ {print ip, mac, $2}' /var/lib/dhcp/dhcpd.leases
```

This script will print the IP address, MAC address, and hostname for each active lease.

#### Managing Stale Leases

Over time, the `dhcpd.leases` file can accumulate stale or expired leases. To manage this, you can manually clean up the file or configure the DHCP server to automatically remove old leases. 

### Security Considerations

- **File Permissions**: Ensure that the `dhcpd.leases` file is readable only by the DHCP server process to prevent unauthorized access.
- **Regular Backups**: Regularly back up the `dhcpd.leases` file to prevent data loss in case of server failure.
- **Monitoring**: Continuously monitor the `dhcpd.leases` file for unusual patterns or unauthorized devices.

### Conclusion

The `dhcpd.leases` file is essential for the operation of the ISC DHCP server, tracking all leases it has granted to clients. Understanding its structure and content allows administrators to effectively manage DHCP leases, troubleshoot network issues, and ensure efficient IP address allocation. Proper management and monitoring of this file are crucial for maintaining a healthy and secure DHCP infrastructure.
