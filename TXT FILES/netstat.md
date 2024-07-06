# netstat 

The `netstat` command in Linux is used to display network connections, routing tables, interface statistics, masquerade connections, and multicast memberships. It provides detailed information about network connections both incoming and outgoing, routing tables, interface statistics, masquerade connections, and multicast memberships. Here’s a detailed explanation of how to use `netstat` and what information it provides:

### Usage of `netstat`

#### Basic Usage

To use `netstat`, open a terminal and type:

```bash
netstat
```

By default, `netstat` displays a list of active sockets for all protocols, including TCP, UDP, RAW, and UNIX sockets.

#### Options and Output

`netstat` typically provides output with different columns representing various network statistics:

1. **Active Internet Connections (TCP/UDP)**
   - `Proto`: Protocol (tcp, udp, raw, etc.)
   - `Recv-Q`: Receive queue size.
   - `Send-Q`: Send queue size.
   - `Local Address`: Local IP address and port number.
   - `Foreign Address`: Remote IP address and port number.
   - `State`: State of the connection (LISTEN, ESTABLISHED, etc.).

   Example output:
   ```
   Proto Recv-Q Send-Q Local Address           Foreign Address         State
   tcp        0      0 0.0.0.0:22              0.0.0.0:*               LISTEN
   tcp        0      0 192.168.1.100:22        192.168.1.101:12345     ESTABLISHED
   udp        0      0 0.0.0.0:68              0.0.0.0:*
   ```

2. **Routing Table**
   - `-r` option: Displays the kernel routing table.

   Example output:
   ```
   Kernel IP routing table
   Destination     Gateway         Genmask         Flags   MSS Window  irtt Iface
   default         gateway         0.0.0.0         UG        0 0          0 eth0
   192.168.1.0     0.0.0.0         255.255.255.0   U         0 0          0 eth0
   ```

3. **Network Interface Statistics**
   - `-i` option: Displays statistics for network interfaces.

   Example output:
   ```
   Kernel Interface table
   Iface   MTU Met   RX-OK RX-ERR RX-DRP RX-OVR    TX-OK TX-ERR TX-DRP TX-OVR Flg
   eth0   1500   0   10000      0      0      0     9500      0      0      0 BMRU
   lo    16436   0     100      0      0      0      100      0      0      0 LRU
   ```

#### Additional Options

- **Filtering by Protocol:** Use `-t` (TCP), `-u` (UDP), `-l` (listening), or `-a` (all, default) options to filter and display specific types of connections.

  ```bash
  netstat -t         # Display TCP connections
  netstat -u         # Display UDP connections
  netstat -l         # Display listening sockets
  ```

- **Continuous Monitoring:** Specify an interval (in seconds) to continuously monitor network connections.

  ```bash
  netstat -c 2       # Display every 2 seconds
  ```

- **Resolve Names:** Use `-n` option to disable name resolution (faster output).

  ```bash
  netstat -n         # Display IP addresses instead of hostnames
  ```

### Use Cases

- **Network Troubleshooting:** `netstat` helps in identifying network connections, checking port availability, and diagnosing network-related issues.

- **Monitoring:** Provides real-time monitoring of network connections and traffic patterns for performance analysis and security auditing.

- **Configuration Verification:** Validates network configuration settings such as routing tables and interface statistics.

### Conclusion

`netstat` is a versatile command-line tool for displaying network connections, routing tables, and interface statistics on Linux systems. By understanding its output and options, administrators and users can effectively monitor network activities, diagnose network problems, and optimize network performance.

# help 

```
-a, --all   Display all connections, including listening sockets.
-n, --numeric   Display numerical addresses instead of names.
-t, --tcp   Display only TCP connections.
-u, --udp   Display only UDP connections.
-r, --route   Display the routing table.
-i, --interface   Display interface statistics.
-l, --listening   Display listening sockets.
-p, --programs   Display the programs that are using network sockets.
-s, --statistics   Display network statistics.
-c, --continuous   Continuously display network connections.
```



