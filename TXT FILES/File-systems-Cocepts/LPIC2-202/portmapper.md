# portmapper
The **portmapper** (or `rpcbind` in newer implementations) is a crucial service in the RPC (Remote Procedure Call) system used by NFS (Network File System) and other network services on Unix-like systems. Here’s an overview of what the portmapper does and its significance in network service management:

### Overview of Portmapper (rpcbind)

#### Purpose
- **Service Mapping**: The portmapper maps RPC service endpoints (services identified by a program number) to their corresponding port numbers on the server.
- **Dynamic Port Assignment**: It dynamically assigns ports to RPC services when they start up, allowing clients to find and communicate with these services efficiently.

#### Key Functions
- **Service Registration**: When an RPC service starts on a server, it registers itself with the portmapper, providing its program number and the port it is using.
- **Port Reservation**: Clients query the portmapper to find out which port a specific RPC service is using, allowing them to establish connections for RPC calls.
- **Service Unregistration**: When an RPC service stops or is terminated, it deregisters itself from the portmapper, freeing up the allocated port.

#### Configuration
- **Configuration File**: Typically, the portmapper configuration file (`rpcbind`) is managed by system services and does not require manual editing under normal circumstances.
- **Security**: Configuration options may include setting access controls to restrict which hosts or networks can query the portmapper, ensuring security against unauthorized access.

#### Usage and Commands
- **Start/Stop**: Portmapper (`rpcbind`) is usually started automatically during system boot. You can check its status or restart it using service management commands like `systemctl` or `service`, depending on your Linux distribution (`systemctl status rpcbind`, `service rpcbind restart`, etc.).

#### Security Considerations
- **Firewall**: Ensure that the portmapper's ports (typically UDP port 111) are open in your firewall to allow clients to access RPC services on the server.
- **Access Controls**: Configure `/etc/hosts.allow` and `/etc/hosts.deny` to control which hosts or networks can access RPC services via the portmapper.

#### Troubleshooting
- **Logs**: Check system logs (`/var/log/messages`, `/var/log/syslog`, etc.) for portmapper (`rpcbind`) related messages to diagnose issues with RPC service registration or port allocation.
- **Service Dependencies**: Ensure that the portmapper (`rpcbind`) is running correctly before attempting to start or access RPC services that rely on it.

### Conclusion
The portmapper (`rpcbind`) plays a vital role in managing RPC services on Unix-like systems, facilitating service discovery and communication between clients and servers. Understanding its function and ensuring proper configuration and security measures are essential for reliable network service operations using NFS and other RPC-based services.
