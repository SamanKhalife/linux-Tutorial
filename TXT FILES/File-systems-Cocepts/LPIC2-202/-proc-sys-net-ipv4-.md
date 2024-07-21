# /proc/sys/net/ipv4/

The `/proc/sys/net/ipv4/` directory in Linux contains various virtual files that expose and allow modification of kernel parameters related to the IPv4 network stack. These parameters control various aspects of network behavior and performance, and changes made here are often used for tuning and configuration of the network stack at runtime.

### Key Files in `/proc/sys/net/ipv4/`

Here are some of the important files and parameters you might encounter in `/proc/sys/net/ipv4/`:

1. **`ip_forward`**
   - **Description:** Controls IP packet forwarding.
   - **File Path:** `/proc/sys/net/ipv4/ip_forward`
   - **Values:**
     - `0`: IP forwarding is disabled.
     - `1`: IP forwarding is enabled.
   - **Usage:** Used to enable or disable routing between different network interfaces.

2. **`conf/<interface>/accept_redirects`**
   - **Description:** Controls whether the system accepts ICMP redirects for a specific interface.
   - **File Path:** `/proc/sys/net/ipv4/conf/<interface>/accept_redirects`
   - **Values:**
     - `0`: ICMP redirects are not accepted.
     - `1`: ICMP redirects are accepted.
   - **Usage:** Configure whether to accept or ignore ICMP redirect messages on an interface.

3. **`conf/<interface>/send_redirects`**
   - **Description:** Controls whether the system sends ICMP redirect messages for a specific interface.
   - **File Path:** `/proc/sys/net/ipv4/conf/<interface>/send_redirects`
   - **Values:**
     - `0`: ICMP redirects are not sent.
     - `1`: ICMP redirects are sent.
   - **Usage:** Configure whether to send ICMP redirect messages from an interface.

4. **`tcp_keepalive_time`**
   - **Description:** Defines the interval between TCP keepalive probes.
   - **File Path:** `/proc/sys/net/ipv4/tcp_keepalive_time`
   - **Values:** Time in seconds.
   - **Usage:** Determines how often TCP keepalive probes are sent to ensure the connection is still alive.

5. **`tcp_max_syn_backlog`**
   - **Description:** Sets the maximum number of queued SYN packets.
   - **File Path:** `/proc/sys/net/ipv4/tcp_max_syn_backlog`
   - **Values:** Integer value.
   - **Usage:** Determines the size of the backlog queue for SYN packets waiting to be accepted by the application.

6. **`tcp_retries2`**
   - **Description:** Defines the number of times TCP retransmits a segment before giving up.
   - **File Path:** `/proc/sys/net/ipv4/tcp_retries2`
   - **Values:** Integer value.
   - **Usage:** Controls how many times TCP will retransmit a segment before aborting the connection.

7. **`ip_local_port_range`**
   - **Description:** Defines the range of local ports for outbound connections.
   - **File Path:** `/proc/sys/net/ipv4/ip_local_port_range`
   - **Values:** Two integers specifying the start and end of the port range.
   - **Usage:** Sets the range of port numbers that can be used for outbound connections.

8. **`ip_dynaddr`**
   - **Description:** Controls the dynamic address assignment for IPv4.
   - **File Path:** `/proc/sys/net/ipv4/ip_dynaddr`
   - **Values:**
     - `0`: Dynamic address assignment is disabled.
     - `1`: Dynamic address assignment is enabled.
   - **Usage:** Determines whether IPv4 addresses can be dynamically reassigned.

9. **`ip_icmp_ratelimit`**
   - **Description:** Sets the rate limit for sending ICMP error messages.
   - **File Path:** `/proc/sys/net/ipv4/ip_icmp_ratelimit`
   - **Values:** Time in seconds.
   - **Usage:** Controls the rate at which ICMP error messages are sent to avoid flooding.

10. **`ip_frag_high_thresh` and `ip_frag_low_thresh`**
    - **Description:** Set thresholds for IP fragment reassembly.
    - **File Path:** `/proc/sys/net/ipv4/ip_frag_high_thresh` and `/proc/sys/net/ipv4/ip_frag_low_thresh`
    - **Values:** Integer values.
    - **Usage:** Define the thresholds for IP fragment queue sizes.

### Example Commands

- **Enable IP Forwarding:**
  ```sh
  echo 1 > /proc/sys/net/ipv4/ip_forward
  ```

- **Check Current Value of TCP Keepalive Time:**
  ```sh
  cat /proc/sys/net/ipv4/tcp_keepalive_time
  ```

- **Set Custom Local Port Range:**
  ```sh
  echo "10000 65000" > /proc/sys/net/ipv4/ip_local_port_range
  ```

### Conclusion

The files in `/proc/sys/net/ipv4/` are used to configure various aspects of the IPv4 network stack at runtime. Modifying these files can help optimize network performance and behavior according to your specific needs. However, be cautious when making changes, as improper settings can affect network stability and performance.
