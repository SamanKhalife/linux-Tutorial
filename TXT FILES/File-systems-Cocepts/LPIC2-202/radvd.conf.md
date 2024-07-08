# radvd.conf
typical `radvd.conf` file structure and its key directives for configuring `radvd` (Router Advertisement Daemon) on a Linux system:

### Overview of radvd.conf

#### File Location
- **Location**: The `radvd.conf` file is typically located in `/etc/radvd.conf` on most Linux distributions.

#### Basic Structure
- **Global Configuration**: The configuration file starts with global settings that apply to all interfaces, followed by interface-specific configurations.

#### Example Configuration

Here’s a simplified example of a `radvd.conf` file:

```plaintext
interface eth0 {
    AdvSendAdvert on;
    MinRtrAdvInterval 3;
    MaxRtrAdvInterval 10;
    AdvLinkMTU 1500;
    AdvDefaultPreference high;

    prefix 2001:db8:1234:5678::/64 {
        AdvOnLink on;
        AdvAutonomous on;
        AdvRouterAddr on;
    };

    RDNSS 2001:db8:aaaa:bbbb::1 {
        AdvRDNSSLifetime 1800;
    };

    DNSSL example.com {
        AdvDNSSLLifetime 1800;
    };
};
```

#### Directives Explanation

- **`interface eth0 { ... }`**: Defines the configuration for the `eth0` interface.
  - `AdvSendAdvert on;`: Enables sending Router Advertisement messages.
  - `MinRtrAdvInterval`, `MaxRtrAdvInterval`: Minimum and maximum intervals between sending RA messages.
  - `AdvLinkMTU`: Specifies the Maximum Transmission Unit (MTU) for the interface.
  - `AdvDefaultPreference`: Sets the default preference for routes advertised by this router.

- **`prefix` Directive**: Configures the IPv6 prefix to advertise on the network.
  - `AdvOnLink`, `AdvAutonomous`, `AdvRouterAddr`: Specifies whether the prefix is on-link, autonomous (used for autoconfiguration), and if the router should use its address as the default gateway.

- **`RDNSS` and `DNSSL` Directives**: These advertise Recursive DNS Server (RDNSS) and DNS Search List (DNSSL) information respectively.
  - `AdvRDNSSLifetime`, `AdvDNSSLLifetime`: Specifies how long the DNS information is valid.

#### Additional Considerations

- **Multiple Interfaces**: You can configure multiple interfaces by repeating the `interface` block with different interface names (`eth1`, `wlan0`, etc.).
- **Security**: Ensure `radvd` is only running on authorized routers and that access to RA messages is secured to prevent unauthorized advertisements.

#### Applying Changes

After modifying `radvd.conf`, you typically restart the `radvd` service to apply the changes:

```bash
sudo systemctl restart radvd
```

#### Testing and Verification

- Use tools like `tcpdump` or `wireshark` to capture IPv6 traffic and verify that clients are receiving the correct Router Advertisement messages (`radvd` sends) with the configured parameters.

### Conclusion

Understanding and correctly configuring `radvd` is crucial for IPv6 network deployments, ensuring that hosts can autoconfigure their IPv6 addresses and access network services seamlessly. Always refer to the `radvd` documentation and your specific network requirements when configuring `radvd.conf`.
