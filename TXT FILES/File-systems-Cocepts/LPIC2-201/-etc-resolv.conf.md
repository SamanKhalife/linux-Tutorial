# /etc/resolv.conf

The `/etc/resolv.conf` file is a configuration file for the DNS resolver on Unix-like operating systems. This file specifies the nameservers that the system uses to resolve domain names into IP addresses. 

#### Format and Usage

The basic format of `/etc/resolv.conf` consists of a series of lines with keywords and values. The most common keywords are `nameserver`, `search`, and `domain`.

##### Keywords

1. **nameserver**
   - Specifies a DNS server IP address.
   - Multiple `nameserver` lines can be included, and they will be queried in order.
   - Example:
     ```sh
     nameserver 8.8.8.8
     nameserver 8.8.4.4
     ```

2. **search**
   - Specifies a list of domains to be searched when resolving a hostname.
   - The resolver appends each domain in the list to the hostname being queried until a match is found.
   - Example:
     ```sh
     search example.com mycompany.local
     ```

3. **domain**
   - Sets the local domain name.
   - It is equivalent to setting a single entry in the `search` line.
   - Example:
     ```sh
     domain example.com
     ```

##### Example /etc/resolv.conf

```sh
# Set the local domain name
domain example.com

# Set the search path
search example.com mycompany.local

# Specify the nameservers
nameserver 8.8.8.8
nameserver 8.8.4.4
nameserver 192.168.1.1
```

In this example:
- The local domain is set to `example.com`.
- The resolver will search `example.com` and `mycompany.local` when resolving hostnames.
- The system will query the DNS servers at `8.8.8.8`, `8.8.4.4`, and `192.168.1.1` in that order.

#### Best Practices

1. **Order of Nameservers**:
   - Place the most reliable and fastest nameservers first. The resolver queries them in the order they appear in the file.

2. **Local DNS Servers**:
   - Include local DNS servers before external ones for faster resolution within a local network.

3. **Limit Search Domains**:
   - Keep the search domain list short to avoid delays in hostname resolution.

4. **Backup Nameservers**:
   - Always provide backup nameservers to ensure resolution continues if the primary server fails.

#### Common Issues and Solutions

1. **Changes Not Persisting**:
   - Some systems overwrite `/etc/resolv.conf` upon reboot or when network services restart. To prevent this, modify the relevant configuration files for your network manager (e.g., NetworkManager, systemd-resolved) or use a custom script.

2. **NetworkManager**:
   - To configure DNS settings permanently with NetworkManager, create or edit a connection profile in `/etc/NetworkManager/system-connections/`.
     ```ini
     [ipv4]
     dns=8.8.8.8;8.8.4.4;
     ignore-auto-dns=true
     ```

3. **systemd-resolved**:
   - If using systemd-resolved, you can configure DNS in `/etc/systemd/resolved.conf`.
     ```ini
     [Resolve]
     DNS=8.8.8.8 8.8.4.4
     ```

4. **DHCP Client Configuration**:
   - If your system uses a DHCP client to obtain DNS settings, you can configure the DHCP client to use specific DNS servers. For example, with `dhclient`, you can edit `/etc/dhcp/dhclient.conf`.
     ```sh
     supersede domain-name-servers 8.8.8.8, 8.8.4.4;
     ```

#### Conclusion

The `/etc/resolv.conf` file is a crucial part of DNS resolution on Unix-like systems. Understanding its format and proper configuration can significantly improve network performance and reliability. By following best practices and addressing common issues, administrators can ensure efficient and effective hostname resolution.
