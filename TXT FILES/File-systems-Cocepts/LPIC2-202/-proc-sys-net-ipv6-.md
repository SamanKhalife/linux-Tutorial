# /proc/sys/net/ipv6/

The `/proc/sys/net/ipv6/` directory in Linux contains virtual files that expose and allow the configuration of kernel parameters related to the IPv6 network stack. These parameters are used to control various aspects of IPv6 networking behavior and performance.

### Key Files in `/proc/sys/net/ipv6/`

Here are some important files and parameters you might encounter in `/proc/sys/net/ipv6/`:

1. **`conf/<interface>/accept_ra`**
   - **Description:** Controls whether the system accepts Router Advertisements (RAs) on a specific interface.
   - **File Path:** `/proc/sys/net/ipv6/conf/<interface>/accept_ra`
   - **Values:**
     - `0`: Router Advertisements are not accepted.
     - `1`: Router Advertisements are accepted.
   - **Usage:** Determines if the interface should accept RAs for autoconfiguration.

2. **`conf/<interface>/autoconf`**
   - **Description:** Controls IPv6 address autoconfiguration on a specific interface.
   - **File Path:** `/proc/sys/net/ipv6/conf/<interface>/autoconf`
   - **Values:**
     - `0`: Autoconfiguration is disabled.
     - `1`: Autoconfiguration is enabled.
   - **Usage:** Enables or disables automatic address configuration on the interface.

3. **`conf/<interface>/disable_ipv6`**
   - **Description:** Disables IPv6 on a specific interface.
   - **File Path:** `/proc/sys/net/ipv6/conf/<interface>/disable_ipv6`
   - **Values:**
     - `0`: IPv6 is enabled on the interface.
     - `1`: IPv6 is disabled on the interface.
   - **Usage:** Control whether IPv6 is enabled or disabled on the interface.

4. **`conf/<interface>/forwarding`**
   - **Description:** Controls IPv6 packet forwarding on a specific interface.
   - **File Path:** `/proc/sys/net/ipv6/conf/<interface>/forwarding`
   - **Values:**
     - `0`: IPv6 forwarding is disabled.
     - `1`: IPv6 forwarding is enabled.
   - **Usage:** Determines whether packets can be forwarded through the interface.

5. **`conf/<interface>/accept_redirects`**
   - **Description:** Controls whether the system accepts IPv6 redirect messages on a specific interface.
   - **File Path:** `/proc/sys/net/ipv6/conf/<interface>/accept_redirects`
   - **Values:**
     - `0`: Redirects are not accepted.
     - `1`: Redirects are accepted.
   - **Usage:** Configure whether the system accepts ICMPv6 redirect messages.

6. **`conf/<interface>/send_redirects`**
   - **Description:** Controls whether the system sends IPv6 redirect messages on a specific interface.
   - **File Path:** `/proc/sys/net/ipv6/conf/<interface>/send_redirects`
   - **Values:**
     - `0`: Redirects are not sent.
     - `1`: Redirects are sent.
   - **Usage:** Configure whether the system sends ICMPv6 redirect messages.

7. **`ip6_forward`**
   - **Description:** Controls IPv6 packet forwarding.
   - **File Path:** `/proc/sys/net/ipv6/ip6_forward`
   - **Values:**
     - `0`: IPv6 forwarding is disabled.
     - `1`: IPv6 forwarding is enabled.
   - **Usage:** Enable or disable IPv6 forwarding across network interfaces.

8. **`ip6frag_high_thresh` and `ip6frag_low_thresh`**
   - **Description:** Set thresholds for IPv6 fragment reassembly.
   - **File Path:** `/proc/sys/net/ipv6/ip6frag_high_thresh` and `/proc/sys/net/ipv6/ip6frag_low_thresh`
   - **Values:** Integer values.
   - **Usage:** Define thresholds for IPv6 fragment queue sizes.

9. **`tcp_max_syn_backlog`**
   - **Description:** Sets the maximum number of queued SYN packets for TCP.
   - **File Path:** `/proc/sys/net/ipv6/tcp_max_syn_backlog`
   - **Values:** Integer value.
   - **Usage:** Determines the size of the backlog queue for SYN packets in IPv6 TCP connections.

10. **`tcp_retries2`**
    - **Description:** Defines the number of times TCP retransmits a segment before giving up.
    - **File Path:** `/proc/sys/net/ipv6/tcp_retries2`
    - **Values:** Integer value.
    - **Usage:** Controls how many times TCP will retransmit a segment before aborting the connection.

### Example Commands

- **Enable IPv6 Forwarding:**
  ```sh
  echo 1 > /proc/sys/net/ipv6/ip6_forward
  ```

- **Check Current Value of IPv6 Forwarding:**
  ```sh
  cat /proc/sys/net/ipv6/ip6_forward
  ```

- **Disable IPv6 on a Specific Interface:**
  ```sh
  echo 1 > /proc/sys/net/ipv6/conf/<interface>/disable_ipv6
  ```

- **Enable IPv6 Autoconfiguration on a Specific Interface:**
  ```sh
  echo 1 > /proc/sys/net/ipv6/conf/<interface>/autoconf
  ```

### Conclusion

The files in `/proc/sys/net/ipv6/` provide a mechanism for configuring and managing IPv6 network stack parameters. Adjusting these parameters can help with tuning network performance, managing IPv6 behavior, and troubleshooting issues related to IPv6 networking. Changes to these parameters are applied immediately but are not persistent across reboots unless configured in system startup scripts or configuration files.
