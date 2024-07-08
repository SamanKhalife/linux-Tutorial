# radvd
`radvd`, which stands for Router Advertisement Daemon. `radvd` is used in IPv6 (Internet Protocol version 6) networks to advertise routing information and network configuration parameters to clients (hosts) on the network. Here’s an overview of `radvd` and its key aspects:

### Overview of radvd

#### Purpose
- **Router Advertisement**: `radvd` sends Router Advertisement (RA) messages periodically or in response to Router Solicitation (RS) messages from clients. RA messages contain essential network configuration information.
- **IPv6 Autoconfiguration**: `radvd` enables hosts to automatically configure their IPv6 addresses and other parameters without manual intervention.

#### Key Functions
- **Prefix Advertisement**: `radvd` advertises IPv6 prefix information that clients use to generate their IPv6 addresses.
- **Router Lifetime**: Specifies how long clients should consider the router (running `radvd`) as the default gateway.
- **Other Configuration Parameters**: Besides prefix information, `radvd` can advertise other parameters like DNS server addresses (`RDNSS`), domain search lists (`DNSSL`), and more.

#### Configuration
- **Configuration File**: The main configuration file for `radvd` is typically located at `/etc/radvd.conf`. This file defines IPv6 prefixes, router advertisement intervals, and other configuration options.
- **Interfaces**: `radvd` is configured to operate on specific network interfaces (`interface` directive), where it listens for RS messages and sends out RA messages.

#### Usage and Commands
- **Start/Stop**: `radvd` is managed like other system services. You can start, stop, or restart it using commands like `systemctl` or `service` depending on your Linux distribution (`systemctl start radvd`, `service radvd restart`, etc.).
- **Monitoring**: Tools like `radvdump` can be used to monitor the RA messages sent by `radvd` and verify that clients are receiving the correct network configuration.

#### Security Considerations
- **Access Control**: Ensure that only authorized routers can run `radvd` on your network. Use firewall rules (`iptables` or `firewalld`) to restrict access to the IPv6 network segments where `radvd` operates.
- **Router Advertisement Guard (RA-Guard)**: In environments with stricter security requirements, consider deploying RA-Guard mechanisms on switches to prevent unauthorized devices from advertising themselves as routers (`radvd`).

#### Troubleshooting
- **Logs**: Check system logs (`/var/log/messages`, `/var/log/syslog`, etc.) for `radvd` related messages to diagnose issues with RA message generation, configuration errors, or client connectivity problems.
- **IPv6 Packet Captures**: Use tools like `tcpdump` or `wireshark` to capture and analyze IPv6 traffic on the network, focusing on RA and RS messages exchanged between routers (`radvd`) and clients.

### Conclusion
`radvd` plays a crucial role in IPv6 networks by facilitating automatic network configuration and router discovery. Understanding its configuration options, security implications, and troubleshooting techniques is essential for maintaining an efficient and secure IPv6 network infrastructure.
