# rpcinfo

The `rpcinfo` command in Linux is used to make remote procedure call (RPC) information inquiries to the portmapper service. It displays a list of all registered RPC services, their port numbers, and other details about their current state.

### Basic Usage

To use `rpcinfo`, open a terminal and type:

```bash
rpcinfo
```

This command will query the local portmapper and display information about all registered RPC services.

### Common `rpcinfo` Commands and Options

1. **Listing All Registered RPC Services**
   ```bash
   rpcinfo -p
   ```
   This command lists all RPC services registered with the local portmapper, including their program number, version number, protocol, and port number.

   Example output:
   ```
   program vers proto   port  service
    100000    4   tcp    111  portmapper
    100000    3   tcp    111  portmapper
    100000    2   tcp    111  portmapper
    100000    4   udp    111  portmapper
    100000    3   udp    111  portmapper
    100000    2   udp    111  portmapper
   ```

2. **Listing RPC Services on a Remote Host**
   ```bash
   rpcinfo -p <hostname>
   ```
   Replace `<hostname>` with the name or IP address of the remote host. This command lists all RPC services registered with the portmapper on the specified remote host.

   Example:
   ```bash
   rpcinfo -p remotehost.example.com
   ```

3. **Displaying Information About a Specific RPC Program**
   ```bash
   rpcinfo -T <transport> <hostname> <program> <version>
   ```
   - `transport`: Specifies the transport protocol (e.g., `udp`, `tcp`).
   - `hostname`: Name or IP address of the host.
   - `program`: RPC program number.
   - `version`: RPC program version number.

   Example:
   ```bash
   rpcinfo -T tcp localhost 100000 4
   ```

4. **Deleting a Specific RPC Service Registration**
   ```bash
   rpcinfo -d <program> <version>
   ```
   This command deletes the registration for the specified RPC program and version number.

   Example:
   ```bash
   rpcinfo -d 100000 2
   ```

5. **Broadcasting an RPC Call**
   ```bash
   rpcinfo -b <program> <version>
   ```
   This command broadcasts an RPC call to all servers on the network to determine which ones have the specified program and version number registered.

   Example:
   ```bash
   rpcinfo -b 100000 2
   ```

### Use Cases

- **Troubleshooting Network Issues**: `rpcinfo` can be used to verify if the portmapper service is running and if RPC services are correctly registered.
- **Checking Service Availability**: By querying local or remote systems, administrators can ensure that necessary RPC services are available and accessible.
- **Security Audits**: Listing all registered RPC services on a system can help in identifying potentially unnecessary or insecure services that should be disabled.

### Summary of `rpcinfo` Commands

| Command                                     | Purpose                                                                 |
|---------------------------------------------|-------------------------------------------------------------------------|
| `rpcinfo -p`                                | List all RPC services registered with the local portmapper              |
| `rpcinfo -p <hostname>`                     | List all RPC services registered with the portmapper on a remote host   |
| `rpcinfo -T <transport> <hostname> <program> <version>` | Display information about a specific RPC program on a host              |
| `rpcinfo -d <program> <version>`            | Delete a specific RPC service registration                              |
| `rpcinfo -b <program> <version>`            | Broadcast an RPC call to all servers on the network                     |

### Conclusion

The `rpcinfo` command is a valuable tool for system administrators working with RPC services. It provides detailed information about registered services and can be used for troubleshooting, checking service availability, and performing security audits. Understanding the various options and use cases of `rpcinfo` enhances the ability to manage and secure systems effectively.
